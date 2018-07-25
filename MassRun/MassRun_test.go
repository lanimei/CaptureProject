package MassRun

import (
	"testing"
	"log"
)



func TestReadJsonFile(t *testing.T) {
	lanimei := ReadJsonFile("test.json")
	if lanimei != nil{
		log.Fatal(lanimei)
	}
}


/*
func Readlanimei(){
	var captureDevice CaptureDevice
	bodys := `{
    "deviceMassArgs":[{
      "deviceName": "lo",
      "filterString": "",
      "ipName": "",
      "MalName": "lanimei1",
      "pcapPath": "/home/lanimei/Project/go/src/capture_packet/CaptureProject"
    },
	{
      "deviceName": "lo",
      "filterString": "",
      "ipName": "",
      "MalName": "lanimei2",
      "pcapPath": "/home/lanimei/Project/go/src/capture_packet/CaptureProject"
    }]
	}`
	if err := json.Unmarshal([]byte(bodys), &captureDevice); err !=nil{
		log.Println("解析json文件出现错误")
		log.Println(err)
	}
	log.Println(captureDevice.DeviceMassArgs[0].MalName)
}


func TestReadlanimei(t *testing.T) {
	Readlanimei()
}
*/
