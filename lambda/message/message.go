package message

import (
	"aws-cost-reporter-function/common"
	"context"
	"fmt"
	"log"
	"os"

	"github.com/aws/aws-sdk-go-v2/service/sns"
)

func PublishMessage(client *common.Client, message *common.Message) {
	// SNSトピックのARN
	topicARN := os.Getenv("SNS_TOPIC_ARN")

	// 送信するメッセージ
	resultMessage := fmt.Sprintf("%s\n%s\n%s", message.TermMessage, message.TotalCostMassage, message.ServiceCostMessage)

	// メッセージを発行
	input := &sns.PublishInput{
		Message:  &resultMessage,
		TopicArn: &topicARN,
	}

	result, err := client.SnsClient.Publish(context.TODO(), input)
	if err != nil {
		log.Fatalf("メッセージの発行に失敗しました: %v", err)
	}

	fmt.Printf("メッセージが正常に発行されました。MessageID: %s\n", *result.MessageId)
}
