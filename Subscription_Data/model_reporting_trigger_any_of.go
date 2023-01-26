/*
 * Unified Data Repository Service API file for subscription data
 *
 * Unified Data Repository Service (subscription data).   The API version is defined in 3GPP TS 29.504.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: -
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_Subscription_Data

type ReportingTriggerAnyOf string

// List of ReportingTriggerAnyOf
const (
	PERIODICAL ReportingTriggerAnyOf = "PERIODICAL"
	EVENT_A2 ReportingTriggerAnyOf = "EVENT_A2"
	EVENT_A2_PERIODIC ReportingTriggerAnyOf = "EVENT_A2_PERIODIC"
	ALL_RRM_EVENT_TRIGGERS ReportingTriggerAnyOf = "ALL_RRM_EVENT_TRIGGERS"
)
