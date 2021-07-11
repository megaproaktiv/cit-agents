package starter

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"

)


const ami = "ami-07920711782de0b25"


// We use this interface to test the functions using a mocked service.
type EC2CreateInstanceAPI interface {
	AuthorizeSecurityGroupIngress(ctx context.Context,
		params *ec2.AuthorizeSecurityGroupIngressInput,
		optFns ...func(*ec2.Options)) (*ec2.AuthorizeSecurityGroupIngressOutput, error)

	RunInstances(ctx context.Context,
		params *ec2.RunInstancesInput,
		optFns ...func(*ec2.Options)) (*ec2.RunInstancesOutput, error)

	CreateTags(ctx context.Context,
		params *ec2.CreateTagsInput,
		optFns ...func(*ec2.Options)) (*ec2.CreateTagsOutput, error)

	CreateSecurityGroup(ctx context.Context,
		params *ec2.CreateSecurityGroupInput,
		optFns ...func(*ec2.Options)) (*ec2.CreateSecurityGroupOutput, error)

	DeleteSecurityGroup(ctx context.Context, 
		params *ec2.DeleteSecurityGroupInput, 
		optFns ...func(*ec2.Options)) (*ec2.DeleteSecurityGroupOutput, error)		

	DescribeSecurityGroups(ctx context.Context,
		 params *ec2.DescribeSecurityGroupsInput, 
		 optFns ...func(*ec2.Options)) (*ec2.DescribeSecurityGroupsOutput, error)	
}

func StartInstance(agent *string, vpcid *string) error {
	Logger.Info("StartInstance start")
	Logger.Info("CreateSecurityGroup")
	sg, err := DoCreateSecurityGroup(vpcid)
	if err != nil {
		Logger.Error(err)
		panic("SecurityGroup Error")
	}

	myIp,_ := GetLocalIp()

	Logger.Info("AuthorizeSecurityGroup")
	AuthorizeSecurityGroup(sg, myIp)


	Logger.Info("RunInstance")
	input := &ec2.RunInstancesInput{
		ImageId:                           aws.String(ami),
		MinCount:                          aws.Int32(1),
		MaxCount:                          aws.Int32(1),
		InstanceType:                      types.InstanceTypeT4gNano,
		InstanceInitiatedShutdownBehavior: types.ShutdownBehaviorStop,
		UserData: EncodeBase64(`#!/bin/bash
		set -xe
		cd /home/ssm-user
		aws s3 cp s3://`+
		*getCDKBucket()+
		`/cit/agents/cfncustom .
		chmod +x cfncustom
		sudo ./cfncustom &
		`),
		IamInstanceProfile: &types.IamInstanceProfileSpecification{
			Arn: aws.String("arn:aws:iam::" + *getAccount()+ ":instance-profile/EC2RoleforSSM"),
		},
		NetworkInterfaces: []types.InstanceNetworkInterfaceSpecification{
            {	
                AssociatePublicIpAddress: aws.Bool(true),
				DeviceIndex: aws.Int32(0),
                DeleteOnTermination: aws.Bool(true),
				Groups: []string{
					*sg,
				},
            },
        },
	}

	result, err := ec2Client.RunInstances(context.TODO(), input)
	if err != nil {
		fmt.Println("Got an error creating an instance:")
		fmt.Println(err)

		inputJSON, err := json.MarshalIndent(input, "", "  ")
		if err != nil {
			Logger.Error(err)
			resultJSON, _ := json.MarshalIndent(result, "", "  ")
			fmt.Println(string(resultJSON))
		}
		fmt.Printf("Parameter %s\n", string(inputJSON))
		return err
	}
	// resultJSON, _ := json.MarshalIndent(result, "", "  ")
	// fmt.Println(string(resultJSON))
	// ip := result.Instances[0].NetworkInterfaces[0].Association.PublicIp
	// Logger.Info("Public IP ", ip)
	return nil
}





func EncodeBase64(raw string) *string {
    data := []byte(raw)
    str := base64.StdEncoding.EncodeToString(data)
	return &str
}