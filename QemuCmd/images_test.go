package QemuCmd

import (
	"testing"
	"log"
)

/*
func TestMipsArgs_LoadImage(t *testing.T) {
	image_str := "/home/lanimei/Project/go/src/capture_packet/CaptureProject/QemuCmd/file/debian_squeeze_mips_lanimei123.qcow2"
	mount_str := "/home/lanimei/Project/go/src/capture_packet/CaptureProject/QemuCmd/test"
	err := LoadImage(image_str, mount_str)
	if err != nil {
		log.Fatal(err)
	}
}


func TestMipsArgs_UnloadImage(t *testing.T) {
	mount_str := "/home/lanimei/Project/go/src/capture_packet/CaptureProject/QemuCmd/test"
	err := UnloadImage(mount_str)
	if err != nil {
		log.Println(err)
	}
}

*/
/*
func TestMakedir(t *testing.T) {
	err := Makedir("/home/lanimei/Project/go/src/capture_packet/CaptureProject/QemuCmd/test")
	if err != nil {
		log.Fatal(err)
	}
}
*/

/*
func TestCpInitImage(t *testing.T) {
	err := CpInitImage("/home/lanimei/Project/go/src/capture_packet/CaptureProject/QemuCmd/file/rc.local", "/home/lanimei/Project/go/src/capture_packet/CaptureProject/QemuCmd/")
	if err != nil {
		log.Fatal(err)
	}
}
*/


func TestCreateImage(t *testing.T) {
	filesPath := FilesPath{
		MalwarePath: "/home/lanimei/Project/go/src/capture_packet/CaptureProject/QemuCmd/malware_one",
		StartScript: "/home/lanimei/Project/go/src/capture_packet/CaptureProject/QemuCmd/file/rc.local",
		ImageSrcPath:"/home/lanimei/Project/go/src/capture_packet/CaptureProject/QemuCmd/debian_squeeze_mips_standard.qcow2",
		ImageNameSpecify: "debian_squeeze_mips_lanimei123.qcow2",
		ImageDestPath: "/home/lanimei/Project/go/src/capture_packet/CaptureProject/QemuCmd/file",
		LoadPath: "/home/lanimei/Project/go/src/capture_packet/CaptureProject/QemuCmd/test",
	}
	ImageErr := CreateImage(filesPath)
	if ImageErr != nil {
		log.Fatal(ImageErr)
	}
}
