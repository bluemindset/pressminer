package clients

import (
	"time"

	"github.com/gocolly/colly/v2"
)

// type Client interface {
// 	FetchSource()
// }
/*
For a powerful Client we use Colly https://github.com/gocolly/colly.
Colly uses Golang’s default http client as networking layer.
HTTP options can be tweaked by changing the default HTTP roundtripper.
The object gathers articles and flush them down. For simplicity we use one
object only for this process
*/

type Miner interface {
	NewClient()
	GetArticles()
	FlushArticles()
}

type Dispatcher struct {
	C        *colly.Collector
	Articles []Article
	DateFrom time.Time
	Category string
}

type Tags struct {
}

func NewClient(opt colly.CollectorOption) *colly.Collector {
	return colly.NewCollector(opt)
}

func (c *Dispatcher) GetArticles() {

}
