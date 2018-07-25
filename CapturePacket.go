package main

import(
	"log"
	"os"
	"os/signal"
	"syscall"
)


func main(){
	err := initConfig()
	if err != nil{
		log.Fatalf("err: %s", err)
	}
	SignalApp()
}

func SignalApp(){   //record the signal of INTERRUPT or SIGQUIT, this method can also receive the other signals
	signalChan := make(chan os.Signal, 1)   //Once this program receive the signal,
	cleaupDone := make(chan bool)
	signal.Notify(signalChan,	// Only when receiving these signals, this program will end.
		os.Interrupt,
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT) // capture all the signals
	go func() {
//		for _ = range signalChan{
		c := <-signalChan
		log.Println("Received an signal: ", c.String())  //print the quit singnal
		cleaupDone <-true
//		}
	}()
	<-cleaupDone
}
