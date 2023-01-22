/*
 * Npcf_AMPolicyControl
 *
 * Access and Mobility Policy Control Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.2.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package Npcf_AMPolicyControl

type RequestTriggerAnyOf string

// List of RequestTriggerAnyOf
const (
	LOC_CH RequestTriggerAnyOf = "LOC_CH"
	PRA_CH RequestTriggerAnyOf = "PRA_CH"
	SERV_AREA_CH RequestTriggerAnyOf = "SERV_AREA_CH"
	RFSP_CH RequestTriggerAnyOf = "RFSP_CH"
	ALLOWED_NSSAI_CH RequestTriggerAnyOf = "ALLOWED_NSSAI_CH"
	UE_AMBR_CH RequestTriggerAnyOf = "UE_AMBR_CH"
	UE_SLICE_MBR_CH RequestTriggerAnyOf = "UE_SLICE_MBR_CH"
	SMF_SELECT_CH RequestTriggerAnyOf = "SMF_SELECT_CH"
	ACCESS_TYPE_CH RequestTriggerAnyOf = "ACCESS_TYPE_CH"
	NWDAF_DATA_CH RequestTriggerAnyOf = "NWDAF_DATA_CH"
	TARGET_NSSAI RequestTriggerAnyOf = "TARGET_NSSAI"
)