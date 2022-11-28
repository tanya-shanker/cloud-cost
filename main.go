package main

import (
	"fmt"

	"github.com/tanya-shanker/cloud-cost/cos"
)

func main() {

	_, err := cos.GetObjectsUsingSDK()
	if err != nil {
		fmt.Println("Error occurred : ", err)
	}
}
