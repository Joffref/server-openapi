/*
 * Unified Data Repository Service API file for Application Data
 *
 * The API version is defined in 3GPP TS 29.504   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: -
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_Application_Data

// UrspRuleRequest - Contains parameters that can be used to guide the URSP.
type UrspRuleRequest struct {

	TrafficDesc TrafficDescriptorComponents `json:"trafficDesc,omitempty"`

	// Unsigned Integer, i.e. only value 0 and integers above 0 are permissible.
	RelatPrecedence int32 `json:"relatPrecedence,omitempty"`

	// Sets of parameters that may be used to guide the Route Selection Descriptors of the URSP.
	RouteSelParamSets []RouteSelectionParameterSet `json:"routeSelParamSets,omitempty"`
}
