package main

import (
	"fmt"
	"os"

	"github.com/IBM/ibm-cos-sdk-go/aws"
	"github.com/IBM/ibm-cos-sdk-go/aws/session"
	"github.com/IBM/ibm-cos-sdk-go/service/s3"
)

func GetObjectsUsingSDK(token string) (string, error) {

	// Create client
	sess := session.Must(session.NewSession())
	client := s3.New(sess, &aws.Config{
		Region:      aws.String("us-south"),
		Endpoint:    aws.String(os.Getenv("BUCKET_ENDPOINT")),
		Credentials: aws.NewConfig().Credentials,
	})

	// Bucket Name
	Bucket := os.Getenv("BUCKET_NAME")

	// Call Function
	Input := &s3.ListObjectsV2Input{
		Bucket: aws.String(Bucket),
	}

	l, e := client.ListObjectsV2(Input)
	fmt.Println(l)
	fmt.Println(e) // prints "<nil>"

	return "", nil
}
