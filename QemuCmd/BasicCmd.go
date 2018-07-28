package QemuCmd

import (
	"os/exec"
	"log"
	"fmt"
	"../utils"
	"strings"
	"path/filepath"
)


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

//建议该MalwarePath下不要包含任何文件夹，只需要是恶意样本文件即可，如果包含了文件夹会导致无法准确识别相关信息
func GetAllMalware(MalwarePath string)(MalwareList []string,  err error) {
	if utils.Debug_ {
		log.Println("BasicCmd.GetAllMalware")
	}
	pattern := MalwarePath + "/*"
	MalwareList, err = filepath.Glob(pattern)
	if err != nil {
		log.Println("GetAllMalware: 获取某目录下所有文件出现错误")
		return nil, err
	}
	return
}




//这里我们通过恶意样本的sha256来判断样本是否已经被养殖过，如果被养殖过。就不会再对其进行养殖。
//获取到Malware的sha256值
func GetMalwareSha256(MalwarePath string)(MalSha256 string, err error) {
	if utils.Debug_ {
		log.Println("BasicCmd.CheckMalwareSha256")
	}
	cmd := exec.Command("sha256sum", MalwarePath)
	out, err := cmd.Output()
	if err != nil {
		log.Println("GetMalwareSha256出现错误：", err)
		return "", err
	}
	outstr := strings.Split(string(out), " ")
	MalSha256 = string(outstr[0])
	return
}


func ParseConfig(ConfigPath string)(filePath *FilesPath){
	if utils.Debug_ {
		log.Println("BasicCmd.GetAllMalware")
	}
	return &FilesPath{
		MalwarePath:"",
		StartScript:"",
		ImageSrcPath:"",
		ImageNameSpecify:"",
		ImageDestPath:"",
		LoadPath:"",
	}
}

