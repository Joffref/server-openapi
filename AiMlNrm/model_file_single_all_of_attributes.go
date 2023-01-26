/*
 * AI/ML NRM
 *
 * OAS 3.0.1 specification of the AI/ML NRM © 2020, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.
 *
 * API version: 17.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_AiMlNrm

import (
	"time"
)

type FileSingleAllOfAttributes struct {

	FileLocation string `json:"fileLocation,omitempty"`

	FileCompression string `json:"fileCompression,omitempty"`

	FileSize int32 `json:"fileSize,omitempty"`

	FileDataType string `json:"fileDataType,omitempty"`

	FileFormat string `json:"fileFormat,omitempty"`

	FileReadyTime time.Time `json:"fileReadyTime,omitempty"`

	FileExpirationTime time.Time `json:"fileExpirationTime,omitempty"`

	FileContent string `json:"fileContent,omitempty"`

	JobRef string `json:"jobRef,omitempty"`

	JobId string `json:"jobId,omitempty"`
}
