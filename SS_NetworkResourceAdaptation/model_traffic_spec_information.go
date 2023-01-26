/*
 * SS_NetworkResourceAdaptation
 *
 * SS Network Resource Adaptation Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_SS_NetworkResourceAdaptation

// TrafficSpecInformation - The traffic classe supported by the DS-TTs and available end-to-end latency value and Priority Code Point (PCP) value. 
type TrafficSpecInformation struct {

	// Integer where the allowed values correspond to the value range of an unsigned 32-bit integer. 
	PcpValue int32 `json:"pcpValue"`

	// Unsigned integer identifying a period of time in units of seconds.
	MaxFramInt int32 `json:"maxFramInt"`

	// Integer where the allowed values correspond to the value range of an unsigned 32-bit integer. 
	MaxFramSize int32 `json:"maxFramSize"`

	// Integer where the allowed values correspond to the value range of an unsigned 32-bit integer. 
	MaxIntFrames int32 `json:"maxIntFrames"`

	// Integer where the allowed values correspond to the value range of an unsigned 32-bit integer. 
	MaxLatency int32 `json:"maxLatency"`
}