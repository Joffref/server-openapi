/*
 * Npcf_SMPolicyControl API
 *
 * Session Management Policy Control Service   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.2.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_Npcf_SMPolicyControl

// SteeringMode - Contains the steering mode value and parameters determined by the PCF.
type SteeringMode struct {

	SteerModeValue SteerModeValue `json:"steerModeValue"`

	Active AccessType `json:"active,omitempty"`

	Standby AccessTypeRm `json:"standby,omitempty"`

	// Unsigned Integer, i.e. only value 0 and integers above 0 are permissible.
	Var3gLoad int32 `json:"3gLoad,omitempty"`

	PrioAcc AccessType `json:"prioAcc,omitempty"`

	ThresValue *ThresholdValue `json:"thresValue,omitempty"`

	SteerModeInd SteerModeIndicator `json:"steerModeInd,omitempty"`
}
