package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/nyaruka/phonenumbers"
)

func handler(phoneNumber string) (string, error) {
	p, err := phonenumbers.Parse(phoneNumber, "US")
	if err != nil {
		return "", err
	}

	if !phonenumbers.IsValidNumber(p) {
		return "", err
	}

	parsedPhone := strings.TrimPrefix(phonenumbers.Format(p, phonenumbers.RFC3966), "tel:")

	return parsedPhone, nil
}

func main() {
	if len(os.Getenv("AWS_EXECUTION_ENV")) == 0 {
		phoneNumber := os.Args[1]
		p, err := handler(phoneNumber)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(p)
		}
		return
	}

	lambda.Start(handler)
}
