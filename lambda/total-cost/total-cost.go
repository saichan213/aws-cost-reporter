package totalcost

import (
	"aws-cost-reporter-function/common"
	"context"
	"fmt"
	"strconv"

	"github.com/aws/aws-sdk-go-v2/service/costexplorer"
	"github.com/aws/aws-sdk-go-v2/service/costexplorer/types"
	"github.com/aws/aws-sdk-go/aws"
)

func GetTotalCost(client *common.Client, date *common.Date) string {

	// Cost Explorer APIリクエストを作成
	input := &costexplorer.GetCostAndUsageInput{
		TimePeriod: &types.DateInterval{
			Start: aws.String(date.StartDate.Format("2006-01-02")),
			End:   aws.String(date.EndDate.Format("2006-01-02")),
		},
		Granularity: types.GranularityMonthly,
		Metrics:     []string{"UnblendedCost"},
	}

	// APIリクエストを実行
	result, err := client.CeClient.GetCostAndUsage(context.TODO(), input)
	if err != nil {
		fmt.Printf("Cost Explorerリクエストエラー: %v\n", err)
		return ""
	}

	// 結果を表示
	if len(result.ResultsByTime) > 0 {
		cost, _ := strconv.ParseFloat(*result.ResultsByTime[0].Total["UnblendedCost"].Amount, 64)
		totalCostMassage := fmt.Sprintf("現在の月のコスト合計: $%.2f\n", cost)
		return totalCostMassage
	} else {
		fmt.Println("トータルコストデータが見つかりません")
		return ""
	}
}
