package Boltdb

import (
	"testing"
	"log"
	"../CapturePacket"
)




func TestSaveBoltDB_and_ReadBoltDB(t *testing.T) {
	err_save := SaveBoltDBString("lanimei", "wocaoni")
	if err_save != nil {
		log.Fatal("存储Boltdb出现问题")
	}
	MalName, err_Read := ReadBoltDBString("lanimei")
	if err_Read != nil {
		log.Fatal("读取BoltDB出现问题", err_Read)
	}
	log.Println(MalName)
}


func TestSerial_DeserialBoltDB(t *testing.T) {
	lanimei := CapturePacket.DeviceMassArgs{
		DeviceName: "lo",
		FilterString: "Port80",
		IPName: "172.16.19.2",
		MalName: "Gafgyt",
		PcapPath: "/home/lanimei/",
	}
	lanimei_bytes, e1 := SerialBoltDB(lanimei)
	if e1 != nil {
		log.Fatal(e1)
	}
	lanimei_2, e2 :=  DeserializeBoltDB(lanimei_bytes)
	if e2 != nil {
		log.Fatal(e2)
	}
	if lanimei_2 != lanimei {
		log.Fatal("结构体struct转换出现问题")
	}
}

func TestSaveBoltDBStruct_and_ReadBoltDB(t *testing.T) {
	lanimei := CapturePacket.DeviceMassArgs{
		DeviceName: "lo",
		FilterString: "Port80",
		IPName: "172.16.19.2",
		MalName: "Gafgyt",
		PcapPath: "/home/lanimei/",
	}
	err := SaveBoltDBStruct("123456", lanimei)
	if err != nil {
		log.Fatal("SaveBoltDBStruct出现问题")
	}
	ReadBoltDBStruct, err2 := ReadBoltDBStruct("123456")
	if err2 != nil {
		log.Fatal("ReadBoltDBStruct出现问题")
	}
	log.Println(ReadBoltDBStruct)
}

func TestMalSha256DB(t *testing.T) {
	firstMalware, err := MalSha256DB("rc.local")
	if err != nil {
		log.Println(firstMalware)
	}
	err = MalSha256DBInsert("rc.local")
	if err != nil {
		log.Println(err)
	}
	firstMalware, err = MalSha256DB("rc.local")
	if err != nil {
		log.Println(firstMalware)
		log.Fatal(err)
	}
	log.Println(firstMalware)
}

