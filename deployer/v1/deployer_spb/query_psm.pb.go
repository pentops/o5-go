// Code generated by protoc-gen-go-psm. DO NOT EDIT.

package deployer_spb

import (
	psm "github.com/pentops/protostate/psm"
)

// State Query Service for %sdeployment
// QuerySet is the query set for the Deployment service.

type DeploymentPSMQuerySet = psm.StateQuerySet[
	*GetDeploymentRequest,
	*GetDeploymentResponse,
	*ListDeploymentsRequest,
	*ListDeploymentsResponse,
	*ListDeploymentEventsRequest,
	*ListDeploymentEventsResponse,
]

func NewDeploymentPSMQuerySet(
	smSpec psm.QuerySpec[
		*GetDeploymentRequest,
		*GetDeploymentResponse,
		*ListDeploymentsRequest,
		*ListDeploymentsResponse,
		*ListDeploymentEventsRequest,
		*ListDeploymentEventsResponse,
	],
	options psm.StateQueryOptions,
) (*DeploymentPSMQuerySet, error) {
	return psm.BuildStateQuerySet[
		*GetDeploymentRequest,
		*GetDeploymentResponse,
		*ListDeploymentsRequest,
		*ListDeploymentsResponse,
		*ListDeploymentEventsRequest,
		*ListDeploymentEventsResponse,
	](smSpec, options)
}

type DeploymentPSMQuerySpec = psm.QuerySpec[
	*GetDeploymentRequest,
	*GetDeploymentResponse,
	*ListDeploymentsRequest,
	*ListDeploymentsResponse,
	*ListDeploymentEventsRequest,
	*ListDeploymentEventsResponse,
]

func DefaultDeploymentPSMQuerySpec(tableSpec psm.StateTableSpec) DeploymentPSMQuerySpec {
	return psm.QuerySpec[
		*GetDeploymentRequest,
		*GetDeploymentResponse,
		*ListDeploymentsRequest,
		*ListDeploymentsResponse,
		*ListDeploymentEventsRequest,
		*ListDeploymentEventsResponse,
	]{
		StateTableSpec: tableSpec,
		ListRequestFilter: func(req *ListDeploymentsRequest) (map[string]interface{}, error) {
			filter := map[string]interface{}{}
			return filter, nil
		},
		ListEventsRequestFilter: func(req *ListDeploymentEventsRequest) (map[string]interface{}, error) {
			filter := map[string]interface{}{}
			filter["deployment_id"] = req.DeploymentId
			return filter, nil
		},
	}
}

// State Query Service for %sstack
// QuerySet is the query set for the Stack service.

type StackPSMQuerySet = psm.StateQuerySet[
	*GetStackRequest,
	*GetStackResponse,
	*ListStacksRequest,
	*ListStacksResponse,
	*ListStackEventsRequest,
	*ListStackEventsResponse,
]

func NewStackPSMQuerySet(
	smSpec psm.QuerySpec[
		*GetStackRequest,
		*GetStackResponse,
		*ListStacksRequest,
		*ListStacksResponse,
		*ListStackEventsRequest,
		*ListStackEventsResponse,
	],
	options psm.StateQueryOptions,
) (*StackPSMQuerySet, error) {
	return psm.BuildStateQuerySet[
		*GetStackRequest,
		*GetStackResponse,
		*ListStacksRequest,
		*ListStacksResponse,
		*ListStackEventsRequest,
		*ListStackEventsResponse,
	](smSpec, options)
}

type StackPSMQuerySpec = psm.QuerySpec[
	*GetStackRequest,
	*GetStackResponse,
	*ListStacksRequest,
	*ListStacksResponse,
	*ListStackEventsRequest,
	*ListStackEventsResponse,
]

func DefaultStackPSMQuerySpec(tableSpec psm.StateTableSpec) StackPSMQuerySpec {
	return psm.QuerySpec[
		*GetStackRequest,
		*GetStackResponse,
		*ListStacksRequest,
		*ListStacksResponse,
		*ListStackEventsRequest,
		*ListStackEventsResponse,
	]{
		StateTableSpec: tableSpec,
		ListRequestFilter: func(req *ListStacksRequest) (map[string]interface{}, error) {
			filter := map[string]interface{}{}
			return filter, nil
		},
		ListEventsRequestFilter: func(req *ListStackEventsRequest) (map[string]interface{}, error) {
			filter := map[string]interface{}{}
			filter["stack_id"] = req.StackId
			return filter, nil
		},
	}
}
