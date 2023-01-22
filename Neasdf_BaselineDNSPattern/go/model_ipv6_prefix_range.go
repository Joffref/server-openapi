/*
 * Neasdf_BaselineDNSPattern
 *
 * EASDF Baseline DNS Pattern Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package Neasdf_BaselineDNSPattern

// Ipv6PrefixRange - Range of IPv6 prefixes
type Ipv6PrefixRange struct {

	Start Ipv6Prefix `json:"start"`

	End Ipv6Prefix `json:"end"`
}