/*
 * Nsmf_PDUSession
 *
 * SMF PDU Session Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.3.0-alpha.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_Nsmf_PDUSession

// SmContextStatusNotification - Data within Notify SM Context Status Request
type SmContextStatusNotification struct {

	StatusInfo StatusInfo `json:"statusInfo"`

	SmallDataRateStatus SmallDataRateStatus `json:"smallDataRateStatus,omitempty"`

	ApnRateStatus ApnRateStatus `json:"apnRateStatus,omitempty"`

	DdnFailureStatus bool `json:"ddnFailureStatus,omitempty"`

	NotifyCorrelationIdsForddnFailure []string `json:"notifyCorrelationIdsForddnFailure,omitempty"`

	// String uniquely identifying a NF instance. The format of the NF Instance ID shall be a  Universally Unique Identifier (UUID) version 4, as described in IETF RFC 4122.  
	NewIntermediateSmfId string `json:"newIntermediateSmfId,omitempty"`

	// String uniquely identifying a NF instance. The format of the NF Instance ID shall be a  Universally Unique Identifier (UUID) version 4, as described in IETF RFC 4122.  
	NewSmfId string `json:"newSmfId,omitempty"`

	// NF Set Identifier (see clause 28.12 of 3GPP TS 23.003), formatted as the following string \"set<Set ID>.<nftype>set.5gc.mnc<MNC>.mcc<MCC>\", or  \"set<SetID>.<NFType>set.5gc.nid<NID>.mnc<MNC>.mcc<MCC>\" with  <MCC> encoded as defined in clause 5.4.2 (\"Mcc\" data type definition)  <MNC> encoding the Mobile Network Code part of the PLMN, comprising 3 digits.    If there are only 2 significant digits in the MNC, one \"0\" digit shall be inserted    at the left side to fill the 3 digits coding of MNC.  Pattern: '^[0-9]{3}$' <NFType> encoded as a value defined in Table 6.1.6.3.3-1 of 3GPP TS 29.510 but    with lower case characters <Set ID> encoded as a string of characters consisting of    alphabetic characters (A-Z and a-z), digits (0-9) and/or the hyphen (-) and that    shall end with either an alphabetic character or a digit.  
	NewSmfSetId string `json:"newSmfSetId,omitempty"`

	// String uniquely identifying a NF instance. The format of the NF Instance ID shall be a  Universally Unique Identifier (UUID) version 4, as described in IETF RFC 4122.  
	OldSmfId string `json:"oldSmfId,omitempty"`

	// String providing an URI formatted according to RFC 3986.
	OldSmContextRef string `json:"oldSmContextRef,omitempty"`

	// String providing an URI formatted according to RFC 3986.
	AltAnchorSmfUri string `json:"altAnchorSmfUri,omitempty"`

	// String uniquely identifying a NF instance. The format of the NF Instance ID shall be a  Universally Unique Identifier (UUID) version 4, as described in IETF RFC 4122.  
	AltAnchorSmfId string `json:"altAnchorSmfId,omitempty"`

	TargetDnaiInfo TargetDnaiInfo `json:"targetDnaiInfo,omitempty"`

	// String providing an URI formatted according to RFC 3986.
	OldPduSessionRef string `json:"oldPduSessionRef,omitempty"`

	// String providing an URI formatted according to RFC 3986.
	InterPlmnApiRoot string `json:"interPlmnApiRoot,omitempty"`
}
