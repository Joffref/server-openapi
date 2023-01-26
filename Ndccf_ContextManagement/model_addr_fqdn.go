/*
 * Ndccf_ContextManagement
 *
 * DCCF Context Management Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_Ndccf_ContextManagement

// AddrFqdn - IP address and/or FQDN.
type AddrFqdn struct {

	IpAddr IpAddr `json:"ipAddr,omitempty"`

	// Indicates an FQDN.
	Fqdn string `json:"fqdn,omitempty"`
}
