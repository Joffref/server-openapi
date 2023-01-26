/*
 * Nmfaf_3caDataManagement
 *
 * MFAF 3GPP Consumer Adaptor (3CA) Data Management Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_Nmfaf_3caDataManagement

// ScpInfo - Information of an SCP Instance
type ScpInfo struct {

	// A map (list of key-value pairs) where the key of the map shall be the string identifying an SCP domain 
	ScpDomainInfoList map[string]ScpDomainInfo `json:"scpDomainInfoList,omitempty"`

	ScpPrefix string `json:"scpPrefix,omitempty"`

	// Port numbers for HTTP and HTTPS. The key of the map shall be \"http\" or \"https\". 
	ScpPorts map[string]int32 `json:"scpPorts,omitempty"`

	AddressDomains []string `json:"addressDomains,omitempty"`

	Ipv4Addresses []string `json:"ipv4Addresses,omitempty"`

	Ipv6Prefixes []Ipv6Prefix `json:"ipv6Prefixes,omitempty"`

	Ipv4AddrRanges []Ipv4AddressRange `json:"ipv4AddrRanges,omitempty"`

	Ipv6PrefixRanges []Ipv6PrefixRange `json:"ipv6PrefixRanges,omitempty"`

	ServedNfSetIdList []string `json:"servedNfSetIdList,omitempty"`

	RemotePlmnList []PlmnId `json:"remotePlmnList,omitempty"`

	RemoteSnpnList []PlmnIdNid `json:"remoteSnpnList,omitempty"`

	IpReachability IpReachability `json:"ipReachability,omitempty"`

	ScpCapabilities []ScpCapability `json:"scpCapabilities,omitempty"`
}
