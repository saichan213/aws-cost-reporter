package servicecosts

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/aws/aws-sdk-go-v2/service/costexplorer"
	"github.com/aws/aws-sdk-go-v2/service/costexplorer/types"
	"github.com/aws/aws-sdk-go/aws"
)

func GetServiceCosts(startDate time.Time, endDate time.Time, client *costexplorer.Client) {
	// Cost Explorer APIリクエストを作成
	input := &costexplorer.GetCostAndUsageInput{
		TimePeriod: &types.DateInterval{
			Start: aws.String(startDate.Format("2006-01-02")),
			End:   aws.String(endDate.Format("2006-01-02")),
		},
		Granularity: types.GranularityMonthly,
		Metrics:     []string{"UnblendedCost"},
		GroupBy: []types.GroupDefinition{
			{
				Type: types.GroupDefinitionTypeDimension,
				Key:  aws.String("SERVICE"),
			},
		},
	}

	// APIリクエストを実行
	result, err := client.GetCostAndUsage(context.TODO(), input)
	if err != nil {
		log.Fatalf("Cost Explorerリクエストエラー: %v", err)
	}

	// 結果を表示
	fmt.Printf("期間: %s - %s\n", *input.TimePeriod.Start, *input.TimePeriod.End)
	for _, group := range result.ResultsByTime[0].Groups {
		cost, _ := strconv.ParseFloat(*group.Metrics["UnblendedCost"].Amount, 64)
		if cost > 0 {
			fmt.Printf("サービス: %s, コスト: $%.2f\n", *&group.Keys[0], cost)
		}
	}
}
