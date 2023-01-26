/*
 * M1_EventDataProcessingProvisioning
 *
 * 5GMS AF M1 Event Data Processing Provisioning API © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved. 
 *
 * API version: 2.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_M1_EventDataProcessingProvisioning

// DataAccessProfile - A data access profile.
type DataAccessProfile struct {

	DataAccessProfileId string `json:"dataAccessProfileId"`

	TargetEventConsumerTypes []EventConsumerType `json:"targetEventConsumerTypes"`

	Parameters []string `json:"parameters"`

	TimeAccessRestrictions DataAccessProfileTimeAccessRestrictions `json:"timeAccessRestrictions,omitempty"`

	UserAccessRestrictions DataAccessProfileUserAccessRestrictions `json:"userAccessRestrictions,omitempty"`

	LocationAccessRestrictions DataAccessProfileLocationAccessRestrictions `json:"locationAccessRestrictions,omitempty"`
}
