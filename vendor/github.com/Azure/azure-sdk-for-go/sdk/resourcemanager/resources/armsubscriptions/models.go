//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armsubscriptions

// AvailabilityZoneMappings - Availability zone mappings for the region
type AvailabilityZoneMappings struct {
	// READ-ONLY; The logical zone id for the availability zone
	LogicalZone *string

	// READ-ONLY; The fully qualified physical zone id of availability zone to which logical zone id is mapped to
	PhysicalZone *string
}

// AvailabilityZonePeers - List of availability zones shared by the subscriptions.
type AvailabilityZonePeers struct {
	// Details of shared availability zone.
	Peers []*Peers

	// READ-ONLY; The availabilityZone.
	AvailabilityZone *string
}

// CheckResourceNameResult - Resource Name valid if not a reserved word, does not contain a reserved word and does not start
// with a reserved word
type CheckResourceNameResult struct {
	// Name of Resource
	Name *string

	// Is the resource name Allowed or Reserved
	Status *ResourceNameStatus

	// Type of Resource
	Type *string
}

// CheckZonePeersRequest - Check zone peers request parameters.
type CheckZonePeersRequest struct {
	// The Microsoft location.
	Location *string

	// The peer Microsoft Azure subscription ID.
	SubscriptionIDs []*string
}

// CheckZonePeersResult - Result of the Check zone peers operation.
type CheckZonePeersResult struct {
	// The Availability Zones shared by the subscriptions.
	AvailabilityZonePeers []*AvailabilityZonePeers

	// the location of the subscription.
	Location *string

	// READ-ONLY; The subscription ID.
	SubscriptionID *string
}

// ErrorAdditionalInfo - The resource management error additional info.
type ErrorAdditionalInfo struct {
	// READ-ONLY; The additional info.
	Info any

	// READ-ONLY; The additional info type.
	Type *string
}

// ErrorDetail - The error detail.
type ErrorDetail struct {
	// READ-ONLY; The error additional info.
	AdditionalInfo []*ErrorAdditionalInfo

	// READ-ONLY; The error code.
	Code *string

	// READ-ONLY; The error details.
	Details []*ErrorDetail

	// READ-ONLY; The error message.
	Message *string

	// READ-ONLY; The error target.
	Target *string
}

// ErrorResponse - Common error response for all Azure Resource Manager APIs to return error details for failed operations.
// (This also follows the OData error response format.)
type ErrorResponse struct {
	// READ-ONLY; The error additional info.
	AdditionalInfo []*ErrorAdditionalInfo

	// READ-ONLY; The error code.
	Code *string

	// READ-ONLY; The error details.
	Details []*ErrorResponse

	// READ-ONLY; The error message.
	Message *string

	// READ-ONLY; The error target.
	Target *string
}

// ErrorResponseAutoGenerated - Common error response for all Azure Resource Manager APIs to return error details for failed
// operations. (This also follows the OData error response format.).
type ErrorResponseAutoGenerated struct {
	// The error object.
	Error *ErrorDetail
}

// Location information.
type Location struct {
	// The availability zone mappings for this region.
	AvailabilityZoneMappings []*AvailabilityZoneMappings

	// Metadata of the location, such as lat/long, paired region, and others.
	Metadata *LocationMetadata

	// READ-ONLY; The display name of the location.
	DisplayName *string

	// READ-ONLY; The fully qualified ID of the location. For example, /subscriptions/8d65815f-a5b6-402f-9298-045155da7d74/locations/westus.
	ID *string

	// READ-ONLY; The location name.
	Name *string

	// READ-ONLY; The display name of the location and its region.
	RegionalDisplayName *string

	// READ-ONLY; The subscription ID.
	SubscriptionID *string

	// READ-ONLY; The location type.
	Type *LocationType
}

// LocationListResult - Location list operation response.
type LocationListResult struct {
	// An array of locations.
	Value []*Location
}

// LocationMetadata - Location metadata information
type LocationMetadata struct {
	// The regions paired to this region.
	PairedRegion []*PairedRegion

	// READ-ONLY; The geography of the location.
	Geography *string

	// READ-ONLY; The geography group of the location.
	GeographyGroup *string

	// READ-ONLY; The home location of an edge zone.
	HomeLocation *string

	// READ-ONLY; The latitude of the location.
	Latitude *string

	// READ-ONLY; The longitude of the location.
	Longitude *string

	// READ-ONLY; The physical location of the Azure location.
	PhysicalLocation *string

	// READ-ONLY; The category of the region.
	RegionCategory *RegionCategory

	// READ-ONLY; The type of the region.
	RegionType *RegionType
}

// ManagedByTenant - Information about a tenant managing the subscription.
type ManagedByTenant struct {
	// READ-ONLY; The tenant ID of the managing tenant. This is a GUID.
	TenantID *string
}

// Operation - Details of a REST API operation, returned from the Resource Provider Operations API
type Operation struct {
	// Localized display information for this particular operation.
	Display *OperationDisplay

	// READ-ONLY; Enum. Indicates the action type. "Internal" refers to actions that are for internal only APIs.
	ActionType *ActionType

	// READ-ONLY; Whether the operation applies to data-plane. This is "true" for data-plane operations and "false" for ARM/control-plane
	// operations.
	IsDataAction *bool

	// READ-ONLY; The name of the operation, as per Resource-Based Access Control (RBAC). Examples: "Microsoft.Compute/virtualMachines/write",
	// "Microsoft.Compute/virtualMachines/capture/action"
	Name *string

	// READ-ONLY; The intended executor of the operation; as in Resource Based Access Control (RBAC) and audit logs UX. Default
	// value is "user,system"
	Origin *Origin
}

// OperationAutoGenerated - Details of a REST API operation, returned from the Resource Provider Operations API
type OperationAutoGenerated struct {
	// Localized display information for this particular operation.
	Display *OperationDisplayAutoGenerated

	// Operation name: {provider}/{resource}/{operation}
	Name *string

	// READ-ONLY; Enum. Indicates the action type. "Internal" refers to actions that are for internal only APIs.
	ActionType *ActionType

	// READ-ONLY; Whether the operation applies to data-plane. This is "true" for data-plane operations and "false" for ARM/control-plane
	// operations.
	IsDataAction *bool

	// READ-ONLY; The intended executor of the operation; as in Resource Based Access Control (RBAC) and audit logs UX. Default
	// value is "user,system"
	Origin *Origin
}

// OperationDisplay - Localized display information for this particular operation.
type OperationDisplay struct {
	// READ-ONLY; The short, localized friendly description of the operation; suitable for tool tips and detailed views.
	Description *string

	// READ-ONLY; The concise, localized friendly name for the operation; suitable for dropdowns. E.g. "Create or Update Virtual
	// Machine", "Restart Virtual Machine".
	Operation *string

	// READ-ONLY; The localized friendly form of the resource provider name, e.g. "Microsoft Monitoring Insights" or "Microsoft
	// Compute".
	Provider *string

	// READ-ONLY; The localized friendly name of the resource type related to this operation. E.g. "Virtual Machines" or "Job
	// Schedule Collections".
	Resource *string
}

// OperationDisplayAutoGenerated - Localized display information for this particular operation.
type OperationDisplayAutoGenerated struct {
	// Description of the operation.
	Description *string

	// Operation type: Read, write, delete, etc.
	Operation *string

	// Service provider: Microsoft.Resources
	Provider *string

	// Resource on which the operation is performed: Profile, endpoint, etc.
	Resource *string
}

// OperationListResult - A list of REST API operations supported by an Azure Resource Provider. It contains an URL link to
// get the next set of results.
type OperationListResult struct {
	// READ-ONLY; URL to get the next set of operation list results (if there are any).
	NextLink *string

	// READ-ONLY; List of operations supported by the resource provider
	Value []*Operation
}

// OperationListResultAutoGenerated - A list of REST API operations supported by an Azure Resource Provider. It contains an
// URL link to get the next set of results.
type OperationListResultAutoGenerated struct {
	// URL to get the next set of operation list results (if there are any).
	NextLink *string

	// List of operations supported by the resource provider
	Value []*OperationAutoGenerated
}

// PairedRegion - Information regarding paired region.
type PairedRegion struct {
	// READ-ONLY; The fully qualified ID of the location. For example, /subscriptions/8d65815f-a5b6-402f-9298-045155da7d74/locations/westus.
	ID *string

	// READ-ONLY; The name of the paired region.
	Name *string

	// READ-ONLY; The subscription ID.
	SubscriptionID *string
}

// Peers - Information about shared availability zone.
type Peers struct {
	// READ-ONLY; The availabilityZone.
	AvailabilityZone *string

	// READ-ONLY; The subscription ID.
	SubscriptionID *string
}

// ResourceName - Name and Type of the Resource
type ResourceName struct {
	// REQUIRED; Name of the resource
	Name *string

	// REQUIRED; The type of the resource
	Type *string
}

// Subscription information.
type Subscription struct {
	// The authorization source of the request. Valid values are one or more combinations of Legacy, RoleBased, Bypassed, Direct
	// and Management. For example, 'Legacy, RoleBased'.
	AuthorizationSource *string

	// An array containing the tenants managing the subscription.
	ManagedByTenants []*ManagedByTenant

	// The subscription policies.
	SubscriptionPolicies *SubscriptionPolicies

	// The tags attached to the subscription.
	Tags map[string]*string

	// READ-ONLY; The subscription display name.
	DisplayName *string

	// READ-ONLY; The fully qualified ID for the subscription. For example, /subscriptions/8d65815f-a5b6-402f-9298-045155da7d74
	ID *string

	// READ-ONLY; The subscription state. Possible values are Enabled, Warned, PastDue, Disabled, and Deleted.
	State *SubscriptionState

	// READ-ONLY; The subscription ID.
	SubscriptionID *string

	// READ-ONLY; The subscription tenant ID.
	TenantID *string
}

// SubscriptionListResult - Subscription list operation response.
type SubscriptionListResult struct {
	// REQUIRED; The URL to get the next set of results.
	NextLink *string

	// An array of subscriptions.
	Value []*Subscription
}

// SubscriptionPolicies - Subscription policies.
type SubscriptionPolicies struct {
	// READ-ONLY; The subscription location placement ID. The ID indicates which regions are visible for a subscription. For example,
	// a subscription with a location placement Id of Public_2014-09-01 has access to Azure
	// public regions.
	LocationPlacementID *string

	// READ-ONLY; The subscription quota ID.
	QuotaID *string

	// READ-ONLY; The subscription spending limit.
	SpendingLimit *SpendingLimit
}

// TenantIDDescription - Tenant Id information.
type TenantIDDescription struct {
	// READ-ONLY; Country/region name of the address for the tenant.
	Country *string

	// READ-ONLY; Country/region abbreviation for the tenant.
	CountryCode *string

	// READ-ONLY; The default domain for the tenant.
	DefaultDomain *string

	// READ-ONLY; The display name of the tenant.
	DisplayName *string

	// READ-ONLY; The list of domains for the tenant.
	Domains []*string

	// READ-ONLY; The fully qualified ID of the tenant. For example, /tenants/8d65815f-a5b6-402f-9298-045155da7d74
	ID *string

	// READ-ONLY; The tenant's branding logo URL. Only available for 'Home' tenant category.
	TenantBrandingLogoURL *string

	// READ-ONLY; Category of the tenant.
	TenantCategory *TenantCategory

	// READ-ONLY; The tenant ID. For example, 8d65815f-a5b6-402f-9298-045155da7d74
	TenantID *string

	// READ-ONLY; The tenant type. Only available for 'Home' tenant category.
	TenantType *string
}

// TenantListResult - Tenant Ids information.
type TenantListResult struct {
	// REQUIRED; The URL to use for getting the next set of results.
	NextLink *string

	// An array of tenants.
	Value []*TenantIDDescription
}
