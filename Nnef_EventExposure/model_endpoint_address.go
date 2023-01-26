/*
 * Nnef_EventExposure
 *
 * NEF Event Exposure Service.   © 2022 , 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.2.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_Nnef_EventExposure

type EndpointAddress struct {

	// String identifying a IPv4 address formatted in the 'dotted decimal' notation as defined in RFC 1166. 
	Ipv4Addr string `json:"ipv4Addr,omitempty"`

	Ipv6Addr Ipv6Addr `json:"ipv6Addr,omitempty"`

	// Integer where the allowed values correspond to the value range of an unsigned 16-bit integer.
	PortNumber int32 `json:"portNumber"`
}
