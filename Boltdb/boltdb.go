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
		b, err := tx.CreateBucketIfNotExists([]byte("PcapMal"))					//这里就相当于创建了一张数据表而已,前一个是键值，后一个是恶意样本名字
		if err != nil {
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
		b, err := tx.CreateBucketIfNotExists([]byte("PcapMal"))					//这里就相当于创建了一张数据表而已,前一个是键值，后一个是恶意样本名字
		if err != nil {
			log.Println("SaveBoltDBStruct 执行数据库创建或读取操作失败!")
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
	if utils.Debug_{
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


