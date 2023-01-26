/*
 * 3gpp-data-reporting
 *
 * API for 3GPP Data Reporting.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_DataReporting

type CommunicationRecordAllOf struct {

	TimeInterval TimeWindow `json:"timeInterval"`

	// Unsigned integer identifying a volume in units of bytes.
	UplinkVolume int64 `json:"uplinkVolume,omitempty"`

	// Unsigned integer identifying a volume in units of bytes.
	DownlinkVolume int64 `json:"downlinkVolume,omitempty"`
}
