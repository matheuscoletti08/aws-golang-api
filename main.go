package main

import (
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func main(){
	lambda.Start(handler)
}


func handler(requests events.APIGatewayProxyRequest)
	(events.APIGatewayProxyResponse, error){
		fmt.Println("request received")

		response := events.APIGatewayProxyResponse(
			Status: 200,
		)
		return response, nil
	}
