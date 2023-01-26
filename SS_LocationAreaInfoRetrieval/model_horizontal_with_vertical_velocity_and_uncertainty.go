/*
 * SS_LocationAreaInfoRetrieval
 *
 * API for SEAL Location Area Info Retrieval.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_SS_LocationAreaInfoRetrieval

// HorizontalWithVerticalVelocityAndUncertainty - Horizontal and vertical velocity with speed uncertainty.
type HorizontalWithVerticalVelocityAndUncertainty struct {

	// Indicates value of horizontal speed.
	HSpeed float32 `json:"hSpeed"`

	// Indicates value of angle.
	Bearing int32 `json:"bearing"`

	// Indicates value of vertical speed.
	VSpeed float32 `json:"vSpeed"`

	VDirection VerticalDirection `json:"vDirection"`

	// Indicates value of speed uncertainty.
	HUncertainty float32 `json:"hUncertainty"`

	// Indicates value of speed uncertainty.
	VUncertainty float32 `json:"vUncertainty"`
}
