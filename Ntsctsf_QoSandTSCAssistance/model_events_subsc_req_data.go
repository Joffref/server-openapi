/*
 * Ntsctsf_QoSandTSCAssistance Service API
 *
 * TSCTSF QoS and TSC Assistance Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_Ntsctsf_QoSandTSCAssistance

// EventsSubscReqData - Identifies the events the application subscribes to.
type EventsSubscReqData struct {

	Events []TscEvent `json:"events"`

	// String providing an URI formatted according to RFC 3986.
	NotifUri string `json:"notifUri"`

	QosMon QosMonitoringInformation `json:"qosMon,omitempty"`

	UsgThres UsageThreshold `json:"usgThres,omitempty"`

	NotifCorreId string `json:"notifCorreId"`
}
