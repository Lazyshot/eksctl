package cloudformation

import (
	"encoding/json"
)

// AWSKinesisFirehoseDeliveryStream_ElasticsearchRetryOptions AWS CloudFormation Resource (AWS::KinesisFirehose::DeliveryStream.ElasticsearchRetryOptions)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisfirehose-deliverystream-elasticsearchretryoptions.html
type AWSKinesisFirehoseDeliveryStream_ElasticsearchRetryOptions struct {

	// DurationInSeconds AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisfirehose-deliverystream-elasticsearchretryoptions.html#cfn-kinesisfirehose-deliverystream-elasticsearchretryoptions-durationinseconds
	DurationInSeconds *Value `json:"DurationInSeconds,omitempty"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSKinesisFirehoseDeliveryStream_ElasticsearchRetryOptions) AWSCloudFormationType() string {
	return "AWS::KinesisFirehose::DeliveryStream.ElasticsearchRetryOptions"
}

func (r *AWSKinesisFirehoseDeliveryStream_ElasticsearchRetryOptions) MarshalJSON() ([]byte, error) {
	return json.Marshal(*r)
}
