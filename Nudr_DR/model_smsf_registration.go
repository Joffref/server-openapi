/*
 * Nudr_DataRepository API OpenAPI file
 *
 * Unified Data Repository Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 2.2.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_Nudr_DR

import (
	"time"
)

type SmsfRegistration struct {

	// String uniquely identifying a NF instance. The format of the NF Instance ID shall be a  Universally Unique Identifier (UUID) version 4, as described in IETF RFC 4122.  
	SmsfInstanceId string `json:"smsfInstanceId"`

	// NF Set Identifier (see clause 28.12 of 3GPP TS 23.003), formatted as the following string \"set<Set ID>.<nftype>set.5gc.mnc<MNC>.mcc<MCC>\", or  \"set<SetID>.<NFType>set.5gc.nid<NID>.mnc<MNC>.mcc<MCC>\" with  <MCC> encoded as defined in clause 5.4.2 (\"Mcc\" data type definition)  <MNC> encoding the Mobile Network Code part of the PLMN, comprising 3 digits.    If there are only 2 significant digits in the MNC, one \"0\" digit shall be inserted    at the left side to fill the 3 digits coding of MNC.  Pattern: '^[0-9]{3}$' <NFType> encoded as a value defined in Table 6.1.6.3.3-1 of 3GPP TS 29.510 but    with lower case characters <Set ID> encoded as a string of characters consisting of    alphabetic characters (A-Z and a-z), digits (0-9) and/or the hyphen (-) and that    shall end with either an alphabetic character or a digit.  
	SmsfSetId string `json:"smsfSetId,omitempty"`

	// A string used to indicate the features supported by an API that is used as defined in clause  6.6 in 3GPP TS 29.500. The string shall contain a bitmask indicating supported features in  hexadecimal representation Each character in the string shall take a value of \"0\" to \"9\",  \"a\" to \"f\" or \"A\" to \"F\" and shall represent the support of 4 features as described in  table 5.2.2-3. The most significant character representing the highest-numbered features shall  appear first in the string, and the character representing features 1 to 4 shall appear last  in the string. The list of features and their numbering (starting with 1) are defined  separately for each API. If the string contains a lower number of characters than there are  defined features for an API, all features that would be represented by characters that are not  present in the string are not supported. 
	SupportedFeatures string `json:"supportedFeatures,omitempty"`

	PlmnId PlmnId `json:"plmnId"`

	SmsfMAPAddress string `json:"smsfMAPAddress,omitempty"`

	SmsfDiameterAddress NetworkNodeDiameterAddress `json:"smsfDiameterAddress,omitempty"`

	// string with format 'date-time' as defined in OpenAPI.
	RegistrationTime time.Time `json:"registrationTime,omitempty"`

	ContextInfo ContextInfo `json:"contextInfo,omitempty"`

	// String providing an URI formatted according to RFC 3986.
	DataRestorationCallbackUri string `json:"dataRestorationCallbackUri,omitempty"`

	ResetIds []string `json:"resetIds,omitempty"`

	SmsfSbiSupInd bool `json:"smsfSbiSupInd,omitempty"`

	UdrRestartInd bool `json:"udrRestartInd,omitempty"`

	// string with format 'date-time' as defined in OpenAPI.
	LastSynchronizationTime time.Time `json:"lastSynchronizationTime,omitempty"`
}
