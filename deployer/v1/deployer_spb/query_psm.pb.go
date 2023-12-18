// Code generated by protoc-gen-go-psm. DO NOT EDIT.

package deployer_spb

import (
	psm "github.com/pentops/protostate/psm"
)

// State Query Service for %sstack
type StackPSMStateQuerySet = psm.StateQuerySet[
	*GetStackRequest,
	*GetStackResponse,
	*ListStacksRequest,
	*ListStacksResponse,
	*ListStackEventsRequest,
	*ListStackEventsResponse,
]

type StackPSMStateQuerySpec = psm.StateQuerySpec[
	*GetStackRequest,
	*GetStackResponse,
	*ListStacksRequest,
	*ListStacksResponse,
	*ListStackEventsRequest,
	*ListStackEventsResponse,
]

// State Query Service for %sdeployment
type DeploymentPSMStateQuerySet = psm.StateQuerySet[
	*GetDeploymentRequest,
	*GetDeploymentResponse,
	*ListDeploymentsRequest,
	*ListDeploymentsResponse,
	*ListDeploymentEventsRequest,
	*ListDeploymentEventsResponse,
]

type DeploymentPSMStateQuerySpec = psm.StateQuerySpec[
	*GetDeploymentRequest,
	*GetDeploymentResponse,
	*ListDeploymentsRequest,
	*ListDeploymentsResponse,
	*ListDeploymentEventsRequest,
	*ListDeploymentEventsResponse,
]
