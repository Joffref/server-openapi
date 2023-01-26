/*
 * 3gpp-analyticsexposure
 *
 * API for Analytics Exposure.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.1.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_AnalyticsExposure

type AnalyticsEventAnyOf string

// List of AnalyticsEventAnyOf
const (
	UE_MOBILITY AnalyticsEventAnyOf = "UE_MOBILITY"
	UE_COMM AnalyticsEventAnyOf = "UE_COMM"
	ABNORMAL_BEHAVIOR AnalyticsEventAnyOf = "ABNORMAL_BEHAVIOR"
	CONGESTION AnalyticsEventAnyOf = "CONGESTION"
	NETWORK_PERFORMANCE AnalyticsEventAnyOf = "NETWORK_PERFORMANCE"
	QOS_SUSTAINABILITY AnalyticsEventAnyOf = "QOS_SUSTAINABILITY"
	DISPERSION AnalyticsEventAnyOf = "DISPERSION"
	DN_PERFORMANCE AnalyticsEventAnyOf = "DN_PERFORMANCE"
	SERVICE_EXPERIENCE AnalyticsEventAnyOf = "SERVICE_EXPERIENCE"
)
