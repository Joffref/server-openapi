/*
 * Nudr_DataRepository API OpenAPI file
 *
 * Unified Data Repository Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 2.2.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_Nudr_DR

import (
	"time"
)

type UpuInfo struct {

	UpuDataList []UpuData1 `json:"upuDataList,omitempty"`

	UpuRegInd bool `json:"upuRegInd,omitempty"`

	// Contains the indication of whether the acknowledgement from UE is needed.
	UpuAckInd bool `json:"upuAckInd,omitempty"`

	// MAC value for protecting UPU procedure (UPU-MAC-IAUSF and UPU-MAC-IUE).
	UpuMacIausf string `json:"upuMacIausf,omitempty"`

	// CounterUPU.
	CounterUpu string `json:"counterUpu,omitempty"`

	// string with format 'date-time' as defined in OpenAPI.
	ProvisioningTime time.Time `json:"provisioningTime"`

	// string with format 'bytes' as defined in OpenAPI
	UpuTransparentContainer string `json:"upuTransparentContainer,omitempty"`
}
