/*
 * Nudm_SDM
 *
 * Nudm Subscriber Data Management Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 2.2.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_Nudm_SDM

type FrameRouteInfo struct {

	// \"String identifying a IPv4 address mask formatted in the 'dotted decimal' notation as defined in RFC 1166.\" 
	Ipv4Mask string `json:"ipv4Mask,omitempty"`

	Ipv6Prefix Ipv6Prefix `json:"ipv6Prefix,omitempty"`
}
