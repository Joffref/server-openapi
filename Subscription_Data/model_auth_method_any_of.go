/*
 * Unified Data Repository Service API file for subscription data
 *
 * Unified Data Repository Service (subscription data).   The API version is defined in 3GPP TS 29.504.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: -
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_Subscription_Data

type AuthMethodAnyOf string

// List of AuthMethodAnyOf
const (
	_5_G_AKA AuthMethodAnyOf = "5G_AKA"
	EAP_AKA_PRIME AuthMethodAnyOf = "EAP_AKA_PRIME"
	EAP_TLS AuthMethodAnyOf = "EAP_TLS"
	EAP_TTLS AuthMethodAnyOf = "EAP_TTLS"
	NONE AuthMethodAnyOf = "NONE"
)
