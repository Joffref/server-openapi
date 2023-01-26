/*
 * Eees Application Context Relocation Service
 *
 * Eees Application Context Relocation Service.   © 2021, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved. 
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_Eees_AppContextRelocation

// EndPoint - The end point information to reach EAS.
type EndPoint struct {

	// Fully Qualified Domain Name
	Fqdn string `json:"fqdn,omitempty"`

	// IPv4 addresses of the edge server.
	Ipv4Addrs []string `json:"ipv4Addrs,omitempty"`

	// IPv6 addresses of the edge server.
	Ipv6Addrs []string `json:"ipv6Addrs,omitempty"`

	// string providing an URI formatted according to IETF RFC 3986.
	Uri string `json:"uri,omitempty"`
}
