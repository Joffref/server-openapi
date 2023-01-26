/*
 * Nadrf_DataManagement
 *
 * ADRF Data Management Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_Nadrf_DataManagement

// EventSubscription1 - Represents a subscription to a single event.
type EventSubscription1 struct {

	Event SmfEvent `json:"event"`

	DnaiChgType DnaiChangeType `json:"dnaiChgType,omitempty"`

	DddTraDescriptors []DddTrafficDescriptor `json:"dddTraDescriptors,omitempty"`

	DddStati []DlDataDeliveryStatus `json:"dddStati,omitempty"`

	AppIds []string `json:"appIds,omitempty"`

	TargetPeriod TimeWindow `json:"targetPeriod,omitempty"`

	// Indicates the subscription for UE transaction dispersion collectionon, if it is included and set to \"true\". Default value is \"false\". 
	TransacDispInd bool `json:"transacDispInd,omitempty"`

	// Indicates Session Management Transaction metrics.
	TransacMetrics []TransactionMetric `json:"transacMetrics,omitempty"`

	UeIpAddr IpAddr `json:"ueIpAddr,omitempty"`
}
