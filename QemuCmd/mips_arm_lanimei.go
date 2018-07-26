package QemuCmd

import (
	"log"
	"../utils"
	"os/exec"
	"fmt"
)


//创建mips和mipsel命令行
//这里需要注意的是, netStr是网卡的物理地址, 而tapStr则是网卡名字, 这里的网卡名tapStr预先使用恶意样本名作为网卡名
func CreateMipsMipselCmd(mStr string, kernelStr string , hdaStr string, isGraphic bool, append string, macStr string, tapStr string, isMips bool)(cmdArg CmdArgs){
		if utils.Debug_ {
			log.Println("QemuCmd.CreateMipsMipselCmd")
		}
		mips := InitCmdConfig()
		mips.MStr = mStr
		mips.Isgraphic = isGraphic
		mips.KernelStr = kernelStr
		mips.HdaStr = hdaStr
		mips.AppendStr = append
		mips.MacAddr = macStr
		TempMacAddr := "nic,macaddr="+macStr
		TempTapStr := "tap,ifname="+tapStr
		mips.CmdQemu = " -M " + mips.MStr + " -kernel " + mips.KernelStr + " -hda " + mips.HdaStr + " -append " + mips.AppendStr + " -net " + TempMacAddr + " -net " + TempTapStr + " -nographic"
		if isMips {
			mips.Cmd = exec.Command(
				"qemu-system-mips",
				"-M",
				mips.MStr,
				"-kernel",
				mips.KernelStr,
				"-hda",
				mips.HdaStr,
				"-append",
				mips.AppendStr,
				"-net",
				TempMacAddr,
				"-net",
				TempTapStr,
				"-nographic",
				)
			mips.CmdQemu = "qemu-system-mips" + mips.CmdQemu
		}else {
			mips.Cmd = exec.Command(
				"qemu-system-mipsel",
				"-M",
				mips.MStr,
				"-kernel",
				mips.KernelStr,
				"-hda",
				mips.HdaStr,
				"-append",
				mips.AppendStr,
				"-net",
				TempMacAddr,
				"-net",
				TempTapStr,
				"-nographic",
			)
			mips.CmdQemu = "qemu-system-mipsel" + mips.CmdQemu
		}
		return mips
}


//思考：
// 1.怎样将函数作为参数进行传递
// 2.函数参数的string...	如何进行传递
// 3.如何获取到IP地址
// 以上两个特性如何使用
//创建armel和armhf命令行
func CreateArmelArmhfCmd(mStr string, kernelStr string , initrdStr string, hda string, drive string, isGraphic bool, append string, macStr string, tapStr string, isArmel bool)(cmdArg CmdArgs){
	if utils.Debug_ {
		log.Println("QemuCmd.CreateArmelArmhfCmd")
	}
	arm := InitCmdConfig()
	arm.MStr = mStr
	arm.Isgraphic = isGraphic
	arm.KernelStr = kernelStr
	if hda == "" {
		arm.DriveStr = "if=sd,file=" + drive
	}else{
		arm.HdaStr = hda
	}
	arm.AppendStr = append
	arm.InitrdStr = initrdStr
	arm.MacAddr = macStr
	arm.TapStr = tapStr
	TempNetStr := "nic,macaddr="+macStr
	TempTapStr := "tap,ifname="+tapStr
	if !isGraphic {
		if isArmel {
			//armel
			arm.Cmd = exec.Command(
				"qemu-system-arm",
				"-M",
				arm.MStr,
				"-kernel",
				arm.KernelStr,
				"-initrd",
				arm.InitrdStr,
				"-hda",
				arm.HdaStr,
				"-append",
				arm.AppendStr,
				"-net",
				TempNetStr,
				"-net",
				TempTapStr,
				"-nographic",
			)
			//该命令行字符串可有可无
			arm.CmdQemu = "qemu-system-arm -M " + arm.MStr + " -kernel " + arm.KernelStr + " -initrd " + arm.InitrdStr + " -hda " + arm.HdaStr + " -append " + arm.AppendStr + " -net " + TempNetStr + " -net " + TempTapStr + " -nographic"
		}else {
			//armhf
			arm.Cmd = exec.Command(
				"qemu-system-arm",
				"-M",
				arm.MStr,
				"-kernel",
				arm.KernelStr,
				"-initrd",
				arm.InitrdStr,
				"-drive",
				arm.DriveStr,
				"-append",
				arm.AppendStr,
				"-net",
				TempNetStr,
				"-net",
				TempTapStr,
				"-nographic",
			)
			//该命令行字符串可有可无
			arm.CmdQemu = "qemu-system-arm -M " + arm.MStr + " -kernel " + arm.KernelStr + " -initrd " + arm.InitrdStr + " -drive " + arm.DriveStr + " -append " + arm.AppendStr + " -net " + TempNetStr + " -net " + TempTapStr + " -nographic"
		}
	}else{
		if isArmel {
			//armel
			arm.Cmd = exec.Command(
				"qemu-system-arm",
				"-m",
				arm.MStr,
				"-kernel",
				arm.KernelStr,
				"-initrd",
				arm.InitrdStr,
				"-hda",
				arm.HdaStr,
				"-append",
				arm.AppendStr,
				"-net",
				TempNetStr,
				"-net",
				TempTapStr,
			)
			arm.CmdQemu = "qemu-system-arm -M " + arm.MStr + " -kernel " + arm.KernelStr + " -initrd " + arm.InitrdStr + " -hda " + arm.HdaStr + " -append " + arm.AppendStr + " -net " + TempNetStr + " -net " + TempTapStr
		}else {
			//armhf
			arm.Cmd = exec.Command(
				"qemu-system-arm",
				"-m",
				arm.MStr,
				"-kernel",
				arm.KernelStr,
				"-initrd",
				arm.InitrdStr,
				"-drive",
				arm.DriveStr,
				"-append",
				arm.AppendStr,
				"-net",
				TempNetStr,
				"-net",
				TempTapStr,
			)
			arm.CmdQemu = "qemu-system-arm -M " + arm.MStr + " -kernel " + arm.KernelStr + " -initrd " + arm.InitrdStr + " -drive " + arm.DriveStr + " -append " + arm.AppendStr + " -net " + TempNetStr + " -net " + TempTapStr
		}
	}
	return arm
}

//怎么来设计Start函数
//设计成100个缓冲区来运行相关的线程样本, 每次结束一个线程样本, 则会直接结束即可
func(cmdArg *CmdArgs) Start32()(errStart error){
	if utils.Debug_ {
		log.Println("QemuCmd.Start")
	}
	if err := cmdArg.Cmd.Start(); err != nil {
		log.Println(cmdArg.HdaStr, ":Start启动qemu进程时出现问题")
		log.Println(err)
		errStart = err
	}
	log.Printf("%d", cmdArg.Cmd.Process.Pid)
	if err := cmdArg.Cmd.Wait(); err != nil {
		//强制kill退出后， kill会打印出相关的字符串
		log.Println(cmdArg.HdaStr, ":Wait等待qemu进程时出现问题")
		log.Println(err)
		errStart = err
	}
	return nil
}

func(cmdArg *CmdArgs) Run32() {
	if utils.Debug_ {
		log.Println("QemuCmd.Run")
	}
}

func(cmdArg *CmdArgs) Stop32() error {
	if utils.Debug_ {
		log.Println("QemuCmd.Stop")
	}
	//这里要注意的一点是， Pid传到kill命令中的参数， 必须是int整数形式， 而不能是string形式
	pid := fmt.Sprintf("%d", cmdArg.Cmd.Process.Pid)
	killcmd := exec.Command("kill", "-s", "9", pid)
	if err := killcmd.Start(); err != nil {
		log.Println("Kill stop qemu进程出现问题")
		log.Println(err)
		return err
	}
	if err := killcmd.Wait(); err != nil {
		log.Println("Kill wait qemu进程出现问题")
		log.Println(err)
		return err
	}
	return nil
}

func(cmdArg *CmdArgs) Restart32(){
	if utils.Debug_ {
		log.Println("QemuCmd.Restart")
	}

}

