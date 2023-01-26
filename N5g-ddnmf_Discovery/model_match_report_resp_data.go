/*
 * N5g-ddnmf_Discovery API
 *
 * N5g-ddnmf_Discovery Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_N5g-ddnmf_Discovery

import (
	"time"
)

// MatchReportRespData - Represents Match Report Acknowledgement
type MatchReportRespData struct {

	ProseAppIdNames []string `json:"proseAppIdNames,omitempty"`

	// string with format 'date-time' as defined in OpenAPI.
	ValidityTime time.Time `json:"validityTime,omitempty"`

	// Contains the metadata.
	MetaData string `json:"metaData,omitempty"`

	MetaDataIndexMasks []string `json:"metaDataIndexMasks,omitempty"`
}