/*
 * Namf_Communication
 *
 * AMF Communication Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.2.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_Namf_Communication

// AddrFqdn - IP address and/or FQDN.
type AddrFqdn struct {

	IpAddr IpAddr `json:"ipAddr,omitempty"`

	// Indicates an FQDN.
	Fqdn string `json:"fqdn,omitempty"`
}
