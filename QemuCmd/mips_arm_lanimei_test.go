package QemuCmd

import(
//	"testing"
//	"log"
//	"time"
)


/*
func TestCreateArmelArmhfCmd(t *testing.T) {

	//armel开始运行
	armel := CreateArmelArmhfCmd(
		armel_m,
		armel_kernel_squeeze,
		armel_initrd_squeeze,
		armel_hda_squeeze,
		"",
		false,
		armel_append_squeeze,
		"00:16:3e:00:00:01",
		"lanimei1",
		true,
		)
	log.Println(armel.CmdQemu)
	go armel.Start32()

	//armhf 开始运行
	armhf := CreateArmelArmhfCmd(
		armhf_m,
		armhf_kernel_wheezy,
		armhf_initrd_wheezy,
		"",
		armhf_drive_wheezy,
		false,
		armhf_append_wheezy,
		"00:16:3e:00:00:02",
		"lanimei2",
		false,
	)
	log.Println(armhf.CmdQemu)
	go armhf.Start32()

	//等待20秒，即结束运行。
	time.Sleep(20 * time.Second)
	armel.Stop32()
	armhf.Stop32()

	//强制退出后， 会打印一些退出的字符串
}
*/
/*
func TestCreateMipsMipselCmd(t *testing.T) {
	time.Sleep(20 * time.Second)
	mips := CreateMipsMipselCmd(
		mips_m,
		mips_kernel_squeeze_32,
		mips_hda_squeeze,
		false,
		mips_append,
		"00:16:3e:00:00:03",
		"lanimei3",
		true,
	)
	log.Println(mips.CmdQemu)
	go mips.Start32()
	mipsel := CreateMipsMipselCmd(
		mipsel_m,
		mipsel_kernel_squeeze_32,
		mipsel_hda_squeeze,
		false,
		mipsel_append,
		"00:16:3e:00:00:04",
		"lanimei4",
		false,
	)
	log.Println(mipsel.CmdQemu)
	go mipsel.Start32()

	//等待20秒，即结束运行。
	time.Sleep(20 * time.Second)
	mips.Stop32()
	mipsel.Stop32()
}

*/