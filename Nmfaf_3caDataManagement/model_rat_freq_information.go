/*
 * Nmfaf_3caDataManagement
 *
 * MFAF 3GPP Consumer Adaptor (3CA) Data Management Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_Nmfaf_3caDataManagement

// RatFreqInformation - Represents the RAT type and/or Frequency information.
type RatFreqInformation struct {

	// Set to \"true\" to indicate to handle all the frequencies the NWDAF received, otherwise  set to \"false\" or omit. The \"allFreq\" attribute and the \"freq\" attribute are mutually  exclusive. 
	AllFreq bool `json:"allFreq,omitempty"`

	// Set to \"true\" to indicate to handle all the RAT Types the NWDAF received, otherwise  set to \"false\" or omit. The \"allRat\" attribute and the \"ratType\" attribute are mutually  exclusive. 
	AllRat bool `json:"allRat,omitempty"`

	// Integer value indicating the ARFCN applicable for a downlink, uplink or bi-directional (TDD) NR global frequency raster, as definition of \"ARFCN-ValueNR\" IE in clause 6.3.2 of 3GPP TS 38.331. 
	Freq int32 `json:"freq,omitempty"`

	RatType RatType `json:"ratType,omitempty"`

	SvcExpThreshold ThresholdLevel `json:"svcExpThreshold,omitempty"`

	MatchingDir MatchingDirection `json:"matchingDir,omitempty"`
}
