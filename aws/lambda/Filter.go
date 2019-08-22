package lambda

import (
	"context"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	_ "strings"
)

func Handler(ctx context.Context, kinesisEvent events.KinesisEvent) {
	fmt.Println("Lambda invokeeeeeeeedddd")

	for _, record := range kinesisEvent.Records {
		kinesisRecord := record.Kinesis
		dataBytes := kinesisRecord.Data
		dataText := string(dataBytes)


		fmt.Printf("%s Data = %s \n", record.EventName, dataText)
	}

}
