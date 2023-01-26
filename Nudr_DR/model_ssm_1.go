/*
 * Nudr_DataRepository API OpenAPI file
 *
 * Unified Data Repository Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 2.2.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_Nudr_DR

// Ssm1 - Source specific IP multicast address
type Ssm1 struct {

	SourceIpAddr IpAddr1 `json:"sourceIpAddr"`

	DestIpAddr IpAddr1 `json:"destIpAddr"`
}
