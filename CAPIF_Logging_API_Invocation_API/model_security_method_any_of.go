/*
 * CAPIF_Logging_API_Invocation_API
 *
 * API for invocation logs.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.2.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_CAPIF_Logging_API_Invocation_API

type SecurityMethodAnyOf string

// List of SecurityMethodAnyOf
const (
	PSK SecurityMethodAnyOf = "PSK"
	PKI SecurityMethodAnyOf = "PKI"
	OAUTH SecurityMethodAnyOf = "OAUTH"
)