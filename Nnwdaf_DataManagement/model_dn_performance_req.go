/*
 * Nnwdaf_DataManagement
 *
 * Nnwdaf_DataManagement API Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_Nnwdaf_DataManagement

// DnPerformanceReq - Represents other DN performance analytics requirements.
type DnPerformanceReq struct {

	DnPerfOrderCriter DnPerfOrderingCriterion `json:"dnPerfOrderCriter,omitempty"`

	Order MatchingDirection `json:"order,omitempty"`

	ReportThresholds []ThresholdLevel `json:"reportThresholds,omitempty"`
}
