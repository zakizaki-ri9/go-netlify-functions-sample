package main

import (
	"encoding/json"
	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

type Result struct {
	IntValue    int    `json:"int_value"`
	StringValue string `json:"string_value"`
}

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	// リクエスト情報をログ出力
	log.Print(request.Body)

	// 返却用のjson情報を生成
	result := Result{1, "Hello"}

	// byte配列に変換
	result_byte, _ := json.Marshal(result)

	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       string(result_byte),
	}, nil
}

func main() {
	lambda.Start(handler)
}
