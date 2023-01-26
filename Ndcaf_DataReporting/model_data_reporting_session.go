/*
 * Ndcaf_DataReporting
 *
 * Data Collection AF: Data Collection and Reporting Configuration API and Data Reporting API © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved. 
 *
 * API version: 1.1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_Ndcaf_DataReporting

import (
	"time"
)

// DataReportingSession - A representation of a Data Reporting Session.
type DataReportingSession struct {

	// String chosen by the 5GMS AF to serve as an identifier in a resource URI.
	SessionId string `json:"sessionId,omitempty"`

	// string with format 'date-time' as defined in OpenAPI.
	ValidUntil time.Time `json:"validUntil,omitempty"`

	// String providing an application identifier.
	ExternalApplicationId string `json:"externalApplicationId"`

	SupportedDomains []DataDomain `json:"supportedDomains"`

	ReportingConditions DataReportingSessionReportingConditions `json:"reportingConditions,omitempty"`
}
