package QemuCmd

import (
	"log"
	"io/ioutil"
	"os/exec"
	"fmt"
	"../utils"
)



//This file is to create a image for mips, mipsel, arm and so on


//将FilesPath当成一个中间变量即可, 其中的MalwarePath必须是一个样本的路径名
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


