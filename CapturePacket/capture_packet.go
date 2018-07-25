package CapturePacket

import (
	"log"
	"time"
	"os"
	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"github.com/google/gopacket/pcap"
	"github.com/google/gopacket/pcapgo"
	"../utils"
	"fmt"
)

//这一段的程序内容只是针对单个网卡进行的抓包模块， 并不涉及到之后多个网卡模块的抓包


//important
//CapturePacket结构体可以是定义的一个数据结构, 也可以是一组数据结构
//之后如果是通过配置文件参数来达到启动多个虚拟机的目的, 可以对参数进行解析, 目前暂时不需要这个功能
type CapturePacket struct{
	Cfg DeviceArgs
	//这里写成数组的形式, 每两分钟需要记录一个数据包, 这里会记录包名, 如果长时间记录，就会导致该数组越来越大，所以这里最好写成一个channel缓存数组的形式
	//两个线程分别用来记住所抓的包， 一个是写入包名， 另一个是读取包名
	PktName []string
}



var CaptureMap = map[string]CapturePacket{}  //记录的一个恶意样本名字(可以是恶意样本的hash值或者样本名字)对应的CapturePacket值, 之后进行扩展
//这里存在一个问题是, 样本养殖后, 如果整个系统重启, 则可能会导致网卡和tap名字和系统对应不上
//这里存储的也是一个样本中的所有抓包信息
//键值对k-v数据存储在数据库中

func NewCapturePacket() Service {
	if utils.Debug_ {
		log.Println("CapturePacket.NewCapturePacket")
	}
	b := &CapturePacket{
		Cfg:	DeviceArgs{},
		PktName:	[]string{},
	}
	return b
}



func (s *CapturePacket) Start(args interface{}) (err error){// 这里的服务直接开启即可
	if utils.Debug_ {
		log.Println("CapturePacket.Start")
	}
	s.Cfg = args.(DeviceArgs)  //这里就直接将参数转换成所需要的cfg参数即可
	if err = s.InitService(); err != nil {
		return
	}
	//这里开始循环抓包, 同时每一个pcap包都会存储在一个文件夹下, 同时pcap包会以时间+mal包命名的形式来对抓住的pcap包进行命名.
	//该循环抓包什么时候停止, 怎么样才能够减少资源的消耗, 这就是一个需要思考的问题
	for {
		/*
		存在的问题:
		1. 生成包名的过程中依旧会产生流量, 这样会导致无法抓到该时段的流量包
		2.
		 */
		MalName := GetTimeString(*s.Cfg.MalName)
		if MalName == "" {
			err = fmt.Errorf("输入的该恶意代码名字为空, 无法生成该恶意代码的补捉的包名")
			break
		}
		//另外启动一个线程来修改CapturePacket中pktName中的数据
		go func(MalName string) {
			if utils.Debug_ {
				log.Println("CapturePacket.Start.go1")
			}
			s.PktName = append(s.PktName, MalName)
		}(MalName)
		err = CapturePcap(*s.Cfg.DeviceName, *s.Cfg.PcapPath, MalName)
		if err != nil{
			break
		}
	}
	return
}

func (s *CapturePacket) Clean(){  //清理如上的抓包函数, 执行Clean函数时, 会将数据转存到k-v数据库中
	if utils.Debug_ {
		log.Println("CapturePacket.Clean")
	}
	s.StopService()
}


func (s *CapturePacket) InitService() (err error){
	if utils.Debug_ {
		log.Println("CapturePacket.InitService")
	}
	return nil
}


func (s *CapturePacket) AddCaptureNode()(err error){   //将节点自己加入到MapNode中, 这个函数在最后执行退出时执行即可
	if utils.Debug_ {
		log.Println("CapturePacket.AddCaptureNode")
	}
	if *s.Cfg.MalName == "" || *s.Cfg.IPName == "" || *s.Cfg.DeviceName == "" {
		err = fmt.Errorf("样本初始化未完成, 未能完成给样本赋值的功能！")
		return
	}
	if _, ok := CaptureMap[*s.Cfg.MalName]; ok {
		err = fmt.Errorf("CaptureMap中键值已存在！")
		return
	}
	CaptureMap[*s.Cfg.MalName] = *s
	return
}


/*
func (s *CapturePacket) CheckArgs() (err error){
	if debug{
		log.Println("CapturePacket.CheckArgs")
	}
	return nil
}
*/

func (s *CapturePacket) StopService() (err error) {
	if utils.Debug_ {
		log.Println("CapturePacket.StopService")
	}
	err = s.AddCaptureNode()
	//StopService的功能可以之后再进行添加
	//停止一切抓包

	return
}


func CapturePcap(deviceName string, pcapPath string, malName string)(e error){  //所有的传入参数均需要在传入之前就处理好
	if utils.Debug_ {
		log.Println("CapturePacket.CapturePcap")
	}
	var snapshotLen uint32  = 1024
	var promiscuous bool   = false
	var timeout     time.Duration = -1 * time.Second
	var packetCount int = 0
	if pcapPath[len(pcapPath)-1] != '/' {
		pcapPath = pcapPath + "/"
	}
	f, _ := os.Create(pcapPath+malName)
	w := pcapgo.NewWriter(f)
	w.WriteFileHeader(snapshotLen, layers.LinkTypeEthernet)
	defer f.Close()

	// Open the device for capturing
	handle, err := pcap.OpenLive(deviceName, int32(snapshotLen), promiscuous, timeout)
	if err != nil {
		e = fmt.Errorf("Error opening device %s: %v", deviceName, err)
		return
	}
	defer handle.Close()

	// Start processing packets
	packetSource := gopacket.NewPacketSource(handle, handle.LinkType())
	for packet := range packetSource.Packets() {
		// Process packet here
//		fmt.Println(packet)
		w.WritePacket(packet.Metadata().CaptureInfo, packet.Data())
		packetCount++

		// Only capture 100 and then stop
		if packetCount > 20000 {
			break
		}
	}
	return
}


func GetTimeString(MalName string) string{      //每隔一段时间获取一个包名, 如果传入的恶意代码名字为空, 就会直接返回失败
	if utils.Debug_ {
		log.Println("CapturePacket.GetTimeString")
	}
	if MalName == "" {
		return ""
	}
	TimeStr := time.Now().Format("_2006-01-02_15:04:05")
	TimeStr = MalName + TimeStr //这里的Malware名字是恶意代码的名字
	return TimeStr
}

