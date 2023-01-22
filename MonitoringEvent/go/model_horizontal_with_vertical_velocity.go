/*
 * 3gpp-monitoring-event
 *
 * API for Monitoring Event.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.2.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package MonitoringEvent

// HorizontalWithVerticalVelocity - Horizontal and vertical velocity.
type HorizontalWithVerticalVelocity struct {

	// Indicates value of horizontal speed.
	HSpeed float32 `json:"hSpeed"`

	// Indicates value of angle.
	Bearing int32 `json:"bearing"`

	// Indicates value of vertical speed.
	VSpeed float32 `json:"vSpeed"`

	VDirection VerticalDirection `json:"vDirection"`
}