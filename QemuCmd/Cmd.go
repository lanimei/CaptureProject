package QemuCmd

import (
	"log"
	"fmt"
	"../utils"
)

const MaxMips = 100 //最大的启动虚拟机个数


//定义全局变量，记录虚拟机的相关信息
//记录下设置的Mac地址以及设置的网卡名
//数据均记录到boltDB数据库中, 定义的3个map字典来对数据进行记录
var QemuArgs []CmdArgs
var MacStr = make(map[string]bool, MaxMips)
var EthStr = make(map[string]bool, MaxMips)   //用来判断网卡名是否已经被使用
var MalEth = make(map[string]string, MaxMips)
var PidStr []string


func init(){
	if utils.Debug_ {
		log.Println("QemuCmd.init")
	}
	//初始化一百条MacStr EthStr MalEth
	for i := 0; i < 100; i++{
		MacStrTemp := "00:16:3e:00:00:" + fmt.Sprintf("%x", i)
		MacStr[MacStrTemp] = false
		EthStrTemp := "eth" + fmt.Sprint("%d", i)
		EthStr[EthStrTemp] = false
		MalEth[EthStrTemp] = ""
	}
}



//详细掌握qemu的启动命令行即可
//初始化一个qemu
//转换成字符串的方法
//strconv.Itoa(i) && fmt.Sprintf("%d", i)
func InitCmdConfig()(CmdArgs){
	if utils.Debug_ {
		log.Println("QemuCmd.InitMipsConfig")
	}
	cmdArg := CmdArgs{
		CmdQemu: "",
		MStr: "",
		KernelStr: "",
		Isgraphic: false,
		InitrdStr: "",
		HdaStr: "",
		DriveStr: "",
		AppendStr: "",
		MacAddr: "",
		TapStr: "",
		Pid: "",
		Feature: "",
	}
	return cmdArg
}
