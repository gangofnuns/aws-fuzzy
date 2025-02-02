package common

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	ec2types "github.com/aws/aws-sdk-go-v2/service/ec2/types"
	nmtypes "github.com/aws/aws-sdk-go-v2/service/networkmanager/types"
)

func GetNMTag(tags []nmtypes.Tag, key string, missing string) string {
	// Get tag Name
	for _, tag := range tags {
		if aws.ToString(tag.Key) == key {
			return aws.ToString(tag.Value)
		}
	}
	return missing
}

func GetEC2Tag(tags []ec2types.Tag, key string, missing string) string {
	// Get tag Name
	for _, tag := range tags {
		if aws.ToString(tag.Key) == key {
			return aws.ToString(tag.Value)
		}
	}
	return missing
}
