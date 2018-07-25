package CapturePacket

import (
	"log"
	"runtime/debug"
	"fmt"
	"../utils"
)



//log log log; log everywhere


//readme
//In this project, we have implement some services include capture_packet, analysis_packet and malware_classify.
//You can also implement some other modules using this service

type Service interface{  //This interfaces will be implemented by other struct
	Start(args interface{})(err error)
	Clean()  //close service
}

var serviceMap = map[string]*ServiceItem{}

type ServiceItem struct{ // include service and args that service need args
	S Service
	Args interface{}
	Name string
}

func Regist(name string, s Service, args interface{}){
	if utils.Debug_ {
		log.Println("service.Register")
	}
	Stop(name)
	serviceMap[name] = &ServiceItem{  //给结构体赋值时, 每一个数据结构都必须赋一个值
		S: s,
		Args: args,
		Name: name,
	}
}

func GetService(name string) *ServiceItem{
	if utils.Debug_ {
		log.Println("service.GetService")
	}
	if s, ok := serviceMap[name]; ok && s.S != nil{
		return s
	}
	return nil
}

func Stop(name string){
	if utils.Debug_ {
		log.Println("service.Stop")
	}
	if s, ok := serviceMap[name]; ok && s.S != nil{
		s.S.Clean()
	}
}

func Run(name string, args ...interface{})(service *ServiceItem, err error){
	if utils.Debug_ {
		log.Println("service.Run")
	}
	service, ok := serviceMap[name]
	if ok {  //attention, the
		defer func() {
			e := recover()   //将错误返回至上一层进行处理
			if e != nil{
				err = fmt.Errorf("%s servcie crashed, ERR: %s\ntrace:%s", name, e, string(debug.Stack()))
			}
		}()
		if len(args) == 1 {
			err = service.S.Start(args[0])   // 这里的err由Start函数来给出, 这一点也比较重要
		} else {
			err = service.S.Start(service.Args)
		}
		if err != nil {
			err = fmt.Errorf("%s servcie fail, ERR: %s", name, err)
		}
	} else {
		err = fmt.Errorf("service %s not found", name)
	}
	return
}

