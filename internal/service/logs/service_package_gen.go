// Code generated by internal/generate/servicepackages/main.go; DO NOT EDIT.

package logs

import (
	"context"

	"github.com/hashicorp/terraform-provider-aws/internal/types"
	"github.com/hashicorp/terraform-provider-aws/names"
)

type servicePackage struct{}

func (p *servicePackage) FrameworkDataSources(ctx context.Context) []*types.ServicePackageFrameworkDataSource {
	return []*types.ServicePackageFrameworkDataSource{}
}

func (p *servicePackage) FrameworkResources(ctx context.Context) []*types.ServicePackageFrameworkResource {
	return []*types.ServicePackageFrameworkResource{}
}

func (p *servicePackage) SDKDataSources(ctx context.Context) []*types.ServicePackageSDKDataSource {
	return []*types.ServicePackageSDKDataSource{
		{
			Factory:  dataSourceDataProtectionPolicyDocument,
			TypeName: "aws_cloudwatch_log_data_protection_policy_document",
		},
		{
			Factory:  dataSourceGroup,
			TypeName: "aws_cloudwatch_log_group",
		},
		{
			Factory:  dataSourceGroups,
			TypeName: "aws_cloudwatch_log_groups",
		},
	}
}

func (p *servicePackage) SDKResources(ctx context.Context) []*types.ServicePackageSDKResource {
	return []*types.ServicePackageSDKResource{
		{
			Factory:  resourceDataProtectionPolicy,
			TypeName: "aws_cloudwatch_log_data_protection_policy",
		},
		{
			Factory:  resourceDestination,
			TypeName: "aws_cloudwatch_log_destination",
			Name:     "Destination",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: "arn",
			},
		},
		{
			Factory:  resourceDestinationPolicy,
			TypeName: "aws_cloudwatch_log_destination_policy",
		},
		{
			Factory:  resourceGroup,
			TypeName: "aws_cloudwatch_log_group",
			Name:     "Log Group",
			Tags:     &types.ServicePackageResourceTags{},
		},
		{
			Factory:  resourceMetricFilter,
			TypeName: "aws_cloudwatch_log_metric_filter",
		},
		{
			Factory:  resourceResourcePolicy,
			TypeName: "aws_cloudwatch_log_resource_policy",
		},
		{
			Factory:  resourceStream,
			TypeName: "aws_cloudwatch_log_stream",
		},
		{
			Factory:  resourceSubscriptionFilter,
			TypeName: "aws_cloudwatch_log_subscription_filter",
		},
		{
			Factory:  resourceQueryDefinition,
			TypeName: "aws_cloudwatch_query_definition",
		},
	}
}

func (p *servicePackage) ServicePackageName() string {
	return names.Logs
}

var ServicePackage = &servicePackage{}
