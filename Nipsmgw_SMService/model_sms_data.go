/*
 * Nipsmgw_SMService Service API
 *
 * IP-SM-GW SMService.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_Nipsmgw_SMService

// SmsData - Information within request message for delivering SMS.
type SmsData struct {

	SmsPayload RefToBinaryData `json:"smsPayload"`
}
