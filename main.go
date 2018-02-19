package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-lambda-go/events"
	"fmt"
	"encoding/json"
)

type response struct {
	name string
	age  int
}

type Beach struct {
	beach string
}


func Handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	fmt.Printf("test request is %s\n", request.Body)
	fmt.Printf("entered the handler\n")

	// Returned result
	var goodBeaches [10]string
	goodBeaches[0] = "before change\n"

	fmt.Printf("before all of the changes\n")
	fmt.Printf(goodBeaches[0])

	// Create pointer to beaches to go to database
	var goodBeachesRef *[10]string = &goodBeaches

	scanDynamoDBItems(goodBeachesRef)

	fmt.Printf("what are the suitable beaches??\n")
	fmt.Printf(goodBeaches[0])


	return events.APIGatewayProxyResponse{
		Body:       "Hello " + request.Body + " the output is ",
		StatusCode: 200,
	}, nil
}

func main() {
	lambda.Start(Handler)
}