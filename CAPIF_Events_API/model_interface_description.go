/*
 * CAPIF_Events_API
 *
 * API for event subscription management.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.2.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_CAPIF_Events_API

// InterfaceDescription - Represents the description of an API's interface.
type InterfaceDescription struct {

	// string identifying a Ipv4 address formatted in the \"dotted decimal\" notation as defined in IETF RFC 1166.
	Ipv4Addr string `json:"ipv4Addr,omitempty"`

	// string identifying a Ipv6 address formatted according to clause 4 in IETF RFC 5952. The mixed Ipv4 Ipv6 notation according to clause 5 of IETF RFC 5952 shall not be used.
	Ipv6Addr string `json:"ipv6Addr,omitempty"`

	// Unsigned integer with valid values between 0 and 65535.
	Port int32 `json:"port,omitempty"`

	// Security methods supported by the interface, it take precedence over the security methods provided in AefProfile, for this specific interface. 
	SecurityMethods []SecurityMethod `json:"securityMethods,omitempty"`
}
