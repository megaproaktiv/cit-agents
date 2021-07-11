package starter_test

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/megaproaktiv/citagents/starter"
)

var vpcid ="vpc-08df7e12dc74dbe32"

func TestStartInstance(t *testing.T) {

	starter.StartInstance(aws.String("cfncustom"), &vpcid)
	
}
