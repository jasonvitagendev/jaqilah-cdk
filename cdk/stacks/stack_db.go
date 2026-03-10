package stacks

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsdynamodb"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
)

type DBStackProps struct {
	awscdk.StackProps
}

func NewDBStack(scope constructs.Construct, id string, props *DBStackProps) awscdk.Stack {
	var sprops awscdk.StackProps
	if props != nil {
		sprops = props.StackProps
	}
	stack := awscdk.NewStack(scope, &id, &sprops)

	// create a dynamodb table - BeachyUserAccount
	awsdynamodb.NewTableV2(stack, jsii.String("UserAccount"), &awsdynamodb.TablePropsV2{
		TableName: jsii.String("BeachyUserAccount"),
		PartitionKey: &awsdynamodb.Attribute{
			Name: jsii.String("userId"),
			Type: awsdynamodb.AttributeType_STRING,
		},
		GlobalSecondaryIndexes: &[]*awsdynamodb.GlobalSecondaryIndexPropsV2{
			{
				IndexName: jsii.String("ByEmail"),
				PartitionKey: &awsdynamodb.Attribute{
					Name: jsii.String("email"),
					Type: awsdynamodb.AttributeType_STRING,
				},
			},
			{
				IndexName: jsii.String("ByPhone"),
				PartitionKey: &awsdynamodb.Attribute{
					Name: jsii.String("phone"),
					Type: awsdynamodb.AttributeType_STRING,
				},
			},
		},
		TimeToLiveAttribute: jsii.String("ttl"),
		Billing:             awsdynamodb.Billing_OnDemand(&awsdynamodb.MaxThroughputProps{}),
		PointInTimeRecovery: jsii.Bool(true),
		Encryption:          awsdynamodb.TableEncryptionV2_DynamoOwnedKey(),
		DeletionProtection:  jsii.Bool(true),
		RemovalPolicy:       awscdk.RemovalPolicy_RETAIN,
	})

	// create a dynamodb table - BeachySubscription
	awsdynamodb.NewTableV2(stack, jsii.String("Subscription"), &awsdynamodb.TablePropsV2{
		TableName: jsii.String("BeachySubscription"),
		PartitionKey: &awsdynamodb.Attribute{
			Name: jsii.String("subscriptionId"),
			Type: awsdynamodb.AttributeType_STRING,
		},
		GlobalSecondaryIndexes: &[]*awsdynamodb.GlobalSecondaryIndexPropsV2{
			{
				IndexName: jsii.String("ByUserId"),
				PartitionKey: &awsdynamodb.Attribute{
					Name: jsii.String("userId"),
					Type: awsdynamodb.AttributeType_STRING,
				},
			},
			{
				IndexName: jsii.String("ByStripeCustomerId"),
				PartitionKey: &awsdynamodb.Attribute{
					Name: jsii.String("stripeCustomerId"),
					Type: awsdynamodb.AttributeType_STRING,
				},
			},
			{
				IndexName: jsii.String("ByStripeSubscriptionId"),
				PartitionKey: &awsdynamodb.Attribute{
					Name: jsii.String("stripeSubscriptionId"),
					Type: awsdynamodb.AttributeType_STRING,
				},
			},
		},
		TimeToLiveAttribute: jsii.String("ttl"),
		Billing:             awsdynamodb.Billing_OnDemand(&awsdynamodb.MaxThroughputProps{}),
		PointInTimeRecovery: jsii.Bool(true),
		Encryption:          awsdynamodb.TableEncryptionV2_DynamoOwnedKey(),
		DeletionProtection:  jsii.Bool(true),
		RemovalPolicy:       awscdk.RemovalPolicy_RETAIN,
	})

	return stack

}
