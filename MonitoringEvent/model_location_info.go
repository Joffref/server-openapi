/*
 * 3gpp-monitoring-event
 *
 * API for Monitoring Event.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.2.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_MonitoringEvent

// LocationInfo - Represents the user location information.
type LocationInfo struct {

	// Unsigned integer identifying a period of time in units of minutes.
	AgeOfLocationInfo int32 `json:"ageOfLocationInfo,omitempty"`

	// Indicates the Cell Global Identification of the user which identifies the cell the UE is registered.
	CellId string `json:"cellId,omitempty"`

	// Indicates the eNodeB in which the UE is currently located.
	EnodeBId string `json:"enodeBId,omitempty"`

	// Identifies the Routing Area Identity of the user where the UE is located.
	RoutingAreaId string `json:"routingAreaId,omitempty"`

	// Identifies the Tracking Area Identity of the user where the UE is located.
	TrackingAreaId string `json:"trackingAreaId,omitempty"`

	// Identifies the PLMN Identity of the user where the UE is located.
	PlmnId string `json:"plmnId,omitempty"`

	// Identifies the TWAN Identity of the user where the UE is located.
	TwanId string `json:"twanId,omitempty"`

	GeographicArea GeographicArea `json:"geographicArea,omitempty"`

	CivicAddress CivicAddress `json:"civicAddress,omitempty"`

	PositionMethod PositioningMethod `json:"positionMethod,omitempty"`

	QosFulfilInd AccuracyFulfilmentIndicator `json:"qosFulfilInd,omitempty"`

	UeVelocity VelocityEstimate `json:"ueVelocity,omitempty"`

	LdrType LdrType `json:"ldrType,omitempty"`

	AchievedQos MinorLocationQoS `json:"achievedQos,omitempty"`
}
