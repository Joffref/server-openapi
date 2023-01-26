/*
 * Nudr_DataRepository API OpenAPI file
 *
 * Unified Data Repository Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 2.2.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_Nudr_DR

type DnnConfiguration1 struct {

	PduSessionTypes PduSessionTypes1 `json:"pduSessionTypes"`

	SscModes SscModes1 `json:"sscModes"`

	IwkEpsInd bool `json:"iwkEpsInd,omitempty"`

	Var5gQosProfile SubscribedDefaultQos1 `json:"5gQosProfile,omitempty"`

	SessionAmbr Ambr1 `json:"sessionAmbr,omitempty"`

	Var3gppChargingCharacteristics string `json:"3gppChargingCharacteristics,omitempty"`

	StaticIpAddress []IpAddress1 `json:"staticIpAddress,omitempty"`

	UpSecurity UpSecurity1 `json:"upSecurity,omitempty"`

	PduSessionContinuityInd PduSessionContinuityInd `json:"pduSessionContinuityInd,omitempty"`

	// Identity of the NEF
	NiddNefId string `json:"niddNefId,omitempty"`

	NiddInfo NiddInformation1 `json:"niddInfo,omitempty"`

	RedundantSessionAllowed bool `json:"redundantSessionAllowed,omitempty"`

	AcsInfo AcsInfo1 `json:"acsInfo,omitempty"`

	Ipv4FrameRouteList []FrameRouteInfo1 `json:"ipv4FrameRouteList,omitempty"`

	Ipv6FrameRouteList []FrameRouteInfo1 `json:"ipv6FrameRouteList,omitempty"`

	AtsssAllowed bool `json:"atsssAllowed,omitempty"`

	SecondaryAuth bool `json:"secondaryAuth,omitempty"`

	UavSecondaryAuth bool `json:"uavSecondaryAuth,omitempty"`

	DnAaaIpAddressAllocation bool `json:"dnAaaIpAddressAllocation,omitempty"`

	DnAaaAddress IpAddress1 `json:"dnAaaAddress,omitempty"`

	AdditionalDnAaaAddresses []IpAddress1 `json:"additionalDnAaaAddresses,omitempty"`

	// Fully Qualified Domain Name
	DnAaaFqdn string `json:"dnAaaFqdn,omitempty"`

	IptvAccCtrlInfo string `json:"iptvAccCtrlInfo,omitempty"`

	Ipv4Index IpIndex `json:"ipv4Index,omitempty"`

	Ipv6Index IpIndex `json:"ipv6Index,omitempty"`

	EcsAddrConfigInfo *EcsAddrConfigInfo1 `json:"ecsAddrConfigInfo,omitempty"`

	AdditionalEcsAddrConfigInfos []EcsAddrConfigInfo1 `json:"additionalEcsAddrConfigInfos,omitempty"`

	SharedEcsAddrConfigInfo string `json:"sharedEcsAddrConfigInfo,omitempty"`

	AdditionalSharedEcsAddrConfigInfoIds []string `json:"additionalSharedEcsAddrConfigInfoIds,omitempty"`

	EasDiscoveryAuthorized bool `json:"easDiscoveryAuthorized,omitempty"`

	OnboardingInd bool `json:"onboardingInd,omitempty"`

	AerialUeInd AerialUeIndication `json:"aerialUeInd,omitempty"`

	SubscribedMaxIpv6PrefixSize int32 `json:"subscribedMaxIpv6PrefixSize,omitempty"`
}