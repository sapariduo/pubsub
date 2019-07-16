package consumer

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/nsqio/go-nsq"
)

var device string

func parseToDB() {
	device = "ruptela"
	cfg := nsq.NewConfig()
	cFlag := nsq.ConfigFlag{cfg}
	log.Println(cFlag)
	cfg.UserAgent = fmt.Sprintf("telematics/%s", device)
	cfg.MaxInFlight = 100
	hupChan := make(chan os.Signal)
	termChan := make(chan os.Signal)
	signal.Notify(hupChan, syscall.SIGHUP)
	signal.Notify(termChan, syscall.SIGINT, syscall.SIGTERM)
}
