/*
 * CAPIF_Security_API
 *
 * API for CAPIF security management.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.2.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package CAPIF_Security_API

// SecurityInformation - Represents the interface details and the security method.
type SecurityInformation struct {

	InterfaceDetails InterfaceDescription `json:"interfaceDetails,omitempty"`

	// Identifier of the API exposing function
	AefId string `json:"aefId,omitempty"`

	// API identifier
	ApiId string `json:"apiId,omitempty"`

	// Security methods preferred by the API invoker for the API interface.
	PrefSecurityMethods []SecurityMethod `json:"prefSecurityMethods"`

	SelSecurityMethod SecurityMethod `json:"selSecurityMethod,omitempty"`

	// Authentication related information
	AuthenticationInfo string `json:"authenticationInfo,omitempty"`

	// Authorization related information
	AuthorizationInfo string `json:"authorizationInfo,omitempty"`
}