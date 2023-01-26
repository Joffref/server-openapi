/*
 * 3gpp-chargeable-party
 *
 * API for Chargeable Party management.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.2.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_ChargeableParty

type EventAnyOf string

// List of EventAnyOf
const (
	SESSION_TERMINATION EventAnyOf = "SESSION_TERMINATION"
	LOSS_OF_BEARER EventAnyOf = "LOSS_OF_BEARER"
	RECOVERY_OF_BEARER EventAnyOf = "RECOVERY_OF_BEARER"
	RELEASE_OF_BEARER EventAnyOf = "RELEASE_OF_BEARER"
	USAGE_REPORT EventAnyOf = "USAGE_REPORT"
	FAILED_RESOURCES_ALLOCATION EventAnyOf = "FAILED_RESOURCES_ALLOCATION"
	SUCCESSFUL_RESOURCES_ALLOCATION EventAnyOf = "SUCCESSFUL_RESOURCES_ALLOCATION"
)
