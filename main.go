package main

import (
	"./config"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials/ec2rolecreds"
	"github.com/aws/aws-sdk-go/aws/ec2metadata"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/hpcloud/tail"
	"io/ioutil"
	"log"
	"os"
	"reflect"
	"strconv"
	"strings"
	"time"
)

var (
	svc        *ec2.EC2
	sess       *session.Session
	instanceId string
	metadata   *ec2metadata.EC2Metadata
)

// Set instance-id.
func GetInstanceId() error {
	var err error
	instanceId, err = metadata.GetMetadata("instance-id")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	return nil
}

// Get region.
func GetRegion() (string, error) {
	region, err := metadata.Region()
	return region, err
}

// Get file size.
func GetFileSize(filepath string) (int64, error) {
	f, err := os.Open(filepath)
	defer f.Close()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(2)
	}
	fi, err := f.Stat()

	return fi.Size(), err

}

// Get postion.
func GetPostion(filepath string) int64 {
	_, err := ioutil.ReadFile(filepath)
	if err != nil {
		log.Printf(err.Error())
	}
	return 0
}

// Check file.
func IsExist(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil
}

// Return postion file size.
func ReadPostion(path string) (int64, error) {
	f, err := os.Open(path)
	if err != nil {
		log.Printf(err.Error())
	}

	bs, err := ioutil.ReadAll(f)
	if err != nil {
		log.Printf(err.Error())
	}

	val, err := strconv.ParseInt(string(bs), 0, 64)
	if err != nil {
		log.Printf(err.Error())
	}

	return val, err
}

// Create postion file.
func CreatePostionFile(postion int, filename string) error {
	val := strconv.Itoa(postion)
	err := ioutil.WriteFile(filename, []byte(val), os.ModePerm)
	if err != nil {
		return err
	}
	return nil
}

func GetOffset(conf *config.ConfToml) (int64, error) {
	if IsExist(conf.PostionFile) {
		return ReadPostion(conf.PostionFile)
	} else {
		return GetFileSize(conf.TailFile)
	}
}

// Create(Update) EC2 tag.
func CreateTags(resources string, key string, name string) {
	params := &ec2.CreateTagsInput{
		Resources: []*string{
			aws.String(resources),
		},
		Tags: []*ec2.Tag{
			{
				Key:   aws.String(key),
				Value: aws.String(name),
			},
		},
		DryRun: aws.Bool(false),
	}

	_, err := svc.CreateTags(params)
	if err != nil {
		log.Printf(err.Error())
	}
}

// Setting aws parameter.
func awsInit() {
	sess = session.New()
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

	if err != GetInstanceId() {
		fmt.Println()
	}
	svc = ec2.New(sess, conf)
}

func main() {

	awsInit()

	conf, err := config.New()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	fmt.Println(reflect.TypeOf(conf.Delay))

	offset, err := GetOffset(conf)
	if err != nil {
		fmt.Println(err.Error())
	}

	tc := tail.Config{
		Follow: true,
		Location: &tail.SeekInfo{
			Offset: offset,
			Whence: os.SEEK_SET,
		},
	}

	t, err := tail.TailFile(conf.TailFile, tc)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	postion, err := GetFileSize(conf.TailFile)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	for line := range t.Lines {
		postion += int64(len(line.Text)) + 1
		fmt.Println(line.Text)
		fmt.Println("postion :             ", postion)

		if err := CreatePostionFile(int(postion), conf.PostionFile); err != nil {
			fmt.Println("Not Created: %s [filename=%s]\n", err.Error(), conf.PostionFile)
			os.Exit(2)
		}

		if strings.Index(line.Text, conf.SearchStart) != -1 {
			time.Sleep(conf.Delay)
			CreateTags(instanceId, conf.TagName, conf.TagStartValue)
		} else if strings.Index(line.Text, conf.SearchEnd) != -1 {
			time.Sleep(conf.Delay)
			CreateTags(instanceId, conf.TagName, conf.TagEndValue)
		}
	}
}
