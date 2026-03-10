package stacks

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsapigatewayv2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsapigatewayv2integrations"
	"github.com/aws/aws-cdk-go/awscdk/v2/awscertificatemanager"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsroute53"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsroute53targets"
	"github.com/aws/aws-cdk-go/awscdklambdagoalpha/v2"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
)

type ApiEndpointsStackProps struct {
	awscdk.StackProps
	HostedZoneId     string
	ZoneDomainName   string
	ApiDomainName    string
	LambdaSourcePath string
}

func NewApiEndpointsStack(scope constructs.Construct, id string, props *ApiEndpointsStackProps) awscdk.Stack {

	var sprops awscdk.StackProps
	if props != nil {
		sprops = props.StackProps
	}
	stack := awscdk.NewStack(scope, &id, &sprops)

	role := awsiam.NewRole(stack, jsii.String("ApiEndpointsLambdaRole"), &awsiam.RoleProps{
		AssumedBy: awsiam.NewServicePrincipal(jsii.String("lambda.amazonaws.com"), nil),
	})

	role.AddManagedPolicy(awsiam.ManagedPolicy_FromAwsManagedPolicyName(jsii.String("service-role/AWSLambdaBasicExecutionRole")))

	role.AddToPolicy(awsiam.NewPolicyStatement(&awsiam.PolicyStatementProps{
		Actions: &[]*string{
			jsii.String("ses:SendEmail"),
		},
		Effect: awsiam.Effect_ALLOW,
		Resources: &[]*string{
			jsii.String("*"),
		},
	}))

	lambda := awscdklambdagoalpha.NewGoFunction(stack, jsii.String("ApiEndpointsLambda"), &awscdklambdagoalpha.GoFunctionProps{
		Entry:        jsii.String(props.LambdaSourcePath),
		Architecture: awslambda.Architecture_ARM_64(),
		Environment: &map[string]*string{
			"NODE_ENV": jsii.String("production"),
		},
		Role: role,
	})

	// public hosted zone
	zone := awsroute53.HostedZone_FromHostedZoneAttributes(stack, jsii.String("MainDomain"), &awsroute53.HostedZoneAttributes{
		HostedZoneId: jsii.String(props.HostedZoneId),
		ZoneName:     jsii.String(props.ZoneDomainName),
	})

	// certificate for api
	cert := awscertificatemanager.NewCertificate(stack, jsii.String("ApiDomainNameCertificate"), &awscertificatemanager.CertificateProps{
		DomainName: jsii.String(props.ApiDomainName),
		Validation: awscertificatemanager.CertificateValidation_FromDns(zone),
	})

	// http api
	domainName := awsapigatewayv2.NewDomainName(stack, jsii.String("ApiEndpointDomainName"), &awsapigatewayv2.DomainNameProps{
		DomainName:  jsii.String(props.ApiDomainName),
		Certificate: cert,
	})

	httpApi := awsapigatewayv2.NewHttpApi(stack, jsii.String("ApiEndpointHttpApi"), &awsapigatewayv2.HttpApiProps{
		DefaultIntegration: awsapigatewayv2integrations.NewHttpLambdaIntegration(jsii.String("ApiEndpointsLambdaIntegration"), lambda, nil),
	})

	awsapigatewayv2.NewApiMapping(stack, jsii.String("ApiEndpointApiMapping"), &awsapigatewayv2.ApiMappingProps{
		Api:        httpApi,
		DomainName: domainName,
		Stage:      httpApi.DefaultStage(),
	})

	// A record alias for api
	awsroute53.NewARecord(stack, jsii.String("ApiAlias"), &awsroute53.ARecordProps{
		Zone:       zone,
		RecordName: jsii.String("api"),
		Target: awsroute53.RecordTarget_FromAlias(
			awsroute53targets.NewApiGatewayv2DomainProperties(domainName.RegionalDomainName(), domainName.RegionalHostedZoneId()),
		),
		Ttl: awscdk.Duration_Seconds(jsii.Number(600)),
	})

	return stack

}
