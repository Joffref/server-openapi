/*
 * Common Data Types
 *
 * Common Data Types for Service Based Interfaces.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.   
 *
 * API version: 1.4.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_CommonData

// SecondaryRatUsageInfo - Secondary RAT Usage Information to report usage data for a secondary RAT for QoS flows and/or the whole PDU session. 
type SecondaryRatUsageInfo struct {

	SecondaryRatType RatType `json:"secondaryRatType"`

	QosFlowsUsageData []QosFlowUsageReport `json:"qosFlowsUsageData,omitempty"`

	PduSessionUsageData []VolumeTimedReport `json:"pduSessionUsageData,omitempty"`
}
