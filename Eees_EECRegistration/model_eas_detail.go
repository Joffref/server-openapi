/*
 * Eees_EECRegistration
 *
 * API for EEC registration. © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved. 
 *
 * API version: 1.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_Eees_EECRegistration

// EasDetail - EAS details.
type EasDetail struct {

	// Application identifier of the EAS.
	EasId string `json:"easId"`

	ExpectedSvcKPIs AcServiceKpis `json:"expectedSvcKPIs,omitempty"`

	MinimumReqSvcKPIs AcServiceKpis `json:"minimumReqSvcKPIs,omitempty"`
}
