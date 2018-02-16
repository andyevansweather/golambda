package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
)

type response struct {
	name string
	age  int
}

func Handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	fmt.Printf("test request is %s\n", request)
	fmt.Printf("entered the handler\n")

	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("eu-west-1")},
	)

	fmt.Printf("what is the error %s \n", err)
	dynamosession := dynamodb.New(sess)

	// Create item in table Movies
	input := &dynamodb.ScanInput{
		TableName: aws.String("Beaches"),
	}

	test, errorm := dynamosession.Scan(input)

	fmt.Printf("the value is %s and the error is %s \n", test, errorm)

	fmt.Printf("entered the handler %s\n", input)
	return events.APIGatewayProxyResponse{
		Body:       "Hello " + request.Body + " the output is ",
		StatusCode: 200,
	}, nil
}

func main() {
	lambda.Start(Handler)
}