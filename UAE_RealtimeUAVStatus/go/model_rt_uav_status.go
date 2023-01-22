/*
 * UAE Server Real-time UAV Status Service
 *
 * UAE Server Real-time UAV Status Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package UAE_RealtimeUAVStatus

// RtUavStatus - Represents real-time UAV status information.
type RtUavStatus struct {

	UavId UavId `json:"uavId,omitempty"`

	UavNetConnStatus UavNetConnStatus `json:"uavNetConnStatus,omitempty"`

	UavLocInfo LocationInfo `json:"uavLocInfo,omitempty"`
}