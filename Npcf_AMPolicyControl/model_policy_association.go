/*
 * Npcf_AMPolicyControl
 *
 * Access and Mobility Policy Control Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.2.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_Npcf_AMPolicyControl

// PolicyAssociation - Represents an individual AM Policy Association resource.
type PolicyAssociation struct {

	Request PolicyAssociationRequest `json:"request,omitempty"`

	// Request Triggers that the PCF subscribes.
	Triggers []RequestTrigger `json:"triggers,omitempty"`

	ServAreaRes ServiceAreaRestriction `json:"servAreaRes,omitempty"`

	WlServAreaRes WirelineServiceAreaRestriction `json:"wlServAreaRes,omitempty"`

	// Unsigned integer representing the \"Subscriber Profile ID for RAT/Frequency Priority\"  as specified in 3GPP TS 36.413. 
	Rfsp int32 `json:"rfsp,omitempty"`

	// Unsigned integer representing the \"Subscriber Profile ID for RAT/Frequency Priority\"  as specified in 3GPP TS 36.413. 
	TargetRfsp int32 `json:"targetRfsp,omitempty"`

	SmfSelInfo *SmfSelectionData `json:"smfSelInfo,omitempty"`

	UeAmbr Ambr `json:"ueAmbr,omitempty"`

	// One or more UE-Slice-MBR(s) for S-NSSAI(s) of serving PLMN as part of the AMF Access and Mobility Policy as determined by the PCF. 
	UeSliceMbrs []UeSliceMbr `json:"ueSliceMbrs,omitempty"`

	// Contains the presence reporting area(s) for which reporting was requested. The praId attribute within the PresenceInfo data type is the key of the map. 
	Pras map[string]PresenceInfo `json:"pras,omitempty"`

	// A string used to indicate the features supported by an API that is used as defined in clause  6.6 in 3GPP TS 29.500. The string shall contain a bitmask indicating supported features in  hexadecimal representation Each character in the string shall take a value of \"0\" to \"9\",  \"a\" to \"f\" or \"A\" to \"F\" and shall represent the support of 4 features as described in  table 5.2.2-3. The most significant character representing the highest-numbered features shall  appear first in the string, and the character representing features 1 to 4 shall appear last  in the string. The list of features and their numbering (starting with 1) are defined  separately for each API. If the string contains a lower number of characters than there are  defined features for an API, all features that would be represented by characters that are not  present in the string are not supported. 
	SuppFeat string `json:"suppFeat"`

	PcfUeInfo *PcfUeCallbackInfo `json:"pcfUeInfo,omitempty"`

	MatchPdus *[]PduSessionInfo `json:"matchPdus,omitempty"`

	AsTimeDisParam *AsTimeDistributionParam `json:"asTimeDisParam,omitempty"`
}
