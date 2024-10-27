package main

import (
	servicecosts "aws-cost-reporter-function/service-costs"
	totalcost "aws-cost-reporter-function/total-cost"
	"context"
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/costexplorer"
)

func main() {
	client := getClient()
	startDate, endDate := getDate()

	servicecosts.GetServiceCosts(startDate, endDate, client)
	totalcost.GetTotalCost(startDate, endDate, client)
}

func getClient() *costexplorer.Client {
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		fmt.Printf("設定の読み込みエラー: %v\n", err)
		return nil
	}

	// Cost Explorerクライアントを作成
	client := costexplorer.NewFromConfig(cfg)

	return client
}

func getDate() (time.Time, time.Time) {
	// 現在の月の開始日と終了日を取得
	now := time.Now()
	startDate := time.Date(now.Year(), now.Month(), 1, 0, 0, 0, 0, time.UTC)
	endDate := now

	return startDate, endDate
}
