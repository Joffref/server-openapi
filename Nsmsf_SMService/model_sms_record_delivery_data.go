/*
 * Nsmsf_SMService Service API
 *
 * SMSF SMService.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 2.2.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_Nsmsf_SMService

// SmsRecordDeliveryData - Represents information on the result of invoking the UplinkSMS service operation.
type SmsRecordDeliveryData struct {

	// Represents a record ID, used to identify a message carrying SMS payload.
	SmsRecordId string `json:"smsRecordId"`

	DeliveryStatus SmsDeliveryStatus `json:"deliveryStatus"`
}