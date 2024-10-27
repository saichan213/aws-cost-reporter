package common

import (
	"time"

	"github.com/aws/aws-sdk-go-v2/service/costexplorer"
	"github.com/aws/aws-sdk-go-v2/service/sns"
)

type Date struct {
	StartDate time.Time
	EndDate   time.Time
}

type Client struct {
	CeClient  *costexplorer.Client
	SnsClient *sns.Client
}

type Message struct {
	TermMessage        string
	ServiceCostMessage string
	TotalCostMassage   string
}
