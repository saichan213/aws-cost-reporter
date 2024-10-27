package main

import (
	"aws-cost-reporter-function/common"
	"aws-cost-reporter-function/message"
	servicecosts "aws-cost-reporter-function/service-costs"
	totalcost "aws-cost-reporter-function/total-cost"
	"context"
	"fmt"
	"time"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/costexplorer"
	"github.com/aws/aws-sdk-go-v2/service/sns"
)

// func main() {
// 	client := getClient()
// 	startDate, endDate := getDate()

// 	servicecosts.GetServiceCosts(startDate, endDate, client)
// 	totalcost.GetTotalCost(startDate, endDate, client)
// }

func main() {
	lambda.Start(HandleRequest)
}

func HandleRequest() {
	client := NewClient()
	date := NewDate()
	costMassage := NewMassage(client, date)

	message.PublishMessage(client, costMassage)
}

func NewClient() *common.Client {
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		fmt.Printf("設定の読み込みエラー: %v\n", err)
		return nil
	}

	// Cost Explorerクライアントを作成
	ceClient := costexplorer.NewFromConfig(cfg)
	snsClient := sns.NewFromConfig(cfg)

	return &common.Client{
		CeClient:  ceClient,
		SnsClient: snsClient,
	}
}

func NewDate() *common.Date {
	// 現在の月の開始日と終了日を取得

	now := time.Now()
	startDate := time.Date(now.Year(), now.Month(), 1, 0, 0, 0, 0, time.UTC)
	endDate := now

	return &common.Date{
		StartDate: startDate,
		EndDate:   endDate,
	}
}

func NewMassage(client *common.Client, date *common.Date) *common.Message {
	termMassage := fmt.Sprintf("期間: %s - %s\n", date.StartDate.Format("2006-01-02"), date.EndDate.Format("2006-01-02"))
	serviceCostMessage := servicecosts.GetServiceCosts(client, date)
	totalCostMassage := totalcost.GetTotalCost(client, date)
	return &common.Message{
		TermMessage:        termMassage,
		ServiceCostMessage: serviceCostMessage,
		TotalCostMassage:   totalCostMassage,
	}
}
