/*
 * Nnwdaf_DataManagement
 *
 * Nnwdaf_DataManagement API Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_Nnwdaf_DataManagement

// N3gaLocation - Contains the Non-3GPP access user location.
type N3gaLocation struct {

	N3gppTai Tai `json:"n3gppTai,omitempty"`

	// This IE shall contain the N3IWF identifier received over NGAP and shall be encoded as a  string of hexadecimal characters. Each character in the string shall take a value of \"0\"  to \"9\", \"a\" to \"f\" or \"A\" to \"F\" and shall represent 4 bits. The most significant  character representing the 4 most significant bits of the N3IWF ID shall appear first in  the string, and the character representing the 4 least significant bit of the N3IWF ID  shall appear last in the string.  
	N3IwfId string `json:"n3IwfId,omitempty"`

	// String identifying a IPv4 address formatted in the 'dotted decimal' notation as defined in RFC 1166. 
	UeIpv4Addr string `json:"ueIpv4Addr,omitempty"`

	UeIpv6Addr Ipv6Addr `json:"ueIpv6Addr,omitempty"`

	// Unsigned Integer, i.e. only value 0 and integers above 0 are permissible.
	PortNumber int32 `json:"portNumber,omitempty"`

	Protocol TransportProtocol `json:"protocol,omitempty"`

	TnapId TnapId `json:"tnapId,omitempty"`

	TwapId TwapId `json:"twapId,omitempty"`

	HfcNodeId HfcNodeId `json:"hfcNodeId,omitempty"`

	// string with format 'bytes' as defined in OpenAPI
	Gli string `json:"gli,omitempty"`

	W5gbanLineType LineType `json:"w5gbanLineType,omitempty"`

	// Global Cable Identifier uniquely identifying the connection between the 5G-CRG or FN-CRG to the 5GS. See clause 28.15.4 of 3GPP TS 23.003. This shall be encoded as a string per clause 28.15.4 of 3GPP TS 23.003, and compliant with the syntax specified  in clause 2.2  of IETF RFC 7542 for the username part of a NAI. The GCI value is specified in CableLabs WR-TR-5WWC-ARCH. 
	Gci string `json:"gci,omitempty"`
}
