package Boltdb

import (
	"../utils"
	"../CapturePacket"
	"encoding/gob"
	"log"
	"bytes"
	"github.com/boltdb/bolt"
	"fmt"
)


var CapturePacketName string = "CapturePacketName.db"
var MalwareDB string = "MalSha256DB.db"

/*
	PcapByte, err_pcap := SerialBoltDB(PcapName)
	if err_pcap != nil {
		log.Println(PcapByte, "SaveBoltDB PcapName序列化失败")
		return err_pcap
	}
	MalByte, err_mal := SerialBoltDB(MalName)
	if err_mal != nil {
		log.Println(MalByte, "SaveBoltDB MalName序列化失败")
		return err_mal
	}
	return b.Put(PcapByte, MalByte)
*/

//初始化过程中先建立db文件，然后在db文件中新建bucket这种类似数据库的文件。

func init(){
	if utils.Debug_ {
		log.Println("BoltDB.init")
	}
	db_packet, err1 :=  bolt.Open(CapturePacketName, 0600, nil)
	if err1 != nil {
		log.Println("数据库db_packet打开失败!")
		return
	}
	db_malware, err2 := bolt.Open(MalwareDB, 0600, nil)
	if err2 != nil {
		log.Println("数据库db_malware打开失败!")
		return
	}
	err1 = db_packet.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte("PcapMal"))
		if err != nil {
			log.Println("boltdb 初始化过程中出现错误 无法创建PcapMal")
			return err
		}
		return nil
	})
	err2 = db_malware.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte("MalSha256"))
		if err != nil {
			log.Println("boltdb初始化出现错误 无法创建MalSha256")
		}
		return nil
	})
	defer func(){
		db_malware.Close()
		db_packet.Close()
	}()
}


func SaveBoltDBString(PcapName string, MalName string)(error) {// 这一个是为了方便在bolt数据库中寻找到自己所需要的包
	if utils.Debug_ {
		log.Println("BoltDB.SaveBoltDB")
	}
	db, err := bolt.Open(CapturePacketName, 0600, nil)
	if err != nil {
		log.Println("SaveBoltDBString 数据库boltdb打开失败!")
		return err
	}
	err = db.Update(func(tx *bolt.Tx) error {			//bolt中的处理事物
		b := tx.Bucket([]byte("PcapMal"))
		if b == nil {
			log.Println("SaveBoltDBString: PcapMal Bucket不存在")
			err = fmt.Errorf("PcapMal Bucket不存在")
			return err
		}
		return b.Put([]byte(PcapName), []byte(MalName))		//这里主要是可以将存储的信息改编为struct的形式， 这一点比较不错。
	})
	if err != nil {
		log.Println("SaveBoltDBString 执行数据库存储操作失败!")
		return err
	}
	defer db.Close()
	return nil
}



func ReadBoltDBString(PcapName string)(MalName string, e error){
	if utils.Debug_ {
		log.Println("BoltDB.ReadBoltDB")
	}
	db,  err := bolt.Open(CapturePacketName, 0600, nil)
	if err != nil {
		e = err
		log.Println("ReadBoltDBString 打开数据库BoltDB失败！")
		return
	}
	e = db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("PcapMal"))
		if b == nil {
			log.Println("ReadBoltDBString: PcapMal Bucket不存在")
			err = fmt.Errorf("PcapMal Bucket不存在")
			return err
		}
		v := b.Get([]byte(PcapName))
		MalName = fmt.Sprintf("%s", v)
		if MalName == ""{
			err = fmt.Errorf("ReadBoltDBString 读取的键所对应的数值为空！")
		}
		return err
	})
	if e != nil {
		log.Println("ReadBoltDBString 执行数据库查询操作失败！")
		return
	}
	defer db.Close()
	return
}

func SaveBoltDBStruct(PcapName string, DeviceMassArgs CapturePacket.DeviceMassArgs)(error){
	if utils.Debug_ {
		log.Println("BoltDB.SaveBoltDBStruct")
	}
	db, err := bolt.Open(CapturePacketName, 0600, nil)
	if err != nil {
		log.Println("SaveBoltDBStruct 数据库boltdb打开失败!")
		return err
	}
	err = db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("PcapMal"))
		if b == nil {
			log.Println("SaveBoltDBStruct: PcapMal Bucket不存在")
			err = fmt.Errorf("PcapMal Bucket不存在")
			return err
		}
		DeviceArgs, err := SerialBoltDB(DeviceMassArgs)
		if err != nil {
			log.Println("SaveBoltDBStruct 序列化出现问题")
			return err
		}
		return b.Put([]byte(PcapName), DeviceArgs)
	})
	if err != nil {
		return err
	}
	defer db.Close()
	return nil
}

func ReadBoltDBStruct(PcapName string)(DeviceMassArgs CapturePacket.DeviceMassArgs, err error){
	if utils.Debug_ {
		log.Println("BoltDB.ReadBoltDBStruct")
	}
	db, e := bolt.Open(CapturePacketName, 0600, nil)
	if e != nil {
		err = e
		log.Println("ReadBoltDBStruct 数据库boltdb打开失败!")
		return CapturePacket.DeviceMassArgs{}, err
	}
	err = db.View(func(tx *bolt.Tx)(e error) {
		b := tx.Bucket([]byte("PcapMal"))
		v := b.Get([]byte(PcapName))
		DeviceMassArgs, e = DeserializeBoltDB(v)
		if e != nil {
			log.Println("ReadBoltDBStruct 读取结构体出现问题")
		}
		return
	})
	defer db.Close()
	return
}

//这两个函数只是简单地将结构体数据转化为[]byte字节而已，并没有特别大的变化

func SerialBoltDB(NameStr CapturePacket.DeviceMassArgs)([]byte, error){
	if utils.Debug_ {
		log.Println("BoltDB.SerialBoltDB")
	}
	var result bytes.Buffer
	encoder := gob.NewEncoder(&result)
	err := encoder.Encode(NameStr)
	if err != nil {
		return nil, err
	}
	return result.Bytes(), nil
}

//解字节化数据转化为结构体的数据

func DeserializeBoltDB(d []byte)(NameStr CapturePacket.DeviceMassArgs, err error) {
	if utils.Debug_ {
		log.Println("BoltDB.DeSerializeBoltDB")
	}
	// decoder := gob.NewDecoder(d)  //思考为什么不能这么进行输入
	decoder := gob.NewDecoder(bytes.NewReader(d))
	err = decoder.Decode(&NameStr)
	if err != nil {
		return CapturePacket.DeviceMassArgs{}, err
	}
	return NameStr, err
}


//查看样本是否已经养殖过，如果已经养殖过就不用重复养殖，查看原始养殖结果即可
//false代表养殖过， 而true则代表还未进行相关养殖
//如果值为false， 未被养殖过，那么就会在数据库中增添新值，准备进行养殖。
func MalSha256DB(MalwareSha256 string)(haveFed string, err error) {
	if utils.Debug_ {
		log.Println("BoltDB.MalSha256DB")
	}
	var db *bolt.DB
	var value []byte
	db, err = bolt.Open(MalwareDB, 0666, nil)
	if err != nil {
		log.Println("MalSha256DB:打开数据库出现问题")
		return "", err
	}
	defer db.Close()
	if err = db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("MalSha256"))
		if b == nil {
			log.Println("MalSha256 Bucket 不存在")
			err = fmt.Errorf("MalSha256 Bucket不存在")
			return err
		}
		value = b.Get([]byte(MalwareSha256))
		if len(value) == 0 {
			err = fmt.Errorf("未找到key相应的value值，请插入相关值")
			return err
		}
		return nil
	}); err != nil {
		log.Println("MaSha256DB: 查询数据库出现问题")
		return "", err
	}
	return string(value), nil
}


func MalSha256DBInsert(MalwareSha256 string)(err error){
	if utils.Debug_ {
		log.Println("BoltDB.MalSha256DBInsert")
	}
	var db *bolt.DB
	db, err = bolt.Open(MalwareDB, 0666, nil)
	if err != nil {
		log.Println("MalSha256DBInsert:打开数据库出现问题")
		return err
	}
	defer db.Close()
	if err = db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("MalSha256"))
		if b == nil {
			log.Println("MalSha256 Bucket不存在")
			err = fmt.Errorf("MalSha256 Bucket不存在")
			return err
		}
		return b.Put([]byte(MalwareSha256), []byte("Yes"))
	}); err != nil {
		log.Println("MalSha256DBInsert: 插入数据库出现问题")
		return err
	}
	return nil
}

