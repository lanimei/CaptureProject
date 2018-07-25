package CapturePacket


type DeviceArgs struct {  //不太理解的地方在于, 这里的变量为什么要定义成指针的形式
	DeviceName	*string		//关于设备device的命名, 这里需要传入的参数可以在以后再进行扩展即可
	FilterString	*string	//关于过滤抓包字符串的相关设计, 这里主要是为了之后的扩展
	IPName *string			//这里是指IP字符串的构建
	MalName *string
	PcapPath *string
}


type DeviceMassArgs struct {   //json文件的配置
	DeviceName string `json:"deviceName"`
	FilterString string `json:"filterString"`
	IPName string `json:"ipName"`
	MalName string `json:"malName"`
	PcapPath string `json:"pcapPath"`
}


type AnalysisPacketArgs struct{	//跟第一个参数一样, 也可以进行扩展
	PacketPath *string
	IsDirectory	*bool	//是否批量分析一个文件夹下的数据包
	TimeStamp	*string //目前将时间暂定为一个时间点即可
	//找什么数据, 怎么来进行寻找, 还得查看C语言中的源码来进行设计, 即在第一版程序中寻找该段代码的实现
}


type MalwareArchClassify struct {	//Arch是architecture的缩写, 这里也就是Malware的体系结构的分类
	SpecifiedMalArchPath *string	//指定存放恶意样本架构所在的目录
	SpecifiedMalPath	*string		//指定恶意样本所在的解析路径
}


