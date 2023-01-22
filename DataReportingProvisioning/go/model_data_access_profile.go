/*
 * 3gpp-data-reporting-provisioning
 *
 * API for 3GPP Data Reporting and Provisioning.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package DataReportingProvisioning

// DataAccessProfile - A data access profile.
type DataAccessProfile struct {

	DataAccessProfileId string `json:"dataAccessProfileId"`

	TargetEventConsumerTypes []EventConsumerType `json:"targetEventConsumerTypes"`

	Parameters []string `json:"parameters"`

	TimeAccessRestrictions DataAccessProfileTimeAccessRestrictions `json:"timeAccessRestrictions,omitempty"`

	UserAccessRestrictions DataAccessProfileUserAccessRestrictions `json:"userAccessRestrictions,omitempty"`

	LocationAccessRestrictions DataAccessProfileLocationAccessRestrictions `json:"locationAccessRestrictions,omitempty"`
}