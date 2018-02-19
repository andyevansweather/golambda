package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"strings"
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

func scanDynamoDBItems(beaches *[10]string) {
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

	// Printing out all the beaches from Beaches database
	for i := 0; i < len(test.Items); i++ {
		fmt.Println("this beach is")
		fmt.Println(test.Items[i]["beach"].S)
		var test2 *dynamodb.AttributeValue = test.Items[i]["beach"]

		fmt.Println("pointers before and after")
		fmt.Println(test2.S)

		// retrieves value held at the pointer
		fmt.Println(*test2.S)

		if strings.Compare(*test2.S, "Exmouth") == 0 {
			//*beaches = append(*beaches, "Exmouth")
			beaches[0] = "test"
		}
	}
}

func Handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	fmt.Printf("test request is %s\n", request)
	fmt.Printf("entered the handler\n")

	var goodBeaches [10]string
	goodBeaches[0] = "before change"

	fmt.Printf("before all of the changes")
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