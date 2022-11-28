package cos

import (
	"fmt"
	"os"

	"github.com/IBM/ibm-cos-sdk-go/aws"
	"github.com/IBM/ibm-cos-sdk-go/aws/credentials/ibmiam"
	"github.com/IBM/ibm-cos-sdk-go/aws/session"
	"github.com/IBM/ibm-cos-sdk-go/service/s3"
)

const (
	authEndpoint = "https://iam.cloud.ibm.com/identity/token"
)

func GetObjectsUsingSDK() (string, error) {

	apiKey := os.Args[1]
	serviceInstanceID := os.Args[2]
	serviceEndpoint := os.Args[3]

	cred := ibmiam.NewStaticCredentials(aws.NewConfig(),
		authEndpoint, apiKey, serviceInstanceID)

	conf := aws.Config{
		Endpoint:    aws.String(serviceEndpoint),
		Credentials: cred,
	}

	sess := session.Must(session.NewSession(&conf))
	client := s3.New(sess)

	buckets, _ := client.ListBuckets(&s3.ListBucketsInput{})
	fmt.Println(buckets)

	for _, b := range buckets.Buckets {
		fmt.Println("bucket name ", *b.Name)
		obj, err := client.ListObjects(&s3.ListObjectsInput{Bucket: b.Name})
		fmt.Println(obj)
		fmt.Println(err)
	}

	return "", nil
}
