package consumer

import (
	"os"
	"sync"

	"github.com/nsqio/go-nsq"
)

type Discoverer struct {
	topics   map[string]*DBIngester
	hupChan  chan os.Signal
	termChan chan os.Signal
	wg       sync.WaitGroup
	cfg      *nsq.Config
}
