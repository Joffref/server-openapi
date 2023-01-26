/*
 * Namf_Communication
 *
 * AMF Communication Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.2.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_Namf_Communication

// Pc5QoSPara - Contains policy data on the PC5 QoS parameters.
type Pc5QoSPara struct {

	Pc5QosFlowList []Pc5QosFlowItem `json:"pc5QosFlowList"`

	// String representing a bit rate; the prefixes follow the standard symbols from The International System of Units, and represent x1000 multipliers, with the exception that prefix \"K\" is used to represent the standard symbol \"k\". 
	Pc5LinkAmbr string `json:"pc5LinkAmbr,omitempty"`
}
