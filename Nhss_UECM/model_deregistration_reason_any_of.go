/*
 * Nhss_UECM
 *
 * HSS UE Context Management.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.1.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_Nhss_UECM

type DeregistrationReasonAnyOf string

// List of DeregistrationReasonAnyOf
const (
	UE_INITIAL_AND_SINGLE_REGISTRATION DeregistrationReasonAnyOf = "UE_INITIAL_AND_SINGLE_REGISTRATION"
	UE_INITIAL_AND_DUAL_REGISTRATION DeregistrationReasonAnyOf = "UE_INITIAL_AND_DUAL_REGISTRATION"
	EPS_TO_5_GS_MOBILITY DeregistrationReasonAnyOf = "EPS_TO_5GS_MOBILITY"
)
