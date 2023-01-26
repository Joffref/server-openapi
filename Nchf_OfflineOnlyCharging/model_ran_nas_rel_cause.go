/*
 * Nchf_OfflineOnlyCharging
 *
 * OfflineOnlyCharging Service © 2021, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved. 
 *
 * API version: 1.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_Nchf_OfflineOnlyCharging

// RanNasRelCause - Contains the RAN/NAS release cause.
type RanNasRelCause struct {

	NgApCause NgApCause `json:"ngApCause,omitempty"`

	// Unsigned Integer, i.e. only value 0 and integers above 0 are permissible.
	Var5gMmCause int32 `json:"5gMmCause,omitempty"`

	// Unsigned Integer, i.e. only value 0 and integers above 0 are permissible.
	Var5gSmCause int32 `json:"5gSmCause,omitempty"`

	// Defines the EPS RAN/NAS release cause.
	EpsCause string `json:"epsCause,omitempty"`
}
