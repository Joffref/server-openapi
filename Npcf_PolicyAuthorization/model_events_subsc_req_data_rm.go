/*
 * Npcf_PolicyAuthorization Service API
 *
 * PCF Policy Authorization Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.2.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_Npcf_PolicyAuthorization

// EventsSubscReqDataRm - This data type is defined in the same way as the EventsSubscReqData data type, but with the OpenAPI nullable property set to true.
type EventsSubscReqDataRm struct {

	Events []AfEventSubscription `json:"events"`

	// String providing an URI formatted according to RFC 3986.
	NotifUri string `json:"notifUri,omitempty"`

	ReqQosMonParams []RequestedQosMonitoringParameter `json:"reqQosMonParams,omitempty"`

	QosMon *QosMonitoringInformationRm `json:"qosMon,omitempty"`

	ReqAnis []RequiredAccessInfo `json:"reqAnis,omitempty"`

	UsgThres *UsageThresholdRm `json:"usgThres,omitempty"`

	NotifCorreId string `json:"notifCorreId,omitempty"`

	DirectNotifInd *bool `json:"directNotifInd,omitempty"`
}
