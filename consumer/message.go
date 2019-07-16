package consumer

import (
	"database/sql"

	"github.com/nsqio/go-nsq"
)

type DBIngester struct {
	topic    string
	consumer *nsq.Consumer

	//sql info
	db        *sql.DB
	tableName string

	logChan  chan *nsq.Message
	termChan chan bool
	hupChan  chan bool
}
