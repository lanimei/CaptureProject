package main

import (
	"log"
	"os"
	"os/exec"
	"fmt"
	"time"
	"./CapturePacket"

	kingpin "gopkg.in/alecthomas/kingpin.v2"
	"bufio"
)

const APP_VERSION = "0.1"

var (
	app *kingpin.Application
	service *CapturePacket.ServiceItem
	cmd     *exec.Cmd
)


func initConfig()(err error){
	if len(os.Args) < 1{
		os.Exit(0)
	}
	//define some struct
	deviceArgs := CapturePacket.DeviceArgs{}
	analysisPacketArgs := CapturePacket.AnalysisPacketArgs{}
	malwareClassifyArgs := CapturePacket.MalwareArchClassify{}

	app := kingpin.New("CApacket", "capture and analysis the packet of the malware!\n")
	app.Author("lanimei").Version(APP_VERSION)

	//There is a prob
	daemon := app.Flag("daemon", "run this capture-analysis app in background").Default("false").Bool()
	forever := app.Flag("forever", "run this capture-analysis app in forever, fail and retry").Default("false").Bool()
	logfile := app.Flag("log", "log file path").Default("").String()
	//##########classify the malware to different arch malware
	//retain
	////////

	//capture the packet of the network adapter
	capturePacket := app.Command("capturePacket", "Capture the packet of the specificed network adapter")
	deviceArgs.DeviceName = capturePacket.Flag("device", "specify the name of network adapter,  such as: eth1").Default("eth1").Short('N').String()
	deviceArgs.FilterString = capturePacket.Flag("filter", "specify the filter string to capture the specified packet").Default("").Short('F').String()
	deviceArgs.IPName = capturePacket.Flag("ipName", "just input ip string").Default("127.0.0.1").Short('I').String()
	deviceArgs.MalName = capturePacket.Flag("malware", "just input malware name or hash").Default("").Short('M').String()
	deviceArgs.PcapPath = capturePacket.Flag("pcapPath", "specify the path of pcap packet").Default("").Short('P').String()

	//Analysis the packet
	analysisPacket := app.Command("analysisPacket", "analysis the packets that have beed captured")
	analysisPacketArgs.IsDirectory = analysisPacket.Flag("isDir", "Analysis lots of packets or just analysis one packet").Short('S').Default("false").Bool()
	analysisPacketArgs.PacketPath = analysisPacket.Flag("pkgPath", "Just analysis one packet or lots of packets. The result depends on the last param").Short('A').Default("").String()
	analysisPacketArgs.TimeStamp = analysisPacket.Flag("time", "Specify the time").Short('T').Default("").String()

	//Classify the malware
	malwareClassify := app.Command("malwareClassify", "Classify the malwares that have been captured!")
	malwareClassifyArgs.SpecifiedMalArchPath = malwareClassify.Flag("archPath", "Put the malware in this path!").Default("").Short('A').String()
	malwareClassifyArgs.SpecifiedMalPath = malwareClassify.Flag("malPath", "").Default("The position of malware").Short('M').String()

	serviceName := kingpin.MustParse(app.Parse(os.Args[1:]))

	if *logfile != ""{
		f, e := os.OpenFile(*logfile, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0600)
		if e != nil {
			log.Fatal(e)
		}
		log.SetFlags(log.Lshortfile | log.LstdFlags)
		log.SetOutput(f)
	}
	log.Println("log")
	if *daemon {
		args := []string{}
		for _, arg := range os.Args[1:] {
			if arg != "--daemon" {
				args = append(args, arg)
			}
		}
		cmd = exec.Command(os.Args[0], args...)
		cmd.Start()
		f := ""
		if *forever {
			f = "forever "
		}
		log.Printf("%s%s [PID] %d running...\n", f, os.Args[0], cmd.Process.Pid)
		os.Exit(0)
	}
	log.Println("daemon")
	if *forever {
		args := []string{}
		for _, arg := range os.Args[1:] {
			if arg != "--forever" {
				args = append(args, arg)
			}
		}
		go func() {
			for {
				if cmd != nil {
					cmd.Process.Kill()
				}
				cmd = exec.Command(os.Args[0], args...)
				cmdReaderStderr, err := cmd.StderrPipe()
				if err != nil {
					log.Printf("ERR:%s,restarting...\n", err)
					continue
				}
				cmdReader, err := cmd.StdoutPipe()
				if err != nil {
					log.Printf("ERR:%s,restarting...\n", err)
					continue
				}
				scanner := bufio.NewScanner(cmdReader)
				scannerStdErr := bufio.NewScanner(cmdReaderStderr)
				go func() {
					for scanner.Scan() {
						fmt.Println(scanner.Text())
					}
				}()
				go func() {
					for scannerStdErr.Scan() {
						fmt.Println(scannerStdErr.Text())
					}
				}()
				if err := cmd.Start(); err != nil {
					log.Printf("ERR:%s,restarting...\n", err)
					continue
				}
				pid := cmd.Process.Pid
				log.Printf("worker %s [PID] %d running...\n", os.Args[0], pid)
				if err := cmd.Wait(); err != nil {
					log.Printf("ERR:%s,restarting...", err)
					continue
				}
				log.Printf("worker %s [PID] %d unexpected exited, restarting...\n", os.Args[0], pid)
				time.Sleep(time.Second * 5)
			}
		}()
		return
	}
//	log.Println("forever")
//	log.Println(lanimei)
//	log.Println(*deviceArgs.DeviceName, "  ", *deviceArgs.FilterString)
	//Register the params to this program
	CapturePacket.Regist("capturePacket", CapturePacket.NewCapturePacket(), deviceArgs)  //service register, CapturePacketName
	log.Println(serviceName)
	log.Println(*deviceArgs.PcapPath)
	service, err = CapturePacket.Run(serviceName)

	//end the classify malware

	return nil
}