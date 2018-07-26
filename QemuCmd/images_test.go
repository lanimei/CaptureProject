package QemuCmd

import (
	"testing"
	"log"
)


func TestMipsArgs_LoadImage(t *testing.T) {
	image_str := "/home/lanimei/Project/go/src/capture_packet/CaptureProject/QemuCmd/debian_squeeze_mips_standard.qcow2"
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


func TestMakedir(t *testing.T) {
	err := Makedir("/home/lanimei/Project/go/src/capture_packet/CaptureProject/QemuCmd/test")
	if err != nil {
		log.Fatal(err)
	}
}

