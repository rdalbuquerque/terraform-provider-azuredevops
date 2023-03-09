// --------------------------------------------------------------------------------------------
// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.
// --------------------------------------------------------------------------------------------
// Generated file, DO NOT EDIT
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// --------------------------------------------------------------------------------------------

package pipelineschecks

import (
	"github.com/google/uuid"
	"github.com/microsoft/azure-devops-go-api/azuredevops"
	"github.com/microsoft/azure-devops-go-api/azuredevops/webapi"
)

type CheckConfiguration struct {
	// Check configuration id.
	Id *int `json:"id,omitempty"`
	// Resource on which check get configured.
	Resource *Resource `json:"resource,omitempty"`
	// Check configuration type
	Type *CheckType `json:"type,omitempty"`
	// The URL from which one can fetch the configured check.
	Url *string `json:"url,omitempty"`
	// Reference links.
	Links interface{} `json:"_links,omitempty"`
	// Identity of person who configured check.
	CreatedBy *webapi.IdentityRef `json:"createdBy,omitempty"`
	// Time when check got configured.
	CreatedOn *azuredevops.Time `json:"createdOn,omitempty"`
	// Identity of person who modified the configured check.
	ModifiedBy *webapi.IdentityRef `json:"modifiedBy,omitempty"`
	// Time when configured check was modified.
	ModifiedOn *azuredevops.Time `json:"modifiedOn,omitempty"`
	// Settings for the check configuration.
	Settings interface{} `json:"settings,omitempty"`
	// Timeout in minutes for the check.
	Timeout *int `json:"timeout,omitempty"`
}

type CheckConfigurationRef struct {
	// Check configuration id.
	Id *int `json:"id,omitempty"`
	// Resource on which check get configured.
	Resource *Resource `json:"resource,omitempty"`
	// Check configuration type
	Type *CheckType `json:"type,omitempty"`
	// The URL from which one can fetch the configured check.
	Url *string `json:"url,omitempty"`
}

type CheckRun struct {
	ResultMessage         *string                `json:"resultMessage,omitempty"`
	Status                *CheckRunStatus        `json:"status,omitempty"`
	CheckConfigurationRef *CheckConfigurationRef `json:"checkConfigurationRef,omitempty"`
	CompletedDate         *azuredevops.Time      `json:"completedDate,omitempty"`
	CreatedDate           *azuredevops.Time      `json:"createdDate,omitempty"`
	Id                    *uuid.UUID             `json:"id,omitempty"`
}

type CheckRunResult struct {
	ResultMessage *string         `json:"resultMessage,omitempty"`
	Status        *CheckRunStatus `json:"status,omitempty"`
}

// [Flags]
type CheckRunStatus string

type checkRunStatusValuesType struct {
	None     CheckRunStatus
	Queued   CheckRunStatus
	Running  CheckRunStatus
	Approved CheckRunStatus
	Rejected CheckRunStatus
	Canceled CheckRunStatus
}

var CheckRunStatusValues = checkRunStatusValuesType{
	None:     "none",
	Queued:   "queued",
	Running:  "running",
	Approved: "approved",
	Rejected: "rejected",
	Canceled: "canceled",
}

type CheckSuite struct {
	// Reference links.
	Links interface{} `json:"_links,omitempty"`
	// List of check runs associated with the given check suite request.
	CheckRuns *[]CheckRun `json:"checkRuns,omitempty"`
	// Completed date of the given check suite request
	CompletedDate *azuredevops.Time `json:"completedDate,omitempty"`
	// Evaluation context for the check suite request
	Context interface{} `json:"context,omitempty"`
	// Unique suite id generated by the pipeline orchestrator for the pipeline check runs request on the list of resources Pipeline orchestrator will used this identifier to map the check requests on a stage
	Id *uuid.UUID `json:"id,omitempty"`
	// Optional message for the given check suite request
	Message *string `json:"message,omitempty"`
	// Overall check runs status for the given suite request. This is check suite status
	Status *CheckRunStatus `json:"status,omitempty"`
}

type CheckSuiteRequest struct {
	Context   interface{} `json:"context,omitempty"`
	Id        *uuid.UUID  `json:"id,omitempty"`
	Resources *[]Resource `json:"resources,omitempty"`
}

type CheckType struct {
	// Gets or sets check type id.
	Id *uuid.UUID `json:"id,omitempty"`
	// Name of the check type.
	Name *string `json:"name,omitempty"`
}

type Resource struct {
	// Id of the resource.
	Id *string `json:"id,omitempty"`
	// Type of the resource.
	Type *string `json:"type,omitempty"`
}
