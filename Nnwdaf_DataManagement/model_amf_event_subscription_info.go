/*
 * Nnwdaf_DataManagement
 *
 * Nnwdaf_DataManagement API Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_Nnwdaf_DataManagement

// AmfEventSubscriptionInfo - Individual AMF Event Subscription Information
type AmfEventSubscriptionInfo struct {

	// String providing an URI formatted according to RFC 3986.
	SubId string `json:"subId"`

	NotifyCorrelationId string `json:"notifyCorrelationId,omitempty"`

	RefIdList []int32 `json:"refIdList"`

	// String providing an URI formatted according to RFC 3986.
	OldSubId string `json:"oldSubId,omitempty"`
}
