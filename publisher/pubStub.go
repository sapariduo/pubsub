package publisher

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/nsqio/go-nsq"
)

var device string
var addr string
var p *nsq.Producer
var (
	topic     = flag.String("topic", "", "NSQ topic to publish to")
	delimiter = flag.String("delimiter", "\n", "character to split input from stdin")
)

func initialize() {
	flag.Parse()
	if len(*topic) == 0 {
		log.Fatal("--topic required")
	}

	if len(*delimiter) != 1 {
		log.Fatal("--delimiter must be a single byte")
	}
	device = "ruptela"
	addr = "localhost:4151"
	dev := &device
	endpoint := &addr
	cfg := nsq.NewConfig()
	stopChan := make(chan bool)
	termChan := make(chan os.Signal, 1)
	signal.Notify(termChan, syscall.SIGINT, syscall.SIGTERM)
	cfg.UserAgent = fmt.Sprintf("listener/%s", *dev)

	producer, err := nsq.NewProducer(*endpoint, cfg)
	if err != nil {
		log.Fatalf("failed to create nsq.Producer - %s", err)
	}

	r := bufio.NewReader(os.Stdin)
	delim := (*delimiter)[0]
	go func() {
		for {
			var err error

			err = readAndPublish(r, delim, producer)

			if err != nil {
				if err != io.EOF {
					log.Fatal(err)
				}
				close(stopChan)
				break
			}
		}
	}()

	select {
	case <-termChan:
	case <-stopChan:
	}

	producer.Stop()
}

func readAndPublish(r *bufio.Reader, delim byte, producer *nsq.Producer) error {
	line, readErr := r.ReadBytes(delim)

	if len(line) > 0 {
		// trim the delimiter
		line = line[:len(line)-1]
	}

	if len(line) == 0 {
		return readErr
	}

	err := producer.Publish(*topic, line)
	if err != nil {
		return err
	}

	return readErr
}
