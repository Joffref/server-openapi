/*
 * TS 28.532 Streaming data reporting service
 *
 * OAS 3.0.1 specification for the Streaming data reporting service (Streaming MnS) © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.
 *
 * API version: 17.1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_StreamingDataMnS

// EventThresholdType - See details in 3GPP TS 32.422 clause 5.10.7, 5.10.7a, 5.10.13 and 5.10.14.
type EventThresholdType struct {

	EventThresholdRSRP EventThresholdTypeEventThresholdRsrp `json:"EventThresholdRSRP,omitempty"`

	EventThresholdRSRQ EventThresholdTypeEventThresholdRsrq `json:"EventThresholdRSRQ,omitempty"`

	EventThreshold1F EventThresholdTypeEventThreshold1F `json:"EventThreshold1F,omitempty"`

	EventThreshold1I int32 `json:"EventThreshold1I,omitempty"`
}
