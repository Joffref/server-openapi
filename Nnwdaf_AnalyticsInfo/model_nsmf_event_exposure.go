/*
 * Nnwdaf_AnalyticsInfo
 *
 * Nnwdaf_AnalyticsInfo Service API.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.2.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_Nnwdaf_AnalyticsInfo

import (
	"time"
)

// NsmfEventExposure - Represents an Individual SMF Notification Subscription resource. The serviveName property corresponds to the serviceName in the main body of the specification. 
type NsmfEventExposure struct {

	// String identifying a Supi that shall contain either an IMSI, a network specific identifier, a Global Cable Identifier (GCI) or a Global Line Identifier (GLI) as specified in clause  2.2A of 3GPP TS 23.003. It shall be formatted as follows  - for an IMSI \"imsi-<imsi>\", where <imsi> shall be formatted according to clause 2.2    of 3GPP TS 23.003 that describes an IMSI.  - for a network specific identifier \"nai-<nai>, where <nai> shall be formatted    according to clause 28.7.2 of 3GPP TS 23.003 that describes an NAI.  - for a GCI \"gci-<gci>\", where <gci> shall be formatted according to clause 28.15.2    of 3GPP TS 23.003.  - for a GLI \"gli-<gli>\", where <gli> shall be formatted according to clause 28.16.2 of    3GPP TS 23.003.To enable that the value is used as part of an URI, the string shall    only contain characters allowed according to the \"lower-with-hyphen\" naming convention    defined in 3GPP TS 29.501. 
	Supi string `json:"supi,omitempty"`

	// String identifying a Gpsi shall contain either an External Id or an MSISDN.  It shall be formatted as follows -External Identifier= \"extid-'extid', where 'extid'  shall be formatted according to clause 19.7.2 of 3GPP TS 23.003 that describes an  External Identifier.  
	Gpsi string `json:"gpsi,omitempty"`

	// Any UE indication. This IE shall be present if the event subscription is applicable to any UE. Default value \"false\" is used, if not present. 
	AnyUeInd bool `json:"anyUeInd,omitempty"`

	// String identifying a group of devices network internal globally unique ID which identifies a set of IMSIs, as specified in clause 19.9 of 3GPP TS 23.003.  
	GroupId string `json:"groupId,omitempty"`

	// Unsigned integer identifying a PDU session, within the range 0 to 255, as specified in  clause 11.2.3.1b, bits 1 to 8, of 3GPP TS 24.007. If the PDU Session ID is allocated by the  Core Network for UEs not supporting N1 mode, reserved range 64 to 95 is used. PDU Session ID  within the reserved range is only visible in the Core Network.  
	PduSeId int32 `json:"pduSeId,omitempty"`

	// String representing a Data Network as defined in clause 9A of 3GPP TS 23.003;  it shall contain either a DNN Network Identifier, or a full DNN with both the Network  Identifier and Operator Identifier, as specified in 3GPP TS 23.003 clause 9.1.1 and 9.1.2. It shall be coded as string in which the labels are separated by dots  (e.g. \"Label1.Label2.Label3\"). 
	Dnn string `json:"dnn,omitempty"`

	Snssai Snssai `json:"snssai,omitempty"`

	// Identifies an Individual SMF Notification Subscription. To enable that the value is used as part of a URI, the string shall only contain characters allowed according to the \"lower-with-hyphen\" naming convention defined in 3GPP TS 29.501. In an OpenAPI schema, the format shall be designated as \"SubId\". 
	SubId string `json:"subId,omitempty"`

	// Notification Correlation ID assigned by the NF service consumer.
	NotifId string `json:"notifId"`

	// String providing an URI formatted according to RFC 3986.
	NotifUri string `json:"notifUri"`

	// Alternate or backup IPv4 address(es) where to send Notifications.
	AltNotifIpv4Addrs []string `json:"altNotifIpv4Addrs,omitempty"`

	// Alternate or backup IPv6 address(es) where to send Notifications.
	AltNotifIpv6Addrs []Ipv6Addr `json:"altNotifIpv6Addrs,omitempty"`

	// Alternate or backup FQDN(s) where to send Notifications.
	AltNotifFqdns []string `json:"altNotifFqdns,omitempty"`

	// Subscribed events
	EventSubs []EventSubscription1 `json:"eventSubs"`

	EventNotifs []EventNotification1 `json:"eventNotifs,omitempty"`

	ImmeRep bool `json:"ImmeRep,omitempty"`

	NotifMethod NotificationMethod1 `json:"notifMethod,omitempty"`

	// Unsigned Integer, i.e. only value 0 and integers above 0 are permissible.
	MaxReportNbr int32 `json:"maxReportNbr,omitempty"`

	// string with format 'date-time' as defined in OpenAPI.
	Expiry time.Time `json:"expiry,omitempty"`

	// indicating a time in seconds.
	RepPeriod int32 `json:"repPeriod,omitempty"`

	Guami Guami `json:"guami,omitempty"`

	ServiveName ServiceName `json:"serviveName,omitempty"`

	// A string used to indicate the features supported by an API that is used as defined in clause  6.6 in 3GPP TS 29.500. The string shall contain a bitmask indicating supported features in  hexadecimal representation Each character in the string shall take a value of \"0\" to \"9\",  \"a\" to \"f\" or \"A\" to \"F\" and shall represent the support of 4 features as described in  table 5.2.2-3. The most significant character representing the highest-numbered features shall  appear first in the string, and the character representing features 1 to 4 shall appear last  in the string. The list of features and their numbering (starting with 1) are defined  separately for each API. If the string contains a lower number of characters than there are  defined features for an API, all features that would be represented by characters that are not  present in the string are not supported. 
	SupportedFeatures string `json:"supportedFeatures,omitempty"`

	// Unsigned integer indicating Sampling Ratio (see clauses 4.15.1 of 3GPP TS 23.502), expressed in percent.  
	SampRatio int32 `json:"sampRatio,omitempty"`

	// Criteria for partitioning the UEs before applying the sampling ratio.
	PartitionCriteria []PartitioningCriteria `json:"partitionCriteria,omitempty"`

	// indicating a time in seconds.
	GrpRepTime int32 `json:"grpRepTime,omitempty"`

	NotifFlag NotificationFlag `json:"notifFlag,omitempty"`
}
