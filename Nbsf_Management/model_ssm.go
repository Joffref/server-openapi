/*
 * Nbsf_Management
 *
 * Binding Support Management Service API.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.3.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_Nbsf_Management

// Ssm - Source specific IP multicast address
type Ssm struct {

	SourceIpAddr IpAddr `json:"sourceIpAddr"`

	DestIpAddr IpAddr `json:"destIpAddr"`
}
