package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"reflect"
)

type response struct {
	name string
	age  int
}

type beach struct {
	beach string
}


/**
	Temp overloaded
 */
func getValuest(dynamodbvalue *string) {
	fmt.Println("get value from pointer")
	fmt.Printf(*dynamodbvalue)
}

func getValues(dynamodbvalue *string, localstorage *string) {
	fmt.Println("test inside function")
	fmt.Printf(*dynamodbvalue)
}

func Handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	fmt.Printf("test request is %s\n", request)
	fmt.Printf("entered the handler\n")

	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("eu-west-1")},
	)

	fmt.Printf("what is the error %s \n", err)
	dynamosession := dynamodb.New(sess)

	// Get items from Beaches
	input := &dynamodb.ScanInput{
		TableName: aws.String("Beaches"),
	}

	test, errorm := dynamosession.Scan(input)

	fmt.Printf("the value is %s and the error is %s \n", test.Items, errorm)

	//var testValue string = "hello"

	//var ptrTestValue *string = &testValue

	// Printing out all the beaches from Beaches database
	for i := 0; i < len(test.Items); i++ {
		fmt.Println("this beach is")
		fmt.Println(test.Items[i]["beach"].S)
		var test2 *dynamodb.AttributeValue = test.Items[i]["beach"]

		fmt.Println("pointers before and after")
		fmt.Println(test2.S)

		// retrieves value held at the pointer
		fmt.Println(*test2.S)

		// attempt at logging out values
		fmt.Println(reflect.TypeOf(test.Items[i]["beach"]))
	}

	fmt.Printf("entered the handler %s\n", input)
	return events.APIGatewayProxyResponse{
		Body:       "Hello " + request.Body + " the output is ",
		StatusCode: 200,
	}, nil
}

func main() {
	lambda.Start(Handler)
}