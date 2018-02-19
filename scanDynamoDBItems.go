package main

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/aws"
	"fmt"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"strings"
)

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
		//fmt.Println("this beach is")
		//fmt.Println(test.Items[i]["beach"].S)
		var test2 *dynamodb.AttributeValue = test.Items[i]["beach"]

		//fmt.Println("pointers before and after")
		//fmt.Println(test2.S)

		// retrieves value held at the pointer
		//fmt.Println(*test2.S)

		if strings.Compare(*test2.S, "Exmouth") == 0 {
			//*beaches = append(*beaches, "Exmouth")
			beaches[0] = "test"
		}
	}
}