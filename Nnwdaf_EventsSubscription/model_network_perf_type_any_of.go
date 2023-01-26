/*
 * Nnwdaf_EventsSubscription
 *
 * Nnwdaf_EventsSubscription Service API.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.2.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_Nnwdaf_EventsSubscription

type NetworkPerfTypeAnyOf string

// List of NetworkPerfTypeAnyOf
const (
	GNB_ACTIVE_RATIO NetworkPerfTypeAnyOf = "GNB_ACTIVE_RATIO"
	GNB_COMPUTING_USAGE NetworkPerfTypeAnyOf = "GNB_COMPUTING_USAGE"
	GNB_MEMORY_USAGE NetworkPerfTypeAnyOf = "GNB_MEMORY_USAGE"
	GNB_DISK_USAGE NetworkPerfTypeAnyOf = "GNB_DISK_USAGE"
	NUM_OF_UE NetworkPerfTypeAnyOf = "NUM_OF_UE"
	SESS_SUCC_RATIO NetworkPerfTypeAnyOf = "SESS_SUCC_RATIO"
	HO_SUCC_RATIO NetworkPerfTypeAnyOf = "HO_SUCC_RATIO"
)
