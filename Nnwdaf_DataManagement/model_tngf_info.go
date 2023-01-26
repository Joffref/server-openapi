/*
 * Nnwdaf_DataManagement
 *
 * Nnwdaf_DataManagement API Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_Nnwdaf_DataManagement

// TngfInfo - Infomation of the TNGF endpoints
type TngfInfo struct {

	Ipv4EndpointAddresses []string `json:"ipv4EndpointAddresses,omitempty"`

	Ipv6EndpointAddresses []Ipv6Addr `json:"ipv6EndpointAddresses,omitempty"`

	// Fully Qualified Domain Name
	EndpointFqdn string `json:"endpointFqdn,omitempty"`
}
