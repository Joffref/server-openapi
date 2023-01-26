/*
 * NSSF NS Selection
 *
 * NSSF Network Slice Selection Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 2.2.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_Nnssf_NSSelection

// AuthorizedNetworkSliceInfo - Contains the authorized network slice information
type AuthorizedNetworkSliceInfo struct {

	AllowedNssaiList []AllowedNssai `json:"allowedNssaiList,omitempty"`

	ConfiguredNssai []ConfiguredSnssai `json:"configuredNssai,omitempty"`

	TargetAmfSet string `json:"targetAmfSet,omitempty"`

	CandidateAmfList []string `json:"candidateAmfList,omitempty"`

	RejectedNssaiInPlmn []Snssai `json:"rejectedNssaiInPlmn,omitempty"`

	RejectedNssaiInTa []Snssai `json:"rejectedNssaiInTa,omitempty"`

	NsiInformation NsiInformation `json:"nsiInformation,omitempty"`

	// A string used to indicate the features supported by an API that is used as defined in clause  6.6 in 3GPP TS 29.500. The string shall contain a bitmask indicating supported features in  hexadecimal representation Each character in the string shall take a value of \"0\" to \"9\",  \"a\" to \"f\" or \"A\" to \"F\" and shall represent the support of 4 features as described in  table 5.2.2-3. The most significant character representing the highest-numbered features shall  appear first in the string, and the character representing features 1 to 4 shall appear last  in the string. The list of features and their numbering (starting with 1) are defined  separately for each API. If the string contains a lower number of characters than there are  defined features for an API, all features that would be represented by characters that are not  present in the string are not supported. 
	SupportedFeatures string `json:"supportedFeatures,omitempty"`

	// String providing an URI formatted according to RFC 3986.
	NrfAmfSet string `json:"nrfAmfSet,omitempty"`

	// String providing an URI formatted according to RFC 3986.
	NrfAmfSetNfMgtUri string `json:"nrfAmfSetNfMgtUri,omitempty"`

	// String providing an URI formatted according to RFC 3986.
	NrfAmfSetAccessTokenUri string `json:"nrfAmfSetAccessTokenUri,omitempty"`

	// Map indicating whether the NRF requires Oauth2-based authorization for accessing its services. The key of the map shall be the name of an NRF service, e.g. \"nnrf-nfm\" or \"nnrf-disc\" 
	NrfOauth2Required map[string]bool `json:"nrfOauth2Required,omitempty"`

	// NF Service Set Identifier (see clause 28.12 of 3GPP TS 23.003) formatted as the following  string \"set<Set ID>.sn<Service Name>.nfi<NF Instance ID>.5gc.mnc<MNC>.mcc<MCC>\", or  \"set<SetID>.sn<ServiceName>.nfi<NFInstanceID>.5gc.nid<NID>.mnc<MNC>.mcc<MCC>\" with  <MCC> encoded as defined in clause 5.4.2 (\"Mcc\" data type definition)   <MNC> encoding the Mobile Network Code part of the PLMN, comprising 3 digits.    If there are only 2 significant digits in the MNC, one \"0\" digit shall be inserted    at the left side to fill the 3 digits coding of MNC.  Pattern: '^[0-9]{3}$' <NID> encoded as defined in clause 5.4.2 (\"Nid\" data type definition)  <NFInstanceId> encoded as defined in clause 5.3.2  <ServiceName> encoded as defined in 3GPP TS 29.510  <Set ID> encoded as a string of characters consisting of alphabetic    characters (A-Z and a-z), digits (0-9) and/or the hyphen (-) and that shall end    with either an alphabetic character or a digit. 
	TargetAmfServiceSet string `json:"targetAmfServiceSet,omitempty"`

	TargetNssai []Snssai `json:"targetNssai,omitempty"`

	NsagInfos []NsagInfo `json:"nsagInfos,omitempty"`
}