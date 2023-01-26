/*
 * Nudr_DataRepository API OpenAPI file
 *
 * Unified Data Repository Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 2.2.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_Nudr_DR

// AcsInfoRm - This data type is defined in the same way as the 'AcsInfo' data type, but with the  OpenAPI 'nullable: true' property. 
type AcsInfoRm struct {

	// String providing an URI formatted according to RFC 3986.
	AcsUrl string `json:"acsUrl,omitempty"`

	// String identifying a IPv4 address formatted in the 'dotted decimal' notation as defined in RFC 1166. 
	AcsIpv4Addr string `json:"acsIpv4Addr,omitempty"`

	AcsIpv6Addr Ipv6Addr `json:"acsIpv6Addr,omitempty"`
}
