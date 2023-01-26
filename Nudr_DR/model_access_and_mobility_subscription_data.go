/*
 * Nudr_DataRepository API OpenAPI file
 *
 * Unified Data Repository Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 2.2.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_Nudr_DR

type AccessAndMobilitySubscriptionData struct {

	// A string used to indicate the features supported by an API that is used as defined in clause  6.6 in 3GPP TS 29.500. The string shall contain a bitmask indicating supported features in  hexadecimal representation Each character in the string shall take a value of \"0\" to \"9\",  \"a\" to \"f\" or \"A\" to \"F\" and shall represent the support of 4 features as described in  table 5.2.2-3. The most significant character representing the highest-numbered features shall  appear first in the string, and the character representing features 1 to 4 shall appear last  in the string. The list of features and their numbering (starting with 1) are defined  separately for each API. If the string contains a lower number of characters than there are  defined features for an API, all features that would be represented by characters that are not  present in the string are not supported. 
	SupportedFeatures string `json:"supportedFeatures,omitempty"`

	Gpsis []string `json:"gpsis,omitempty"`

	// Identifier of a group of NFs.
	HssGroupId string `json:"hssGroupId,omitempty"`

	InternalGroupIds []string `json:"internalGroupIds,omitempty"`

	// A map(list of key-value pairs) where GroupId serves as key of SharedDataId
	SharedVnGroupDataIds map[string]string `json:"sharedVnGroupDataIds,omitempty"`

	SubscribedUeAmbr AmbrRm `json:"subscribedUeAmbr,omitempty"`

	Nssai *Nssai `json:"nssai,omitempty"`

	RatRestrictions []RatType `json:"ratRestrictions,omitempty"`

	ForbiddenAreas []Area `json:"forbiddenAreas,omitempty"`

	ServiceAreaRestriction ServiceAreaRestriction `json:"serviceAreaRestriction,omitempty"`

	CoreNetworkTypeRestrictions []CoreNetworkType `json:"coreNetworkTypeRestrictions,omitempty"`

	// Unsigned integer representing the 'Subscriber Profile ID for RAT/Frequency Priority'  as specified in 3GPP TS 36.413 with the OpenAPI 'nullable: true' property.  
	RfspIndex *int32 `json:"rfspIndex,omitempty"`

	// indicating a time in seconds with OpenAPI defined 'nullable: true' property.
	SubsRegTimer *int32 `json:"subsRegTimer,omitempty"`

	UeUsageType int32 `json:"ueUsageType,omitempty"`

	MpsPriority bool `json:"mpsPriority,omitempty"`

	McsPriority bool `json:"mcsPriority,omitempty"`

	// indicating a time in seconds with OpenAPI defined 'nullable: true' property.
	ActiveTime *int32 `json:"activeTime,omitempty"`

	SorInfo SorInfo `json:"sorInfo,omitempty"`

	SorInfoExpectInd bool `json:"sorInfoExpectInd,omitempty"`

	SorafRetrieval bool `json:"sorafRetrieval,omitempty"`

	SorUpdateIndicatorList []SorUpdateIndicator `json:"sorUpdateIndicatorList,omitempty"`

	UpuInfo UpuInfo `json:"upuInfo,omitempty"`

	RoutingIndicator string `json:"routingIndicator,omitempty"`

	MicoAllowed bool `json:"micoAllowed,omitempty"`

	SharedAmDataIds []string `json:"sharedAmDataIds,omitempty"`

	OdbPacketServices OdbPacketServices `json:"odbPacketServices,omitempty"`

	SubscribedDnnList []AccessAndMobilitySubscriptionDataSubscribedDnnListInner `json:"subscribedDnnList,omitempty"`

	// indicating a time in seconds.
	ServiceGapTime int32 `json:"serviceGapTime,omitempty"`

	MdtUserConsent MdtUserConsent `json:"mdtUserConsent,omitempty"`

	MdtConfiguration MdtConfiguration `json:"mdtConfiguration,omitempty"`

	TraceData *TraceData1 `json:"traceData,omitempty"`

	CagData CagData `json:"cagData,omitempty"`

	// String representing the STN-SR as defined in clause 18.6 of 3GPP TS 23.003.
	StnSr string `json:"stnSr,omitempty"`

	// String representing the C-MSISDN as defined in clause 18.7 of 3GPP TS 23.003.
	CMsisdn string `json:"cMsisdn,omitempty"`

	NbIoTUePriority int32 `json:"nbIoTUePriority,omitempty"`

	NssaiInclusionAllowed bool `json:"nssaiInclusionAllowed,omitempty"`

	// string with format 'bytes' as defined in OpenAPI
	RgWirelineCharacteristics string `json:"rgWirelineCharacteristics,omitempty"`

	EcRestrictionDataWb EcRestrictionDataWb `json:"ecRestrictionDataWb,omitempty"`

	EcRestrictionDataNb bool `json:"ecRestrictionDataNb,omitempty"`

	ExpectedUeBehaviourList ExpectedUeBehaviourData `json:"expectedUeBehaviourList,omitempty"`

	PrimaryRatRestrictions []RatType `json:"primaryRatRestrictions,omitempty"`

	SecondaryRatRestrictions []RatType `json:"secondaryRatRestrictions,omitempty"`

	EdrxParametersList []EdrxParameters `json:"edrxParametersList,omitempty"`

	PtwParametersList []PtwParameters `json:"ptwParametersList,omitempty"`

	IabOperationAllowed bool `json:"iabOperationAllowed,omitempty"`

	// A map (list of key-value pairs where PlmnId serves as key) of PlmnRestriction
	AdjacentPlmnRestrictions map[string]PlmnRestriction `json:"adjacentPlmnRestrictions,omitempty"`

	WirelineForbiddenAreas []WirelineArea `json:"wirelineForbiddenAreas,omitempty"`

	WirelineServiceAreaRestriction WirelineServiceAreaRestriction `json:"wirelineServiceAreaRestriction,omitempty"`

	PcfSelectionAssistanceInfos []PcfSelectionAssistanceInfo `json:"pcfSelectionAssistanceInfos,omitempty"`

	AerialUeSubInfo AerialUeSubscriptionInfo `json:"aerialUeSubInfo,omitempty"`

	RoamingRestrictions RoamingRestrictions `json:"roamingRestrictions,omitempty"`

	RemoteProvInd bool `json:"remoteProvInd,omitempty"`
}
