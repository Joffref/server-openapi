/*
 * Nmfaf_3caDataManagement
 *
 * MFAF 3GPP Consumer Adaptor (3CA) Data Management Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package Nmfaf_3caDataManagement

// TwifInfo - Addressing information (IP addresses, FQDN) of the TWIF
type TwifInfo struct {

	Ipv4EndpointAddresses []string `json:"ipv4EndpointAddresses,omitempty"`

	Ipv6EndpointAddresses []Ipv6Addr `json:"ipv6EndpointAddresses,omitempty"`

	// Fully Qualified Domain Name
	EndpointFqdn string `json:"endpointFqdn,omitempty"`
}