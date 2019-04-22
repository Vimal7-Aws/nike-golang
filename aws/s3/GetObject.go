package s3

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

func GetObject(inBucket string, inKey string) {

	config := aws.Config {

		Region:      aws.String("us-west-2"),
		Credentials: credentials.NewSharedCredentials("", "default"),

	}

	sess, err := session.NewSession(&config)

	svc := s3.New(sess)

	input := s3.GetObjectInput{
		Bucket: aws.String(inBucket),
		Key:    aws.String(inKey),
	}

	result, err := svc.GetObject(&input)

	if err != nil {

		if aerr, oke := err.(awserr.Error); oke {

			switch aerr.Code() {
			case s3.ErrCodeNoSuchKey:
				fmt.Println(s3.ErrCodeNoSuchKey, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}

		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}

	fmt.Println(result)
}
