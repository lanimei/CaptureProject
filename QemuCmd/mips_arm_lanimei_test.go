package QemuCmd

import(
	"testing"
	"log"
	"time"
)



func TestCreateArmelArmhfCmd(t *testing.T) {
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
	time.Sleep(10 * time.Second)
	armel.Stop32()
/*
	armhf := CreateArmelArmhfCmd(
		armhf_m,
		armhf_kernel_wheezy,
		armhf_initrd_wheezy,
		"",
		armhf_drive_wheezy,
		true,
		armhf_append_wheezy,
		"00:16:3e:00:00:02",
		"lanimei12",
		false,
	)
	log.Println(armhf.CmdQemu)
*/
}


/*
func TestCreateMipsMipselCmd(t *testing.T) {
	mips := CreateMipsMipselCmd(
		mips_m,
		mips_kernel_squeeze_32,
		mips_hda_squeeze,
		false,
		mips_append,
		"00:16:3e:00:00:03",
		"lanimei123",
		true,
	)
	log.Println(mips.CmdQemu)
	mips.Start32()
	mipsel := CreateMipsMipselCmd(
		mipsel_m,
		mipsel_kernel_squeeze_32,
		mipsel_hda_squeeze,
		false,
		"root=/dev/sda1 console=tty0",
		"00:16:3e:00:00:03",
		"lanimei123",
		false,
	)
	log.Println(mipsel.CmdQemu)
}
*/
