/*
 * Nudsf_Timer
 *
 * Nudsf Timer Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_Nudsf_Timer

// SearchCondition - A logical condition
type SearchCondition struct {

	Cond ConditionOperator `json:"cond"`

	Units []SearchExpression `json:"units"`

	// Represents the Identifier of a Meta schema.
	SchemaId string `json:"schemaId,omitempty"`
}
