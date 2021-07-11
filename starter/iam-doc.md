 _, err = iamsvc.CreateRole(&iam.CreateRoleInput{
        RoleName: aws.String("AgentsRole"),
        Description: aws.String("Allows EC2 instances to call AWS services on your behalf."),
        AssumeRolePolicyDocument: &iam.AssumeRolePolicyDocument{
            Version: aws.String("2012-10-17"),
            Statement: []*iam.Statement{
                &iam.Statement{
                    Effect: aws.String("Allow"),
                    Action: []*string{
                        aws.String("sts:AssumeRole"),
                    },
                    Principal: &iam.Principal{
                        Service: []*string{
                            aws.String("ec2.amazonaws.com"),
                        },
                    },
                },
            },
        },
    })
    _, err = iamsvc.CreateInstanceProfile(&iam.CreateInstanceProfileInput{
        InstanceProfileName: aws.String("AgensRole"),
    })
    _, err = iamsvc.AttachRolePolicy(&iam.AttachRolePolicyInput{
        RoleName: aws.String("AgensRole"),
        PolicyArn: aws.String("arn:aws:iam::aws:policy/AmazonSSMManagedInstanceCore"),
    })