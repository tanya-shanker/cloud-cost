package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

type Bucket struct {
	Name     string   `xml:"Name" json:"Name"`
	Contents []Object `xml:"Contents" json:"Contents"`
}

type Object struct {
	Key string `xml:"Key" json:"Key"`
}

type ErrStruct struct {
	Code       string `xml:"Code" json:"Code"`
	Messgae    string `xml:"Message" json:"Message"`
	StatusCode string `xml:"httpStatusCode" json:"httpStatusCode"`
	Resource   string `xml:"Resource" json:"Resource"`
}

func getURL() string {

	BUCKET_ENDPOINT := os.Getenv("BUCKET_ENDPOINT")
	BUCKET_NAME := os.Getenv("BUCKET_NAME")

	bucketUrl := fmt.Sprintf("https://%s/%s", BUCKET_ENDPOINT, BUCKET_NAME)

	url := fmt.Sprintf("%s?%s", bucketUrl, "list-type=2")

	fmt.Println("url :", url)

	return url
}

func GetObjects(token string) (string, error) {
	req, _ := http.NewRequest("GET", getURL(), nil)

	req.Header.Add("Authorization", token)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println("Http req error :", err)
		return "", err
	}

	var readErr error
	var responseBody []byte
	if res.Body != nil {
		defer res.Body.Close()
		responseBody, readErr = ioutil.ReadAll(res.Body)
		if readErr != nil {
			err = fmt.Errorf(" ERRORMSG_READ_RESPONSE_BODY %s", readErr.Error())
			return "", err
		}

		errStruct := &ErrStruct{}
		err = xml.Unmarshal(responseBody, errStruct)
		if nil != err {
			fmt.Println("Error unmarshalling XML", err)
		}

		response, err := json.Marshal(errStruct)
		if nil != err {
			fmt.Println("Error marshalling to JSON", err)
		}

		fmt.Printf("%s\n", response)

	}

	return "", nil
}

func ObjectAnalysis(bucket Bucket) {

	for _, obj := range bucket.Contents {
		if strings.Contains(obj.Key, "cloud-cost") {

		}
	}
}

func main() {
	iam_token := os.Getenv("CLOUD_COST_TOKEN")
	fmt.Println("iam token :", iam_token)
	list, err := GetObjects(iam_token)
	if err != nil {
		fmt.Println("Error occurred : ", err)
	}
	fmt.Println("object list :", list)
}
