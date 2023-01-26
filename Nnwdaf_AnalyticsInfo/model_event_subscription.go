/*
 * Nnwdaf_AnalyticsInfo
 *
 * Nnwdaf_AnalyticsInfo Service API.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.2.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_Nnwdaf_AnalyticsInfo

// EventSubscription - Represents a subscription to a single event.
type EventSubscription struct {

	// FALSE represents not applicable for all slices. TRUE represents applicable for all slices. 
	AnySlice bool `json:"anySlice,omitempty"`

	// Identification(s) of application to which the subscription applies.
	AppIds []string `json:"appIds,omitempty"`

	// Identification(s) of DNN to which the subscription applies.
	Dnns []string `json:"dnns,omitempty"`

	Dnais []string `json:"dnais,omitempty"`

	Event NwdafEvent `json:"event"`

	ExtraReportReq EventReportingRequirement `json:"extraReportReq,omitempty"`

	// Identification(s) of LADN DNN to indicate the LADN service area as the AOI.
	LadnDnns []string `json:"ladnDnns,omitempty"`

	// Indicates that the NWDAF shall report the corresponding network slice load level to the NF  service consumer where the load level of the network slice identified by snssais is  reached. 
	LoadLevelThreshold int32 `json:"loadLevelThreshold,omitempty"`

	NotificationMethod NotificationMethod `json:"notificationMethod,omitempty"`

	MatchingDir MatchingDirection `json:"matchingDir,omitempty"`

	// Shall be supplied in order to start reporting when an average load level is reached. 
	NfLoadLvlThds []ThresholdLevel `json:"nfLoadLvlThds,omitempty"`

	NfInstanceIds []string `json:"nfInstanceIds,omitempty"`

	NfSetIds []string `json:"nfSetIds,omitempty"`

	NfTypes []NfType `json:"nfTypes,omitempty"`

	NetworkArea NetworkAreaInfo `json:"networkArea,omitempty"`

	VisitedAreas []NetworkAreaInfo `json:"visitedAreas,omitempty"`

	// Unsigned Integer, i.e. only value 0 and integers above 0 are permissible.
	MaxTopAppUlNbr int32 `json:"maxTopAppUlNbr,omitempty"`

	// Unsigned Integer, i.e. only value 0 and integers above 0 are permissible.
	MaxTopAppDlNbr int32 `json:"maxTopAppDlNbr,omitempty"`

	NsiIdInfos []NsiIdInfo `json:"nsiIdInfos,omitempty"`

	NsiLevelThrds []int32 `json:"nsiLevelThrds,omitempty"`

	QosRequ QosRequirement `json:"qosRequ,omitempty"`

	QosFlowRetThds []RetainabilityThreshold `json:"qosFlowRetThds,omitempty"`

	RanUeThrouThds []string `json:"ranUeThrouThds,omitempty"`

	// indicating a time in seconds.
	RepetitionPeriod int32 `json:"repetitionPeriod,omitempty"`

	// Identification(s) of network slice to which the subscription applies. It corresponds to  snssais in the data model definition of 3GPP TS 29.520.  
	Snssaia []Snssai `json:"snssaia,omitempty"`

	TgtUe TargetUeInformation `json:"tgtUe,omitempty"`

	CongThresholds []ThresholdLevel `json:"congThresholds,omitempty"`

	NwPerfRequs []NetworkPerfRequirement `json:"nwPerfRequs,omitempty"`

	BwRequs []BwRequirement `json:"bwRequs,omitempty"`

	ExcepRequs []Exception `json:"excepRequs,omitempty"`

	ExptAnaType ExpectedAnalyticsType `json:"exptAnaType,omitempty"`

	ExptUeBehav ExpectedUeBehaviourData `json:"exptUeBehav,omitempty"`

	RatFreqs []RatFreqInformation `json:"ratFreqs,omitempty"`

	ListOfAnaSubsets []AnalyticsSubset `json:"listOfAnaSubsets,omitempty"`

	DisperReqs []DispersionRequirement `json:"disperReqs,omitempty"`

	RedTransReqs []RedundantTransmissionExpReq `json:"redTransReqs,omitempty"`

	WlanReqs []WlanPerformanceReq `json:"wlanReqs,omitempty"`

	UpfInfo UpfInformation `json:"upfInfo,omitempty"`

	AppServerAddrs []AddrFqdn `json:"appServerAddrs,omitempty"`

	DnPerfReqs []DnPerformanceReq `json:"dnPerfReqs,omitempty"`
}
