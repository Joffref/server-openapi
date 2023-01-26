/*
 * Nudm_SDM
 *
 * Nudm Subscriber Data Management Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 2.2.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_Nudm_SDM

type ReportingTriggerAnyOf string

// List of ReportingTriggerAnyOf
const (
	PERIODICAL ReportingTriggerAnyOf = "PERIODICAL"
	EVENT_A2 ReportingTriggerAnyOf = "EVENT_A2"
	EVENT_A2_PERIODIC ReportingTriggerAnyOf = "EVENT_A2_PERIODIC"
	ALL_RRM_EVENT_TRIGGERS ReportingTriggerAnyOf = "ALL_RRM_EVENT_TRIGGERS"
)
