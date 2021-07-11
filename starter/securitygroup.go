package starter

import (
	"context"
	"errors"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"

	"github.com/aws/smithy-go"

)

var aSGgroupName = "agent"

func DoCreateSecurityGroup(vpcid *string) (*string, error) {
	if existAgentSG() {
		deleteAgentSG()
	}
	return CreateSecurityGroup(vpcid)
}

func CreateSecurityGroup(vpcid *string) (*string, error) {
	var ae smithy.APIError
	createSG, err := ec2Client.CreateSecurityGroup(context.TODO(), &ec2.CreateSecurityGroupInput{
		Description: &aSGgroupName,
		GroupName:   &aSGgroupName,
		DryRun:      aws.Bool(false),
		TagSpecifications: []types.TagSpecification{
			{
				ResourceType: types.ResourceTypeSecurityGroup,
				Tags: []types.Tag{
					{
						Key:   aws.String("Name"),
						Value: aws.String("api"),
					},
				},
			},
		},
		VpcId: vpcid,
	})
	if err != nil {
		if errors.As(err, &ae) {
			switch ae.ErrorCode() {
			case "InvalidGroup.Duplicate":

			}
		}
		return nil, err
	}
	return createSG.GroupId, nil
}

func AuthorizeSecurityGroup(securityGroupID *string, myIp *string) error {
	_, err := ec2Client.AuthorizeSecurityGroupIngress(context.TODO(),
		&ec2.AuthorizeSecurityGroupIngressInput{
			GroupId: securityGroupID,
			IpPermissions: []types.IpPermission{
				{
					IpProtocol: aws.String("tcp"),
					FromPort:   aws.Int32(443),
					ToPort:     aws.Int32(443),
					IpRanges: []types.IpRange{
						{
							Description: aws.String("mock api port"),
							CidrIp:      myIp,
						},
					},
				},
				{
					IpProtocol: aws.String("tcp"),
					FromPort:   aws.Int32(8081),
					ToPort:     aws.Int32(8081),
					IpRanges: []types.IpRange{
						{
							Description: aws.String("mock control port"),
							CidrIp:      myIp,
						},
					},
				},
			},
		})
	return err
}

// existAgentSG - do the SG name of the agent exist
func existAgentSG() bool {
	var ae smithy.APIError

	res, err := ec2Client.DescribeSecurityGroups(context.TODO(),
		&ec2.DescribeSecurityGroupsInput{
			GroupNames: []string{ aSGgroupName},
	})	

	if err != nil {
		if errors.As(err, &ae) {
			switch ae.ErrorCode() {
			case "InvalidGroup.NotFound":
				return false
			}
		}
	}	
	
	if len(res.SecurityGroups) == 1 {
		return true
	}
	return false
}

func deleteAgentSG() error {
	Logger.Info("Deleting old SG")
	_, err := ec2Client.DeleteSecurityGroup(context.TODO(),
	&ec2.DeleteSecurityGroupInput{
		GroupName: &aSGgroupName,
	})
	if err != nil {
		Logger.Error(err)
		return err
	}
	return nil
}