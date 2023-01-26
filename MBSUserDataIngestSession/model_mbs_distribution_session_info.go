/*
 * 3gpp-mbs-ud-ingest
 *
 * API for MBS User Data Ingest Session.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_MBSUserDataIngestSession

// MbsDistributionSessionInfo - Represents MBS Distribution Session information.
type MbsDistributionSessionInfo struct {

	MbsDistSessionId string `json:"mbsDistSessionId,omitempty"`

	MbsSessionId MbsSessionId `json:"mbsSessionId,omitempty"`

	MbsServInfo MbsServiceInfo `json:"mbsServInfo,omitempty"`

	// String representing a bit rate; the prefixes follow the standard symbols from The International System of Units, and represent x1000 multipliers, with the exception that prefix \"K\" is used to represent the standard symbol \"k\". 
	MaxContBitRate string `json:"maxContBitRate"`

	// Unsigned integer indicating Packet Delay Budget (see clauses 5.7.3.4 and 5.7.4 of 3GPP TS 23.501), expressed in milliseconds. 
	MaxContDelay int32 `json:"maxContDelay,omitempty"`

	DistrMethod DistributionMethod `json:"distrMethod"`

	FecConfig FecConfig `json:"fecConfig,omitempty"`

	ObjDistrInfo ObjectDistrMethInfo `json:"objDistrInfo,omitempty"`

	PckDistrInfo PacketDistrMethInfo `json:"pckDistrInfo,omitempty"`

	TrafficMarkingInfo string `json:"trafficMarkingInfo,omitempty"`

	MbsDistSessState DistSessionState `json:"mbsDistSessState,omitempty"`

	TgtServAreas MbsServiceArea `json:"tgtServAreas,omitempty"`

	ExtTgtServAreas ExternalMbsServiceArea `json:"extTgtServAreas,omitempty"`

	// MBS Frequency Selection Area Identifier
	MbsFSAId string `json:"mbsFSAId,omitempty"`

	// Represents an indication that this MBS Distribution Session belongs to a location- dependent MBS. This attribute shall be set to \"true\" to indicate that the MBS  Distribution Session belongs to a location-dependent MBS; or set to \"false\" to  indicate that the MBS Distribution Session does not belong to a location-dependent MBS. The default value is \"false\", if omitted. 
	LocationDependent bool `json:"locationDependent,omitempty"`

	// Represents an indication that this MBS Distribution Session belongs to a multiplex, i.e.  forms part of a set of MBS Distribution Sessions under the same parent MBS User Data  Ingest Session with identical or empty sets of target service areas and multiplexed onto  the same MBS Session at the MB-SMF. 
	MultiplexedServFlag bool `json:"multiplexedServFlag,omitempty"`

	// Represents an indication that this MBS Distribution Session is not open to any UE, i.e.  restricted to a set of UEs according to their MBS related subscription information. This attribute may be included only if the parent MBS User Service is of Multicast service type. This attribute shall be set to \"true\" to indicate that this MBS Distribution Session is restricted to a set of UE(s); or set to \"false\" to indicate that this MBS Distribution Session is open to any UE. The default value is \"false\", if omitted. 
	RestrictedFlag bool `json:"restrictedFlag,omitempty"`
}
