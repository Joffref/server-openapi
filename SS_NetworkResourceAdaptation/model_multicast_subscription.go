/*
 * SS_NetworkResourceAdaptation
 *
 * SS Network Resource Adaptation Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_SS_NetworkResourceAdaptation

import (
	"time"
)

// MulticastSubscription - Represents a multicast subscription.
type MulticastSubscription struct {

	ValGroupId string `json:"valGroupId"`

	AnncMode ServiceAnnoucementMode `json:"anncMode"`

	MultiQosReq string `json:"multiQosReq"`

	LocArea MbmsLocArea `json:"locArea,omitempty"`

	// string with format 'date-time' as defined in OpenAPI.
	Duration time.Time `json:"duration,omitempty"`

	// Integer where the allowed values correspond to the value range of an unsigned 32-bit integer. 
	Tmgi int32 `json:"tmgi,omitempty"`

	LocalMbmsInfo LocalMbmsInfo `json:"localMbmsInfo,omitempty"`

	LocalMbmsActInd bool `json:"localMbmsActInd,omitempty"`

	// String providing an URI formatted according to RFC 3986.
	NotifUri string `json:"notifUri"`

	ReqTestNotif bool `json:"reqTestNotif,omitempty"`

	WsNotifCfg WebsockNotifConfig `json:"wsNotifCfg,omitempty"`

	// A string used to indicate the features supported by an API that is used as defined in clause  6.6 in 3GPP TS 29.500. The string shall contain a bitmask indicating supported features in  hexadecimal representation Each character in the string shall take a value of \"0\" to \"9\",  \"a\" to \"f\" or \"A\" to \"F\" and shall represent the support of 4 features as described in  table 5.2.2-3. The most significant character representing the highest-numbered features shall  appear first in the string, and the character representing features 1 to 4 shall appear last  in the string. The list of features and their numbering (starting with 1) are defined  separately for each API. If the string contains a lower number of characters than there are  defined features for an API, all features that would be represented by characters that are not  present in the string are not supported. 
	SuppFeat string `json:"suppFeat,omitempty"`

	// String identifying a IPv4 address formatted in the 'dotted decimal' notation as defined in RFC 1166. 
	UpIpv4Addr string `json:"upIpv4Addr,omitempty"`

	UpIpv6Addr Ipv6Addr `json:"upIpv6Addr,omitempty"`

	// Unsigned integer with valid values between 0 and 65535.
	UpPortNum int32 `json:"upPortNum,omitempty"`

	RadioFreqs []int32 `json:"radioFreqs,omitempty"`
}