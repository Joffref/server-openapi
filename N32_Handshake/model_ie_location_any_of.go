/*
 * N32 Handshake API
 *
 * N32-c Handshake Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.3.0-alpha.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_N32_Handshake

type IeLocationAnyOf string

// List of IeLocationAnyOf
const (
	URI_PARAM IeLocationAnyOf = "URI_PARAM"
	HEADER IeLocationAnyOf = "HEADER"
	BODY IeLocationAnyOf = "BODY"
	MULTIPART_BINARY IeLocationAnyOf = "MULTIPART_BINARY"
)