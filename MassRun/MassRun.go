package MassRun

import (
	"../CapturePacket"
	"../utils"    //
	"log"
//	"io/ioutil"
	"encoding/json"
	"io/ioutil"
)

type DeviceMassArgs struct {   //json文件的配置
	DeviceName string `json:"deviceName"`
	FilterString string `json:"filterString"`
	IPName string `json:"ipName"`
	MalName string `json:"malName"`
	PcapPath string `json:"pcapPath"`
}

//这里定义的数据结构需要是大写字幕开头, 否则会导致json函数无法进行解析.
// 在传递函数参数时, 无法继续向下继续传递参数.
type CaptureDevice struct {
	DeviceMassArgs []DeviceMassArgs `json:"deviceMassArgs"`
}

var captureMassDevice CaptureDevice
var CapturePktMassMap = map[string]*CapturePacket.CapturePacket{}   //熟练使用包导入和包导出的


func MassStart(){
	if utils.Debug_ {
		log.Println("MassRun.MassStart")
	}

}



// 这里处理的json文件中数据属于数组数据
// json可以处理健值数据,也可以处理数组型数据

func ReadJsonFile(jsonPath string)error {
	if utils.Debug_{
		log.Println("MassRun.ReadJsonFile")
	}
	bytes, err := ioutil.ReadFile(jsonPath)
	if err != nil{
		log.Println("读取json文件出现错误")
		return err
	}
	if err := json.Unmarshal([]byte(bytes), &captureMassDevice); err !=nil{
		log.Println("解析json文件出现错误")
		return err
	}
	log.Println(len(captureMassDevice.DeviceMassArgs))
	log.Println(captureMassDevice.DeviceMassArgs[5].MalName)
	return nil
}

