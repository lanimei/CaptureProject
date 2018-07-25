package CapturePacket

import (
	"testing"
//	"log"
)

//go test -v .

func TestGetTimeString(t *testing.T) {
	lanimei := GetTimeString("wocao")
	t.Log(lanimei)
}

/*
func TestCapturePcap(t *testing.T) {
	e := CapturePcap("lo", "lanimei123", "/home/lanimei/Project/go/src/capture_packet/CaptureProject/CapturePacket/")
	if e != nil {
		log.Fatal(e)
	}
	log.Println("运行完毕!")
}
*/