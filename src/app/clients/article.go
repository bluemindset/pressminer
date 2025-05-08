package clients

import (
	"time"
)

type Article struct {
	Title       string
	Description string
	Date        time.Time
	ImagePath   string
}
