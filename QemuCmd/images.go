package QemuCmd

import (
	"log"
	"io/ioutil"
	"os/exec"
	"fmt"
	"../utils"
)

//创建恶意样本的镜像文件
func CreateImage(MalwarePath string, Start string, MonitorStr string)(err error){

	return nil
}


//guestmount的命令操作可参考网址： https://www.wolfcstech.com/2017/10/31/qcow2_on_linux/

//加载镜像文件，并映射到目录中
//这里的os.exec用法如下所示
/*
$ aaa -a a -b b -c c -d d -e d
golang中必须是这种形式exec.Command("aaa", "-a", "a", "-b", "b")
而不能是如下的形式 exec.Command("aaa -a a -b b")
 */
func LoadImage(ImagePath string, MountPath string)(err error){
	if utils.Debug_ {
		log.Println("utils.LoadImage")
	}
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
func UnloadImage(MountPath string)(err error){
	if utils.Debug_ {
		log.Println("utils.UnloadImage")
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


func Makedir(DirPath string)(err error){
	if utils.Debug_ {
		log.Println("utils.Makedir")
	}
	if DirPath == "" {
		err = fmt.Errorf("DirPath为空")
		return
	}
	cmd := exec.Command("mkdir", "-p", DirPath)
	err = cmd.Start()
	if err != nil {
		log.Fatal(err)
	}
	return
}