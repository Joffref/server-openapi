/*
 * Nnwdaf_AnalyticsInfo
 *
 * Nnwdaf_AnalyticsInfo Service API.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.2.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_Nnwdaf_AnalyticsInfo

type PdnConnectivityStatReport struct {

	PdnConnStat PdnConnectivityStatus `json:"pdnConnStat"`

	// String representing a Data Network as defined in clause 9A of 3GPP TS 23.003;  it shall contain either a DNN Network Identifier, or a full DNN with both the Network  Identifier and Operator Identifier, as specified in 3GPP TS 23.003 clause 9.1.1 and 9.1.2. It shall be coded as string in which the labels are separated by dots  (e.g. \"Label1.Label2.Label3\"). 
	Dnn string `json:"dnn,omitempty"`

	// Unsigned integer identifying a PDU session, within the range 0 to 255, as specified in  clause 11.2.3.1b, bits 1 to 8, of 3GPP TS 24.007. If the PDU Session ID is allocated by the  Core Network for UEs not supporting N1 mode, reserved range 64 to 95 is used. PDU Session ID  within the reserved range is only visible in the Core Network.  
	PduSeId int32 `json:"pduSeId,omitempty"`

	// String identifying a IPv4 address formatted in the 'dotted decimal' notation as defined in RFC 1166. 
	Ipv4Addr string `json:"ipv4Addr,omitempty"`

	Ipv6Prefixes []Ipv6Prefix `json:"ipv6Prefixes,omitempty"`

	Ipv6Addrs []Ipv6Addr `json:"ipv6Addrs,omitempty"`

	PduSessType PduSessionType `json:"pduSessType,omitempty"`
}
