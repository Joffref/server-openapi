/*
 * Namf_Communication
 *
 * AMF Communication Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.2.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_Namf_Communication

type LocationFilterAnyOf string

// List of LocationFilterAnyOf
const (
	TAI LocationFilterAnyOf = "TAI"
	CELL_ID LocationFilterAnyOf = "CELL_ID"
	RAN_NODE LocationFilterAnyOf = "RAN_NODE"
	N3_IWF LocationFilterAnyOf = "N3IWF"
	UE_IP LocationFilterAnyOf = "UE_IP"
	UDP_PORT LocationFilterAnyOf = "UDP_PORT"
	TNAP_ID LocationFilterAnyOf = "TNAP_ID"
	GLI LocationFilterAnyOf = "GLI"
	TWAP_ID LocationFilterAnyOf = "TWAP_ID"
)