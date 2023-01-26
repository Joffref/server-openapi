/*
 * Nadrf_DataManagement
 *
 * ADRF Data Management Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_Nadrf_DataManagement

import (
	"time"
)

// EventNotification1 - Represents a notification related to a single event that occurred.
type EventNotification1 struct {

	Event SmfEvent `json:"event"`

	// string with format 'date-time' as defined in OpenAPI.
	TimeStamp time.Time `json:"timeStamp"`

	// String identifying a Supi that shall contain either an IMSI, a network specific identifier, a Global Cable Identifier (GCI) or a Global Line Identifier (GLI) as specified in clause  2.2A of 3GPP TS 23.003. It shall be formatted as follows  - for an IMSI \"imsi-<imsi>\", where <imsi> shall be formatted according to clause 2.2    of 3GPP TS 23.003 that describes an IMSI.  - for a network specific identifier \"nai-<nai>, where <nai> shall be formatted    according to clause 28.7.2 of 3GPP TS 23.003 that describes an NAI.  - for a GCI \"gci-<gci>\", where <gci> shall be formatted according to clause 28.15.2    of 3GPP TS 23.003.  - for a GLI \"gli-<gli>\", where <gli> shall be formatted according to clause 28.16.2 of    3GPP TS 23.003.To enable that the value is used as part of an URI, the string shall    only contain characters allowed according to the \"lower-with-hyphen\" naming convention    defined in 3GPP TS 29.501. 
	Supi string `json:"supi,omitempty"`

	// String identifying a Gpsi shall contain either an External Id or an MSISDN.  It shall be formatted as follows -External Identifier= \"extid-'extid', where 'extid'  shall be formatted according to clause 19.7.2 of 3GPP TS 23.003 that describes an  External Identifier.  
	Gpsi string `json:"gpsi,omitempty"`

	UeIpAddr IpAddr `json:"ueIpAddr,omitempty"`

	// Transaction Information.
	TransacInfos []TransactionInfo `json:"transacInfos,omitempty"`

	// DNAI (Data network access identifier), see clause 5.6.7 of 3GPP TS 23.501.
	SourceDnai string `json:"sourceDnai,omitempty"`

	// DNAI (Data network access identifier), see clause 5.6.7 of 3GPP TS 23.501.
	TargetDnai string `json:"targetDnai,omitempty"`

	DnaiChgType DnaiChangeType `json:"dnaiChgType,omitempty"`

	// String identifying a IPv4 address formatted in the 'dotted decimal' notation as defined in RFC 1166. 
	SourceUeIpv4Addr string `json:"sourceUeIpv4Addr,omitempty"`

	SourceUeIpv6Prefix Ipv6Prefix `json:"sourceUeIpv6Prefix,omitempty"`

	// String identifying a IPv4 address formatted in the 'dotted decimal' notation as defined in RFC 1166. 
	TargetUeIpv4Addr string `json:"targetUeIpv4Addr,omitempty"`

	TargetUeIpv6Prefix Ipv6Prefix `json:"targetUeIpv6Prefix,omitempty"`

	SourceTraRouting *RouteToLocation `json:"sourceTraRouting,omitempty"`

	TargetTraRouting *RouteToLocation `json:"targetTraRouting,omitempty"`

	// String identifying a MAC address formatted in the hexadecimal notation according to clause 1.1 and clause 2.1 of RFC 7042. 
	UeMac string `json:"ueMac,omitempty"`

	// String identifying a IPv4 address formatted in the 'dotted decimal' notation as defined in RFC 1166. 
	AdIpv4Addr string `json:"adIpv4Addr,omitempty"`

	AdIpv6Prefix Ipv6Prefix `json:"adIpv6Prefix,omitempty"`

	// String identifying a IPv4 address formatted in the 'dotted decimal' notation as defined in RFC 1166. 
	ReIpv4Addr string `json:"reIpv4Addr,omitempty"`

	ReIpv6Prefix Ipv6Prefix `json:"reIpv6Prefix,omitempty"`

	PlmnId PlmnId `json:"plmnId,omitempty"`

	AccType AccessType `json:"accType,omitempty"`

	// Unsigned integer identifying a PDU session, within the range 0 to 255, as specified in  clause 11.2.3.1b, bits 1 to 8, of 3GPP TS 24.007. If the PDU Session ID is allocated by the  Core Network for UEs not supporting N1 mode, reserved range 64 to 95 is used. PDU Session ID  within the reserved range is only visible in the Core Network.  
	PduSeId int32 `json:"pduSeId,omitempty"`

	RatType RatType `json:"ratType,omitempty"`

	DddStatus DlDataDeliveryStatus `json:"dddStatus,omitempty"`

	DddTraDescriptor DddTrafficDescriptor `json:"dddTraDescriptor,omitempty"`

	// string with format 'date-time' as defined in OpenAPI.
	MaxWaitTime time.Time `json:"maxWaitTime,omitempty"`

	CommFailure CommunicationFailure `json:"commFailure,omitempty"`

	// String identifying a IPv4 address formatted in the 'dotted decimal' notation as defined in RFC 1166. 
	Ipv4Addr string `json:"ipv4Addr,omitempty"`

	Ipv6Prefixes []Ipv6Prefix `json:"ipv6Prefixes,omitempty"`

	Ipv6Addrs []Ipv6Addr `json:"ipv6Addrs,omitempty"`

	PduSessType PduSessionType `json:"pduSessType,omitempty"`

	// Unsigned integer identifying a QoS flow, within the range 0 to 63.
	Qfi int32 `json:"qfi,omitempty"`

	// String providing an application identifier.
	AppId string `json:"appId,omitempty"`

	// Descriptor(s) for non-IP traffic. It allows the encoding of multiple UL and/or DL flows. Each entry of the array describes a single Ethernet flow. 
	EthFlowDescs []EthFlowDescription `json:"ethFlowDescs,omitempty"`

	// Contains the UL and/or DL Ethernet flows. Each entry of the array describes a single Ethernet flow. 
	EthfDescs []EthFlowDescription `json:"ethfDescs,omitempty"`

	// Descriptor(s) for IP traffic. It allows the encoding of multiple UL and/or DL flows. Each entry of the array describes a single IP flow. 
	FlowDescs []string `json:"flowDescs,omitempty"`

	// Contains the UL and/or DL IP flows. Each entry of the array describes a single IP flow. 
	FDescs []string `json:"fDescs,omitempty"`

	// String representing a Data Network as defined in clause 9A of 3GPP TS 23.003;  it shall contain either a DNN Network Identifier, or a full DNN with both the Network  Identifier and Operator Identifier, as specified in 3GPP TS 23.003 clause 9.1.1 and 9.1.2. It shall be coded as string in which the labels are separated by dots  (e.g. \"Label1.Label2.Label3\"). 
	Dnn string `json:"dnn,omitempty"`

	Snssai Snssai `json:"snssai,omitempty"`

	UlDelays []int32 `json:"ulDelays,omitempty"`

	DlDelays []int32 `json:"dlDelays,omitempty"`

	RtDelays []int32 `json:"rtDelays,omitempty"`

	TimeWindow TimeWindow `json:"timeWindow,omitempty"`

	SmNasFromUe SmNasFromUe `json:"smNasFromUe,omitempty"`

	SmNasFromSmf SmNasFromSmf `json:"smNasFromSmf,omitempty"`

	// Indicates whether the redundant transmission is setup or terminated. Set to \"true\" if  the redundant transmission is setup, otherwise set to \"false\" if the redundant  transmission is terminated. Default value is set to \"false\". 
	UpRedTrans bool `json:"upRedTrans,omitempty"`

	SsId string `json:"ssId,omitempty"`

	BssId string `json:"bssId,omitempty"`

	// string with format 'date-time' as defined in OpenAPI.
	StartWlan time.Time `json:"startWlan,omitempty"`

	// string with format 'date-time' as defined in OpenAPI.
	EndWlan time.Time `json:"endWlan,omitempty"`

	PduSessInfos []PduSessionInformation `json:"pduSessInfos,omitempty"`

	UpfInfo UpfInformation `json:"upfInfo,omitempty"`
}
