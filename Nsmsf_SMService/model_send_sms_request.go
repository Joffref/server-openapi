/*
 * Nsmsf_SMService Service API
 *
 * SMSF SMService.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 2.2.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_Nsmsf_SMService

import (
	"os"
)

type SendSmsRequest struct {

	JsonData SmsRecordData `json:"jsonData,omitempty"`

	BinaryPayload *os.File `json:"binaryPayload,omitempty"`
}
