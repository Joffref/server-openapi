/*
 * EES UE Location Information_API
 *
 * API for EES UE Location Information.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package Eees_UELocation

type AccuracyAnyOf string

// List of AccuracyAnyOf
const (
	CGI_ECGI AccuracyAnyOf = "CGI_ECGI"
	ENODEB AccuracyAnyOf = "ENODEB"
	TA_RA AccuracyAnyOf = "TA_RA"
	PLMN AccuracyAnyOf = "PLMN"
	TWAN_ID AccuracyAnyOf = "TWAN_ID"
	GEO_AREA AccuracyAnyOf = "GEO_AREA"
	CIVIC_ADDR AccuracyAnyOf = "CIVIC_ADDR"
)