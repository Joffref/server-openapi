/*
 * Namf_Communication
 *
 * AMF Communication Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.2.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_Namf_Communication

import (
	"time"
)

// QosSustainabilityInfo - Represents the QoS Sustainability information.
type QosSustainabilityInfo struct {

	AreaInfo NetworkAreaInfo `json:"areaInfo,omitempty"`

	// string with format 'date-time' as defined in OpenAPI.
	StartTs time.Time `json:"startTs,omitempty"`

	// string with format 'date-time' as defined in OpenAPI.
	EndTs time.Time `json:"endTs,omitempty"`

	QosFlowRetThd RetainabilityThreshold `json:"qosFlowRetThd,omitempty"`

	// String representing a bit rate; the prefixes follow the standard symbols from The International System of Units, and represent x1000 multipliers, with the exception that prefix \"K\" is used to represent the standard symbol \"k\". 
	RanUeThrouThd string `json:"ranUeThrouThd,omitempty"`

	Snssai Snssai `json:"snssai,omitempty"`

	// Unsigned Integer, i.e. only value 0 and integers above 0 are permissible.
	Confidence int32 `json:"confidence,omitempty"`
}
