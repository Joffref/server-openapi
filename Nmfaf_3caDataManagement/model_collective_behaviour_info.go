/*
 * Nmfaf_3caDataManagement
 *
 * MFAF 3GPP Consumer Adaptor (3CA) Data Management Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_Nmfaf_3caDataManagement

// CollectiveBehaviourInfo - Contains the collective behaviour information to be reported to the subscriber.
type CollectiveBehaviourInfo struct {

	ColAttrib []PerUeAttribute `json:"colAttrib"`

	// Total number of UEs that fulfil a collective within the area of interest.
	NoOfUes int32 `json:"noOfUes,omitempty"`

	AppIds []string `json:"appIds,omitempty"`

	ExtUeIds []string `json:"extUeIds,omitempty"`

	UeIds []string `json:"ueIds,omitempty"`
}