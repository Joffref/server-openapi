/*
 * AUSF API
 *
 * AUSF UE Authentication Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.2.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_Nausf_UEAuthentication

// ServerAddressingInfo - Contains addressing information (IP addresses and/or FQDNs) of a server.
type ServerAddressingInfo struct {

	Ipv4Addresses []string `json:"ipv4Addresses,omitempty"`

	Ipv6Addresses []Ipv6Addr `json:"ipv6Addresses,omitempty"`

	FqdnList []string `json:"fqdnList,omitempty"`
}
