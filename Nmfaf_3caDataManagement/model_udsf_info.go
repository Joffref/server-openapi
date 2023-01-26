/*
 * Nmfaf_3caDataManagement
 *
 * MFAF 3GPP Consumer Adaptor (3CA) Data Management Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_Nmfaf_3caDataManagement

// UdsfInfo - Information related to UDSF
type UdsfInfo struct {

	// Identifier of a group of NFs.
	GroupId string `json:"groupId,omitempty"`

	SupiRanges []SupiRange `json:"supiRanges,omitempty"`

	// A map (list of key-value pairs) where realmId serves as key and each value in the map is an array of IdentityRanges. Each IdentityRange is a range of storageIds. 
	StorageIdRanges map[string][]IdentityRange `json:"storageIdRanges,omitempty"`
}