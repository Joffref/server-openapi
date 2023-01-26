/*
 * 3gpp-bdt
 *
 * API for BDT resouce management.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.2.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_ResourceManagementOfBdt

// TransferPolicy - Represents an offered transfer policy sent from the SCEF to the SCS/AS, or a selected transfer policy sent from the SCS/AS to the SCEF.
type TransferPolicy struct {

	// Identifier for the transfer policy
	BdtPolicyId int32 `json:"bdtPolicyId"`

	// integer indicating a bandwidth in bits per second.
	MaxUplinkBandwidth int32 `json:"maxUplinkBandwidth,omitempty"`

	// integer indicating a bandwidth in bits per second.
	MaxDownlinkBandwidth int32 `json:"maxDownlinkBandwidth,omitempty"`

	// Indicates the rating group during the time window.
	RatingGroup int32 `json:"ratingGroup"`

	TimeWindow TimeWindow `json:"timeWindow"`
}
