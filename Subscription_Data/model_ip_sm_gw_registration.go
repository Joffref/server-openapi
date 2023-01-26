/*
 * Unified Data Repository Service API file for subscription data
 *
 * Unified Data Repository Service (subscription data).   The API version is defined in 3GPP TS 29.504.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: -
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_Subscription_Data

type IpSmGwRegistration struct {

	IpSmGwMapAddress string `json:"ipSmGwMapAddress,omitempty"`

	IpSmGwDiameterAddress NetworkNodeDiameterAddress `json:"ipSmGwDiameterAddress,omitempty"`

	// String identifying a IPv4 address formatted in the 'dotted decimal' notation as defined in RFC 1166. 
	IpsmgwIpv4 string `json:"ipsmgwIpv4,omitempty"`

	IpsmgwIpv6 Ipv6Addr `json:"ipsmgwIpv6,omitempty"`

	// Fully Qualified Domain Name
	IpsmgwFqdn string `json:"ipsmgwFqdn,omitempty"`

	// String uniquely identifying a NF instance. The format of the NF Instance ID shall be a  Universally Unique Identifier (UUID) version 4, as described in IETF RFC 4122.  
	NfInstanceId string `json:"nfInstanceId,omitempty"`

	UnriIndicator bool `json:"unriIndicator,omitempty"`

	ResetIds []string `json:"resetIds,omitempty"`

	IpSmGwSbiSupInd bool `json:"ipSmGwSbiSupInd,omitempty"`
}
