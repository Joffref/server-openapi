/*
 * Common Data Types
 *
 * Common Data Types for Service Based Interfaces.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.   
 *
 * API version: 1.4.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_CommonData

type AuthStatusAnyOf string

// List of AuthStatusAnyOf
const (
	EAP_SUCCESS AuthStatusAnyOf = "EAP_SUCCESS"
	EAP_FAILURE AuthStatusAnyOf = "EAP_FAILURE"
	PENDING AuthStatusAnyOf = "PENDING"
)
