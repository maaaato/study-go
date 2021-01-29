package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/maaaato/monitor/config"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/ec2metadata"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/hpcloud/tail"
)

var (
	svc        *ec2.EC2
	sess       *session.Session
	instanceID string
	metadata   *ec2metadata.EC2Metadata
)

// GetInstanceID Set instance-id.
func GetInstanceID() error {
	var err error
	instanceID, err = metadata.GetMetadata("instance-id")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	return nil
}

// GetRegion Get region.
func GetRegion() (string, error) {
	region, err := metadata.Region()
	return region, err
}

// GetFileSize Get file size.
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

// Getposition Get position.
func Getposition(filepath string) int64 {
	_, err := ioutil.ReadFile(filepath)
	if err != nil {
		log.Printf(err.Error())
	}
	return 0
}

// IsExist Check file.
func IsExist(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil
}

// Readposition Return position file size.
func Readposition(path string) (int64, error) {
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

// CreatePositionFile Create position file.
func CreatePositionFile(position int, filename string) error {
	val := strconv.Itoa(position)
	err := ioutil.WriteFile(filename, []byte(val), os.ModePerm)
	if err != nil {
		return err
	}
	return nil
}

// GetOffset
func GetOffset(conf *config.ConfToml) (int64, error) {
	if IsExist(conf.PositionFile) {
		return Readposition(conf.PositionFile)
	} else {
		return GetFileSize(conf.TailFile)
	}
}

// CreateTags Create(Update) EC2 tag.
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

func main() {

	//awsInit()

	conf, err := config.LoadConfig("monitor.toml")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

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

	position, err := GetFileSize(conf.TailFile)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	fmt.Println("aaaaaaaa")
	for line := range t.Lines {
		position += int64(len(line.Text)) + 1
		fmt.Println(line.Text)
		//fmt.Println("position :             ", position)

		if err := CreatePositionFile(int(position), conf.PositionFile); err != nil {
			fmt.Printf("Not Created: %s [filename=%s]\n", err.Error(), conf.PositionFile)
			os.Exit(2)
		}

		if strings.Index(line.Text, conf.SearchStart) != -1 {
			// If the SearchStart found in the line, Update the tags by TagStartValue.
			time.Sleep(conf.Delay)
			fmt.Println("start")
			//CreateTags(instanceId, conf.TagName, conf.TagStartValue)
		} else if strings.Index(line.Text, conf.SearchEnd) != -1 {
			// If the SearchEnd found in the line, Update the tags by TagEndValue.
			time.Sleep(conf.Delay)
			fmt.Println("end")
			//CreateTags(instanceId, conf.TagName, conf.TagEndValue)
		}
	}
}
