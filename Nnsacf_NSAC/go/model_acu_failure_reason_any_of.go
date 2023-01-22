/*
 * Nnsacf_NSAC
 *
 * Nnsacf_NSAC Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package Nnsacf_NSAC

type AcuFailureReasonAnyOf string

// List of AcuFailureReasonAnyOf
const (
	SLICE_NOT_FOUND AcuFailureReasonAnyOf = "SLICE_NOT_FOUND"
	EXCEED_MAX_UE_NUM AcuFailureReasonAnyOf = "EXCEED_MAX_UE_NUM"
	EXCEED_MAX_UE_NUM_3_GPP AcuFailureReasonAnyOf = "EXCEED_MAX_UE_NUM_3GPP"
	EXCEED_MAX_UE_NUM_N3_GPP AcuFailureReasonAnyOf = "EXCEED_MAX_UE_NUM_N3GPP"
	EXCEED_MAX_PDU_NUM AcuFailureReasonAnyOf = "EXCEED_MAX_PDU_NUM"
	EXCEED_MAX_PDU_NUM_3_GPP AcuFailureReasonAnyOf = "EXCEED_MAX_PDU_NUM_3GPP"
	EXCEED_MAX_PDU_NUM_N3_GPP AcuFailureReasonAnyOf = "EXCEED_MAX_PDU_NUM_N3GPP"
)