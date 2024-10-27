package totalcost

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/aws/aws-sdk-go-v2/service/costexplorer"
	"github.com/aws/aws-sdk-go-v2/service/costexplorer/types"
	"github.com/aws/aws-sdk-go/aws"
)

func GetTotalCost(startDate time.Time, endDate time.Time, client *costexplorer.Client) {

	// Cost Explorer APIリクエストを作成
	input := &costexplorer.GetCostAndUsageInput{
		TimePeriod: &types.DateInterval{
			Start: aws.String(startDate.Format("2006-01-02")),
			End:   aws.String(endDate.Format("2006-01-02")),
		},
		Granularity: types.GranularityMonthly,
		Metrics:     []string{"UnblendedCost"},
	}

	// APIリクエストを実行
	result, err := client.GetCostAndUsage(context.TODO(), input)
	if err != nil {
		fmt.Printf("Cost Explorerリクエストエラー: %v\n", err)
		return
	}

	// 結果を表示
	if len(result.ResultsByTime) > 0 {
		cost, _ := strconv.ParseFloat(*result.ResultsByTime[0].Total["UnblendedCost"].Amount, 64)
		fmt.Printf("現在の月のコスト合計: $%.2f\n", cost)
	} else {
		fmt.Println("コストデータが見つかりません")
	}
}
