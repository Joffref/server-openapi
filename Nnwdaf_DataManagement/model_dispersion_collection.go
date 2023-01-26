/*
 * Nnwdaf_DataManagement
 *
 * Nnwdaf_DataManagement API Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_Nnwdaf_DataManagement

// DispersionCollection - Dispersion collection per UE location or per slice.
type DispersionCollection struct {

	UeLoc UserLocation `json:"ueLoc,omitempty"`

	Snssai Snssai `json:"snssai,omitempty"`

	Supis []string `json:"supis,omitempty"`

	Gpsis []string `json:"gpsis,omitempty"`

	AppVolumes []ApplicationVolume `json:"appVolumes,omitempty"`

	// Unsigned Integer, i.e. only value 0 and integers above 0 are permissible.
	DisperAmount int32 `json:"disperAmount,omitempty"`

	DisperClass DispersionClass `json:"disperClass,omitempty"`

	// Integer where the allowed values correspond to 1, 2, 3 only.
	UsageRank int32 `json:"usageRank,omitempty"`

	// Unsigned integer indicating Sampling Ratio (see clauses 4.15.1 of 3GPP TS 23.502), expressed in percent.  
	PercentileRank int32 `json:"percentileRank,omitempty"`

	// Unsigned integer indicating Sampling Ratio (see clauses 4.15.1 of 3GPP TS 23.502), expressed in percent.  
	UeRatio int32 `json:"ueRatio,omitempty"`

	// Unsigned Integer, i.e. only value 0 and integers above 0 are permissible.
	Confidence int32 `json:"confidence,omitempty"`
}
