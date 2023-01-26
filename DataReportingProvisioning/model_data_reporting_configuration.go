/*
 * 3gpp-data-reporting-provisioning
 *
 * API for 3GPP Data Reporting and Provisioning.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_DataReportingProvisioning

// DataReportingConfiguration - A Data Reporting Configuration subresource.
type DataReportingConfiguration struct {

	// String chosen by the 5GMS AF to serve as an identifier in a resource URI.
	DataReportingConfigurationId string `json:"dataReportingConfigurationId"`

	DataCollectionClientType DataCollectionClientType `json:"dataCollectionClientType"`

	// Uniform Resource Locator, comforming with the URI Generic Syntax specified in IETF RFC 3986.
	AuthorizationURL string `json:"authorizationURL,omitempty"`

	DataAccessProfiles []DataAccessProfile `json:"dataAccessProfiles"`
}
