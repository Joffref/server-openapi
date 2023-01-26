/*
 * Nnwdaf_AnalyticsInfo
 *
 * Nnwdaf_AnalyticsInfo Service API.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.2.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_Nnwdaf_AnalyticsInfo

// IpEndPoint - IP addressing information of a given NFService; it consists on, e.g. IP address, TCP port, transport protocol... 
type IpEndPoint struct {

	// String identifying a IPv4 address formatted in the 'dotted decimal' notation as defined in RFC 1166. 
	Ipv4Address string `json:"ipv4Address,omitempty"`

	Ipv6Address Ipv6Addr `json:"ipv6Address,omitempty"`

	Transport TransportProtocol1 `json:"transport,omitempty"`

	Port int32 `json:"port,omitempty"`
}
