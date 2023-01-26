/*
 * Ndccf_ContextManagement
 *
 * DCCF Context Management Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_Ndccf_ContextManagement

// WlanPerformanceReq - Represents other WLAN performance analytics requirements.
type WlanPerformanceReq struct {

	SsIds []string `json:"ssIds,omitempty"`

	BssIds []string `json:"bssIds,omitempty"`

	WlanOrderCriter WlanOrderingCriterion `json:"wlanOrderCriter,omitempty"`

	Order MatchingDirection `json:"order,omitempty"`
}
