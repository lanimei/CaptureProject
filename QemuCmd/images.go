package QemuCmd

import (
	"log"
	"io/ioutil"
	"os/exec"
	"fmt"
	"../utils"
	"strings"
)



//This file is to create a image for mips, mipsel, arm and so on


//将FilesPath当成一个中间变量即可, 其中的MalwarePath必须是一个样本的路径名
//StartScript就只是一个文件夹路径名而已， 里面不包含任何文件
type FilesPath struct {
	MalwarePath string `json:"malware_path"`
	StartScript string `json:"start_script"`
	ImageSrcPath string `json:"image_src_path"`
	ImageDestPath string `json:"image_dest_path"`
	LoadPath string `json:"load_path"`
}



//这里面对应的变量都是指定了文件路径名，已经包括文件名字等相关信息
type MalwareSample struct {
	MalwareName string
	MalwarePath string
	ImageNameSpecify string		//一个恶意样本对应一个镜像的承诺
	StartScript string         //启动恶意样本的脚本在之后的代码中可以进行相关的扩展， 例如mipsel的恶意样本可以进行相关扩展即可
	MalSha256 string
	Arch string
	SqueezeOr string
	Pid string
	ImageSrcPath string
}


//默认养殖100个恶意样本, 这里默认养殖100个arm， mips的恶意样本，并且镜像首选为squeeze版本的恶意样本
var Malwares []MalwareSample
var filePaths *FilesPath


func InitFilesPath(){
	if utils.Debug_ {
		log.Println("QemuCmd.init")
	}
	var err error
	filePaths, err = ParseConfig("test.json")
	if err != nil {
		log.Println("Init 初始化QemuCmd失败")
		return
	}
	fileList, err := GetAllMalware(filePaths.MalwarePath)
	if err != nil {
		log.Println("获得恶意样本名字")
	}
	for _, itemPath := range fileList {
		if err := AppendMalware(itemPath, "squeeze"); err != nil {
			log.Println(err)
			continue
		}
	}
}

//暂时只做镜像文件为mipsel和armel的镜像文件和恶意样本
func AppendMalware(itemPath string, squeezeOr string)(err error){
	if utils.Debug_ {
		log.Println("QemuCmd.AppendMalware")
	}
	var item MalwareSample
	item.Arch, err = MalwareClassify(itemPath)
	if err != nil {
		log.Println("未找到该种类型的恶意代码文件,无法养殖该文件:", itemPath)
		return
	}
	if strings.Compare("mips", item.Arch) == 0 && strings.Contains(mipsel_hda_squeeze, squeezeOr) {
		item.ImageSrcPath = filePaths.ImageSrcPath + "/" + mipsel_hda_squeeze
	}else if strings.Compare("arm", item.Arch) == 0 && strings.Contains(armel_hda_squeeze, squeezeOr) {
		item.ImageSrcPath = filePaths.ImageSrcPath + "/" + armel_hda_squeeze
	}else {
		log.Println("lanimei123")
		err = fmt.Errorf("暂时只做mips和arm的恶意样本镜像文件")
		return
	}
	splitSlices := strings.Split(itemPath, "/")
	item.MalwareName = splitSlices[len(splitSlices)-1]
	item.MalwarePath = itemPath
	item.StartScript = filePaths.StartScript + "/rclocal_" + item.MalwareName
	item.SqueezeOr = squeezeOr
	item.ImageNameSpecify = "debian_" + item.Arch + "el_" + item.SqueezeOr + "_" + item.MalwareName
	item.MalSha256, err = GetMalwareSha256(itemPath)
	if err != nil {
		log.Println("无法计算该恶意样本的sha256", itemPath)
		return
	}
	Malwares = append(Malwares, item)
	return	nil
}


//路径名必须不能加 /  如 /home/lanimei
//创建恶意样本的镜像文件
func(malware *MalwareSample)CreateImage()(err error) {
	if utils.Debug_ {
		log.Println("QemuCmd.CreateImage")
	}
	//创建image存放的目的路径
	err = MakeDir(filePaths.ImageDestPath)
	if err != nil {
		log.Println("CreateImage MakeDir error")
		return
	}
	//创建image加载的路径
	err = MakeDir(filePaths.LoadPath)
	if err != nil {
		log.Println("CreateImage MakeDir LoadPath error")
		return
	}
	imageNameSpecifyPath := filePaths.ImageDestPath + "/" + malware.ImageNameSpecify   //文件路劲+文件名
	err = CpInitImage(malware.ImageSrcPath, imageNameSpecifyPath)
	if err != nil {
		log.Println("CreateImage CpInitImage error")
		return
	}
	err = LoadImage(imageNameSpecifyPath, filePaths.LoadPath)
	if err != nil {
		log.Println("CreateImage LoadImage error")
		return
	}
	home_malware := filePaths.LoadPath + "/home/"
	err = CpInitImage(malware.MalwarePath, home_malware)
	if err != nil {
		log.Println("CreateImage CpInitImage Malware error")
		return
	}
	startScript := filePaths.LoadPath + "/etc/rc.local"
	err = CpInitImage(malware.StartScript, startScript)
	if err != nil {
		log.Println("CreateImage CpInitImage startup error")
		return
	}
	err = UnloadImage(filePaths.LoadPath)
	if err != nil {
		log.Println("CreateImage UnloadImage error")
		return
	}
	return nil
}


//该恶意样本所对应的Startup文件由主文件filePaths来生成
//使用ioutil包中的读写函数来对文件进行读写
func(malware *MalwareSample) CreateStartup()(err error){
	if utils.Debug_ {
		log.Println("QemuCmd.CreateStartup")
	}
	//这里首先说到的是
	rcLocalOne := "#!/bin/sh -e\n# rc.local\nsleep 4m\nchmod a+x /home/"
	rcLocalTwo := "\nsleep 4m\n/home/"
	rcLocalExit := "\nexit 0"
	rcLocalResult := rcLocalOne + malware.MalwareName + rcLocalTwo + malware.MalwareName +rcLocalExit
	if err := ioutil.WriteFile(malware.StartScript, []byte(rcLocalResult), 0644); err != nil {
		log.Println("写入malware rc.local出现问题，请即使修改！")
		return err
	}
	return nil
}


//guestmount的命令操作可参考网址： https://www.wolfcstech.com/2017/10/31/qcow2_on_linux/

//这里需要注意的是，每一次执行cmd.exec指令时，即每一次 cmd.Start(), 都应该对应着一次cmd.Wait()


//加载镜像文件，并映射到目录中
//这里的os.exec用法如下所示
/*
$ aaa -a a -b b -c c -d d -e d
golang中必须是这种形式exec.Command("aaa", "-a", "a", "-b", "b")
而不能是如下的形式 exec.Command("aaa -a a -b b")
 */
func LoadImage(ImagePath string, MountPath string)(err error) {
	if utils.Debug_ {
		log.Println("QemuCmd.LoadImage")
	}
	log.Println(ImagePath)
	log.Println(MountPath)
	if ImagePath == "" || MountPath == "" {
		err = fmt.Errorf("ImagePath为空或者MountPath为空")
		return
	}
	args := "/dev/sda1"
	cmd := exec.Command("guestmount", "-a", ImagePath, "-m", args, MountPath)
	stderr, stderr2 := cmd.StderrPipe()
	if stderr2 != nil {
		log.Fatal(stderr2)
	}
	err = cmd.Start()
	if err != nil {
		log.Fatal(err)
	}
	slurp, _ := ioutil.ReadAll(stderr)
	fmt.Printf("%s\n", slurp)
	log.Println("Waiting for command to finish...")
	err = cmd.Wait()
	log.Printf("Command finished with error: %v", err)
	return
}


//卸载镜像文件，将镜像文件所对应的目录卸载
func UnloadImage(MountPath string)(err error) {
	if utils.Debug_ {
		log.Println("QemuCmd.UnloadImage")
	}
	if MountPath == "" {
		err = fmt.Errorf("MountPath为空")
		return
	}
	cmd := exec.Command("guestunmount", MountPath)
	stderr, stderr2 := cmd.StderrPipe()
	if stderr2 != nil {
		log.Fatal(stderr2)
	}
	err = cmd.Start()
	if err != nil {
		log.Fatal(err)
	}
	slurp, _ := ioutil.ReadAll(stderr)   //将错误直接进行输出即可
	fmt.Printf("%s\n", slurp)
	log.Println("Waiting for command to finish...")
	err = cmd.Wait()
	log.Printf("Command finished with error: %v", err)
	return nil
}


