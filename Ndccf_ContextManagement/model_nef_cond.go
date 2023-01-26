/*
 * Ndccf_ContextManagement
 *
 * DCCF Context Management Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_Ndccf_ContextManagement

// NefCond - Subscription to a set of NF Instances (NEFs), identified by Event ID(s) provided by AF, S-NSSAI(s), AF Instance ID, Application Identifier, External Identifier, External Group Identifier, or domain name. 
type NefCond struct {

	ConditionType string `json:"conditionType"`

	AfEvents []AfEvent `json:"afEvents,omitempty"`

	SnssaiList []Snssai `json:"snssaiList,omitempty"`

	PfdData PfdData `json:"pfdData,omitempty"`

	GpsiRanges []IdentityRange `json:"gpsiRanges,omitempty"`

	ExternalGroupIdentifiersRanges []IdentityRange `json:"externalGroupIdentifiersRanges,omitempty"`

	ServedFqdnList []string `json:"servedFqdnList,omitempty"`
}
