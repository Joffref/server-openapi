/*
 * Npcf_SMPolicyControl API
 *
 * Session Management Policy Control Service   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.2.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_Npcf_SMPolicyControl

type SteerModeValueAnyOf string

// List of SteerModeValueAnyOf
const (
	ACTIVE_STANDBY SteerModeValueAnyOf = "ACTIVE_STANDBY"
	LOAD_BALANCING SteerModeValueAnyOf = "LOAD_BALANCING"
	SMALLEST_DELAY SteerModeValueAnyOf = "SMALLEST_DELAY"
	PRIORITY_BASED SteerModeValueAnyOf = "PRIORITY_BASED"
)
