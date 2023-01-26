/*
 * Nnwdaf_DataManagement
 *
 * Nnwdaf_DataManagement API Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_Nnwdaf_DataManagement

// TargetUeIdentification - Identifies the UE to which the request applies.
type TargetUeIdentification struct {

	Supis []string `json:"supis,omitempty"`

	InterGroupIds []string `json:"interGroupIds,omitempty"`

	AnyUeId bool `json:"anyUeId,omitempty"`
}
