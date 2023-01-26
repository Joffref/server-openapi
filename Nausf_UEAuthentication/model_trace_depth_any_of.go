/*
 * AUSF API
 *
 * AUSF UE Authentication Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.2.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_Nausf_UEAuthentication

type TraceDepthAnyOf string

// List of TraceDepthAnyOf
const (
	MINIMUM TraceDepthAnyOf = "MINIMUM"
	MEDIUM TraceDepthAnyOf = "MEDIUM"
	MAXIMUM TraceDepthAnyOf = "MAXIMUM"
	MINIMUM_WO_VENDOR_EXTENSION TraceDepthAnyOf = "MINIMUM_WO_VENDOR_EXTENSION"
	MEDIUM_WO_VENDOR_EXTENSION TraceDepthAnyOf = "MEDIUM_WO_VENDOR_EXTENSION"
	MAXIMUM_WO_VENDOR_EXTENSION TraceDepthAnyOf = "MAXIMUM_WO_VENDOR_EXTENSION"
)