/*
 * Ndccf_ContextManagement
 *
 * DCCF Context Management Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_Ndccf_ContextManagement

// ServiceExperienceInfoPerFlow - Contains service experience information associated with a service flow.
type ServiceExperienceInfoPerFlow struct {

	SvcExprc SvcExperience `json:"svcExprc,omitempty"`

	TimeIntev TimeWindow `json:"timeIntev,omitempty"`

	// DNAI (Data network access identifier), see clause 5.6.7 of 3GPP TS 23.501.
	Dnai string `json:"dnai,omitempty"`

	IpTrafficFilter FlowInfo `json:"ipTrafficFilter,omitempty"`

	EthTrafficFilter EthFlowDescription `json:"ethTrafficFilter,omitempty"`
}
