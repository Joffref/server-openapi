/*
 * Npcf_SMPolicyControl API
 *
 * Session Management Policy Control Service   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.2.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_Npcf_SMPolicyControl

import (
	"time"
)

// ConditionData - Contains conditions of applicability for a rule.
type ConditionData struct {

	// Uniquely identifies the condition data within a PDU session.
	CondId string `json:"condId"`

	// string with format 'date-time' as defined in OpenAPI with 'nullable:true' property.  
	ActivationTime *time.Time `json:"activationTime,omitempty"`

	// string with format 'date-time' as defined in OpenAPI with 'nullable:true' property.  
	DeactivationTime *time.Time `json:"deactivationTime,omitempty"`

	AccessType AccessType `json:"accessType,omitempty"`

	RatType RatType `json:"ratType,omitempty"`
}