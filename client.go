package main

import (
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials/ec2rolecreds"
	"github.com/aws/aws-sdk-go/aws/ec2metadata"
	"github.com/aws/aws-sdk-go/service/ec2"
)

type Client struct {
	Service   *ec2.EC2
	InstansID string
}

// Setting aws parameter.
func newClient() {
	c := Client{}
	metadata = ec2metadata.New(sess)
	cred := ec2rolecreds.NewCredentialsWithClient(metadata)

	region, err := GetRegion()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	conf := &aws.Config{
		Credentials: cred,
		Region:      &region,
	}

	if err != GetInstanceID() {
		fmt.Println()
	}
	svc = ec2.New(sess, conf)
}
