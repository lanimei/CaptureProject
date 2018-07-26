package QemuCmd

import (
	"os/exec"
)

//1.关于qemu指令的执行， 这里主要涉及到产生包括恶意样本的images
//2.开机命令执行， 执行恶意样本
//3.生成可执行样本的pcap包信息
//4.掌控镜像执行的指令相关信息
//5.开辟端口，随时等待外部的连接


//记录样本对应的虚拟机名字, 键值是虚拟机的名字，存储的值是样本的值
//思考：这里是否可以定义两个KeyMap, 即 键值(虚拟机名)---->值(样本名)	键值(样本名)---->值(虚拟机名)
var lanimei map[string]string


type QemuCmd interface{
	Start32()
	Run32()
	Stop32()
	Restart32()
}

//armel
//用另一种标识来替代输入
const (
	 armel_m = "versatilepb"
	 armel_kernel_squeeze = "vmlinuz-2.6.32-5-versatile"
	 armel_initrd_squeeze = "initrd.img-2.6.32-5-versatile"
	 armel_hda_squeeze = "debian_squeeze_armel_standard.qcow2"
	 armel_append_squeeze = "root=/dev/sda1 console=tty0"

	 armel_kernel_wheezy = "vmlinuz-3.2.0-4-versatile"
	 armel_initrd_wheezy = "initrd.img-3.2.0-4-versatile"
	 armel_hda_wheezy = "debian_wheezy_armel_standard.qcow2"
	 armel_append_wheezy = "root=/dev/sda1 console=tty0"

	 Armel_device = "/dev/sda1"
)

//armhf
const (
	armhf_m = "vexpress-a9"
	armhf_kernel_wheezy = "vmlinuz-3.2.0-4-vexpress"
	armhf_initrd_wheezy = "initrd.img-3.2.0-4-vexpress"
	armhf_drive_wheezy = "debian_wheezy_armhf_standard.qcow2"
	armhf_append_wheezy = "root=/dev/mmcblk0p2 console=tty0"

	Armhf_device = "/dev/mmcblk0p2"
)

//mips, 64位则使用64位代码即可
const (
	mips_m = "malta"
	mips_kernel_squeeze_32 = "vmlinux-2.6.32-5-4kc-malta-mips"
	mips_kernel_wheezy_32 = "vmlinux-3.2.0-4-4kc-malta"
	mips_kernel_squeeze_64 = "vmlinux-2.6.32-5-5kc-malta"
	mips_kernel_wheezy_64 = "vmlinux-3.2.0-4-5kc-malta"
	mips_hda_squeeze = "debian_squeeze_mips_standard.qcow2"
	mips_hda_wheezy = "debian_wheezy_mips_standard.qcow2"
	mips_append = "root=/dev/sda1 console=tty0"
	Mips_device = "/dev/sda1"
)

//mipsel
const (
	mipsel_m = "malta"
	mipsel_kernel_squeeze_32 = "vmlinux-2.6.32-5-4kc-malta-mipsel"
	mipsel_kernel_wheezy_32 = "vmlinux-3.2.0-4-4kc-malta"
	mipsel_kernel_squeeze_64 = "vmlinux-2.6.32-5-5kc-malta"
	mipsel_kernel_wheezy_64 = "vmlinux-3.2.0-4-5kc-malta"
	mipsel_hda_squeeze = "debian_squeeze_mipsel_standard.qcow2"
	mipsel_hda_wheezy = "debian_wheezy_mips_standard.qcow2"
	mipsel_append = "root=/dev/sda1 console=tty0"
	Mipsel_device = "/dev/sda1"
)


//mips,arm,x86,powerpc,sparc,sh4

type CmdArgs struct {
	CmdQemu string		//由以下所有参数生成
	Cmd *exec.Cmd

	MStr string			//对应-M参数, 架构所属的参数
	KernelStr string	//分为启动32位的内核还是启动64位的内核
	Isgraphic bool
	InitrdStr string
	HdaStr string		//启动squeezy版本硬盘还是wheezy版本的硬盘
	DriveStr string
	AppendStr string
	MacAddr string	//Mac的具体地址
	TapStr string	//网卡的名字，一般用恶意样本的名字来进行命名
	Pid string
	Feature string //记录qemu-system进程的pid号等进程相关信息
}

//make image, and get the name of images
const(
	image_debian = "debian"

	image_wheezy = "wheezy"
	image_squeeze = "squeeze"

	image_armel = "armel"
	image_armhf = "armhf"

	image_mips = "mips"
	image_mipsel = "mipsel"

	image_x86 = "x86"
	image_amd = "amd64"

)


/*

//armel
type ArmelArgs struct{
	ArmelQemu string
	MStr string
	KernelStr string
	Nographic bool
	InitrdStr string
	HdaStr string
	AppendStr string
	NetStr string   //该参数可由ArmelArgs中的MacAddr和IP推到而来
	MacAddr string
	DeviceName string
	EthStr string
	VM_str string
	Feature string //32 bit or 64 bit, wheezy or squeezy
}

//armhf
type ArmhfArgs struct{
	ArmhfQemu string
	MStr string   //对应-M参数
	KernelStr string
	Nographic bool
	InitrdStr string
	DriveStr string
	AppendStr string
	NetStr string   //该参数可由ArmelArgs中的MacAddr和IP推到而来
	MacAddr string
	DeviceName string
	EthStr string
	VM_str string
	Feature string //32 bit or 64 bit, wheezy or squeezy
}

//i386
type I386Args struct{
	I386Qemu string
	MStr string   //对应-M参数
	KernelStr string
	Nographic bool
	InitrdStr string
	HdaStr string
	AppendStr string
	NetStr string   //该参数可由I386Args中的MacAddr和IP推到而来
	MacAddr string
	EthStr string
	VM_str string
	Feature string //32 bit or 64 bit, wheezy or squeezy
}

//amd64
type Amd64Args struct{
	I386Qemu string
	MStr string   //对应-M参数
	KernelStr string
	Nographic bool
	InitrdStr string
	HdaStr string
	AppendStr string
	NetStr string   //该参数可由Amd64Args中的MacAddr和IP推到而来
	MacAddr string
	EthStr string
	VM_str string
	Feature string //32 bit or 64 bit, wheezy or squeezy
}

//mips
type MipsArgs struct {
	MipsQemu string
	MStr string			//对应-M参数, 架构所属的参数
	KernelStr string	//分为启动32位的内核还是启动64位的内核
	Nographic bool
	InitrdStr string
	HdaStr string		//启动squeezy版本硬盘还是wheezy版本的硬盘
	AppendStr string
	NetStr string   //该参数可由MipsArgs中的MacAddr和IP推到而来
	MacAddr string	//Mac的具体地址
	EthStr string
	VM_str string
	Feature string //32 bit or 64 bit, wheezy or squeezy
}

//mipsel
type MipselArgs struct {
	MipselQemu string
	MStr string   //对应-M参数
	KernelStr string
	Nographic bool
	InitrdStr string
	HdaStr string
	AppendStr string
	NetStr string   //该参数可由MipselArgs中的MacAddr和IP推到而来
	MacAddr string
	EthStr string
	VM_str string
	Feature string //32 bit or 64 bit, wheezy or squeezy
}

//powerpc
type PowerPCArgs struct {
	PowerPCQemu string
	MStr string   //对应-M参数
	KernelStr string
	Nographic bool
	InitrdStr string
	HdaStr string
	AppendStr string
	NetStr string   //该参数可由PowerPCArgs中的MacAddr和IP推到而来
	MacAddr string
	EthStr string
	VM_str string
	Feature bool //32 bit or 64 bit, wheezy or squeezy
}

//sh4
type Sh4Args struct {
	Sh4Qemu string
	MStr string   //对应-M参数
	KernelStr string
	Nographic bool
	InitrdStr string
	HdaStr string
	AppendStr string
	NetStr string   //该参数可由Sh4Args中的MacAddr和IP推到而来
	MacAddr string
	EthStr string
	VM_str string
	Feature string //32 bit or 64 bit, wheezy or squeezy
}

//sparc
type SparcArgs struct {
	MipsQemu string
	MStr string   //对应-M参数
	KernelStr string
	Nographic bool
	InitrdStr string
	HdaStr string
	AppendStr string
	NetStr string   //该参数可由SparcArgs中的MacAddr和IP推到而来
	MacAddr string
	EthStr string
	VM_str string
	Feature string //32 bit or 64 bit, wheezy or squeezy
}

*/