/*
 * Nudr_DataRepository API OpenAPI file
 *
 * Unified Data Repository Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 2.2.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_Nudr_DR

// NetworkAreaInfo2 - Describes a network area information in which the NF service consumer requests the number of UEs. 
type NetworkAreaInfo2 struct {

	// Contains a list of E-UTRA cell identities.
	Ecgis []Ecgi1 `json:"ecgis,omitempty"`

	// Contains a list of NR cell identities.
	Ncgis []Ncgi1 `json:"ncgis,omitempty"`

	// Contains a list of NG RAN nodes.
	GRanNodeIds []GlobalRanNodeId1 `json:"gRanNodeIds,omitempty"`

	// Contains a list of tracking area identities.
	Tais []Tai1 `json:"tais,omitempty"`
}
