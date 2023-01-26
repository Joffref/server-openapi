/*
 * Neasdf_BaselineDNSPattern
 *
 * EASDF Baseline DNS Pattern Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_Neasdf_BaselineDNSPattern

// Ipv4AddressRange - Range of IPv4 addresses
type Ipv4AddressRange struct {

	// String identifying a IPv4 address formatted in the 'dotted decimal' notation as defined in RFC 1166. 
	Start string `json:"start"`

	// String identifying a IPv4 address formatted in the 'dotted decimal' notation as defined in RFC 1166. 
	End string `json:"end"`
}