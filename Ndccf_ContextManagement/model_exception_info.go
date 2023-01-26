/*
 * Ndccf_ContextManagement
 *
 * DCCF Context Management Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_Ndccf_ContextManagement

// ExceptionInfo - Represents the exceptions information provided by the AF.
type ExceptionInfo struct {

	IpTrafficFilter FlowInfo `json:"ipTrafficFilter,omitempty"`

	EthTrafficFilter EthFlowDescription `json:"ethTrafficFilter,omitempty"`

	Exceps []Exception `json:"exceps"`
}
