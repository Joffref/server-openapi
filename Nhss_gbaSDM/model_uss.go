/*
 * Nhss_gbaSDM
 *
 * Nhss Subscriber Data Management Service for GBA.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_Nhss_gbaSDM

// Uss - User Security Settings for a given GAA Service
type Uss struct {

	// Integer where the allowed values correspond to the value range of an unsigned 32-bit integer. 
	GsId int32 `json:"gsId"`

	// Integer where the allowed values correspond to the value range of an unsigned 32-bit integer. 
	GsType int32 `json:"gsType"`

	UeIds []UeIdsItem `json:"ueIds"`

	// Character string representing a NAF Group
	NafGroup string `json:"nafGroup,omitempty"`

	Flags []FlagsItem `json:"flags,omitempty"`

	KeyChoice KeyChoice `json:"keyChoice,omitempty"`
}