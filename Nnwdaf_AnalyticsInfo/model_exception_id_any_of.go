/*
 * Nnwdaf_AnalyticsInfo
 *
 * Nnwdaf_AnalyticsInfo Service API.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.2.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_Nnwdaf_AnalyticsInfo

type ExceptionIdAnyOf string

// List of ExceptionIdAnyOf
const (
	UNEXPECTED_UE_LOCATION ExceptionIdAnyOf = "UNEXPECTED_UE_LOCATION"
	UNEXPECTED_LONG_LIVE_FLOW ExceptionIdAnyOf = "UNEXPECTED_LONG_LIVE_FLOW"
	UNEXPECTED_LARGE_RATE_FLOW ExceptionIdAnyOf = "UNEXPECTED_LARGE_RATE_FLOW"
	UNEXPECTED_WAKEUP ExceptionIdAnyOf = "UNEXPECTED_WAKEUP"
	SUSPICION_OF_DDOS_ATTACK ExceptionIdAnyOf = "SUSPICION_OF_DDOS_ATTACK"
	WRONG_DESTINATION_ADDRESS ExceptionIdAnyOf = "WRONG_DESTINATION_ADDRESS"
	TOO_FREQUENT_SERVICE_ACCESS ExceptionIdAnyOf = "TOO_FREQUENT_SERVICE_ACCESS"
	UNEXPECTED_RADIO_LINK_FAILURES ExceptionIdAnyOf = "UNEXPECTED_RADIO_LINK_FAILURES"
	PING_PONG_ACROSS_CELLS ExceptionIdAnyOf = "PING_PONG_ACROSS_CELLS"
)
