/*
 * Nudm_EE
 *
 * Nudm Event Exposure Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.2.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package Nudm_EE

type RevokedCauseAnyOf string

// List of RevokedCauseAnyOf
const (
	NOT_ALLOWED RevokedCauseAnyOf = "NOT_ALLOWED"
	EXCLUDED_FROM_GROUP RevokedCauseAnyOf = "EXCLUDED_FROM_GROUP"
	GPSI_REMOVED RevokedCauseAnyOf = "GPSI_REMOVED"
)