package main

import (
	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	// リクエスト情報をログ出力
	log.Print(request.Body)

	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       "Hello AWS Lambda and Netlify",
	}, nil
}

func main() {
	lambda.Start(handler)
}
