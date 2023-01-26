/*
 * Nmfaf_3caDataManagement
 *
 * MFAF 3GPP Consumer Adaptor (3CA) Data Management Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_Nmfaf_3caDataManagement

// SnssaiTaiMapping - List of restricted or unrestricted S-NSSAIs per TAI(s)
type SnssaiTaiMapping struct {

	ReportingArea TargetArea `json:"reportingArea"`

	AccessTypeList []AccessType `json:"accessTypeList,omitempty"`

	SupportedSnssaiList []SupportedSnssai `json:"supportedSnssaiList,omitempty"`
}
