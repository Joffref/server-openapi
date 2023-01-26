/*
 * Neasdf_BaselineDNSPattern
 *
 * EASDF Baseline DNS Pattern Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_Neasdf_BaselineDNSPattern

type PatchOperationAnyOf string

// List of PatchOperationAnyOf
const (
	ADD PatchOperationAnyOf = "add"
	COPY PatchOperationAnyOf = "copy"
	MOVE PatchOperationAnyOf = "move"
	REMOVE PatchOperationAnyOf = "remove"
	REPLACE PatchOperationAnyOf = "replace"
	TEST PatchOperationAnyOf = "test"
)
