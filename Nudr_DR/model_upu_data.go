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

// UpuData - Used to store the status of the latest UPU data update.
type UpuData struct {

	// string with format 'date-time' as defined in OpenAPI.
	ProvisioningTime time.Time `json:"provisioningTime"`

	UeUpdateStatus UeUpdateStatus `json:"ueUpdateStatus"`

	// MAC value for protecting UPU procedure (UPU-MAC-IAUSF and UPU-MAC-IUE).
	UpuXmacIue string `json:"upuXmacIue,omitempty"`

	// MAC value for protecting UPU procedure (UPU-MAC-IAUSF and UPU-MAC-IUE).
	UpuMacIue string `json:"upuMacIue,omitempty"`
}
