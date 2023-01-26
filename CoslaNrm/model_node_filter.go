/*
 * coslaNrm
 *
 * OAS 3.0.1 specification of the Cosla NRM © 2020, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.
 *
 * API version: 17.3.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_CoslaNrm

type NodeFilter struct {

	AreaOfInterest AreaOfInterest `json:"areaOfInterest,omitempty"`

	NetworkDomain string `json:"networkDomain,omitempty"`

	CpUpType string `json:"cpUpType,omitempty"`

	Sst int32 `json:"sst,omitempty"`
}