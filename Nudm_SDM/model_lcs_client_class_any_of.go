/*
 * Nudm_SDM
 *
 * Nudm Subscriber Data Management Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 2.2.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_Nudm_SDM

type LcsClientClassAnyOf string

// List of LcsClientClassAnyOf
const (
	BROADCAST_SERVICE LcsClientClassAnyOf = "BROADCAST_SERVICE"
	OM_IN_HPLMN LcsClientClassAnyOf = "OM_IN_HPLMN"
	OM_IN_VPLMN LcsClientClassAnyOf = "OM_IN_VPLMN"
	ANONYMOUS_LOCATION_SERVICE LcsClientClassAnyOf = "ANONYMOUS_LOCATION_SERVICE"
	SPECIFIC_SERVICE LcsClientClassAnyOf = "SPECIFIC_SERVICE"
)
