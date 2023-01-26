/*
 * Npcf_PolicyAuthorization Service API
 *
 * PCF Policy Authorization Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.2.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_Npcf_PolicyAuthorization

// ResourcesAllocationInfo - Describes the status of the PCC rule(s) related to certain media components.
type ResourcesAllocationInfo struct {

	McResourcStatus MediaComponentResourcesStatus `json:"mcResourcStatus,omitempty"`

	Flows []Flows `json:"flows,omitempty"`

	AltSerReq string `json:"altSerReq,omitempty"`
}