package servicecosts

import (
	"aws-cost-reporter-function/common"
	"context"
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/aws/aws-sdk-go-v2/service/costexplorer"
	"github.com/aws/aws-sdk-go-v2/service/costexplorer/types"
	"github.com/aws/aws-sdk-go/aws"
)

func GetServiceCosts(client *common.Client, date *common.Date) string {
	// Cost Explorer APIリクエストを作成
	input := &costexplorer.GetCostAndUsageInput{
		TimePeriod: &types.DateInterval{
			Start: aws.String(date.StartDate.Format("2006-01-02")),
			End:   aws.String(date.EndDate.Format("2006-01-02")),
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
	result, err := client.CeClient.GetCostAndUsage(context.TODO(), input)
	if err != nil {
		log.Fatalf("Cost Explorerリクエストエラー: %v", err)
	}

	// 結果を表示

	serviceCosetMessages := []string{}
	for _, group := range result.ResultsByTime[0].Groups {
		cost, _ := strconv.ParseFloat(*group.Metrics["UnblendedCost"].Amount, 64)
		if cost > 0 {
			message := fmt.Sprintf("サービス: %s, コスト: $%.2f\n", *&group.Keys[0], cost)
			serviceCosetMessages = append(serviceCosetMessages, message)
		}
	}
	serviceCosetMessage := strings.Join(serviceCosetMessages, "\n")
	return serviceCosetMessage
}
