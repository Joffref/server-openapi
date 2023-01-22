/*
 * 3gpp-traffic-influence
 *
 * API for AF traffic influence   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.2.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package TrafficInfluence

type AfResultStatusAnyOf string

// List of AfResultStatusAnyOf
const (
	SUCCESS AfResultStatusAnyOf = "SUCCESS"
	TEMPORARY_CONGESTION AfResultStatusAnyOf = "TEMPORARY_CONGESTION"
	RELOC_NO_ALLOWED AfResultStatusAnyOf = "RELOC_NO_ALLOWED"
	OTHER AfResultStatusAnyOf = "OTHER"
)