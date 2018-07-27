package QemuCmd

import (
	"log"
	"io/ioutil"
	"os/exec"
	"fmt"
	"../utils"
	"strings"
)

type FilesPath struct {
	MalwarePath string
	StartScript string
	ImageSrcPath string
	ImageNameSpecify string
	ImageDestPath string
	LoadPath string
}


//路径名必须加上 /  如 /home/lanimei/
//创建恶意样本的镜像文件
func CreateImage(filePaths FilesPath)(err error) {
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
	imageNameSpecifyPath := filePaths.ImageDestPath + "/" + filePaths.ImageNameSpecify   //文件路劲+文件名
	err = CpInitImage(filePaths.ImageSrcPath, imageNameSpecifyPath)
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
	err = CpInitImage(filePaths.MalwarePath, home_malware)
	if err != nil {
		log.Println("CreateImage CpInitImage Malware error")
		return
	}
	startScript := filePaths.LoadPath + "/etc/"
	err = CpInitImage(filePaths.StartScript, startScript)
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


func MakeDir(DirPath string)(err error) {
	if utils.Debug_ {
		log.Println("QemuCmd.Makedir")
	}
	if DirPath == "" {
		err = fmt.Errorf("DirPath为空")
		return
	}
	cmd := exec.Command("mkdir", "-p", DirPath)
	err = cmd.Start()
	if err != nil {
		log.Println("Makedir: ", err)
	}
	err = cmd.Wait()
	log.Printf("Command finished with error: %v", err)
	return
}


func CpInitImage(srcPath string, destPath string)(err error) {
	if utils.Debug_ {
		log.Println("QemuCmd.CpInitImage")
	}
	if srcPath == "" || destPath == "" {
		err = fmt.Errorf("srcPath为空或destPath为空")
		return
	}
	cmd := exec.Command("cp", srcPath, destPath)
	err = cmd.Start()
	if err != nil {
		log.Println("CpInitImage", err)
	}
	err = cmd.Wait()
	log.Printf("Command finished with error: %v", err)
	return
}




//这里首先只针对mipsel和arm样本做养殖
func MalwareClassify(MalwarePath string)(string, error){
	if utils.Debug_ {
		log.Println("QemuCmd.MalwareClassify")
	}
	if MalwarePath == "" {
		err := fmt.Errorf("QemuCmd.MalwareClassify: MalwarePath为空")
		label := ""
		return label, err
	}
	cmd := exec.Command("file", MalwarePath)
	out, err := cmd.Output()
	if err != nil {
		log.Println(out)
		return "", err
	}
	outstr := string(out)
	if strings.Contains(outstr, "MIPS") {
		return "mips", nil
	}
	if strings.Contains(outstr, "ARM") {
		return "arm", nil
	}
	if strings.Contains(outstr, "80386") {
		return "i386", nil
	}
	if strings.Contains(outstr, "x86-64") {
		return "x86-64", nil
	}
	if strings.Contains(outstr, "PowerPC") {
		return "PowerPC", nil
	}
	if strings.Contains(outstr, "SPARC") {
		return "SPARC", nil
	}
	if strings.Contains(outstr, "Renesas") {
		return "Renesas", nil
	}
	err = fmt.Errorf("QemuCmd.MalwareClassify：未找到该文件的类型")
	return "", err
}