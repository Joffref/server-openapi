/*
 * Namf_Communication
 *
 * AMF Communication Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.2.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_Namf_Communication

// PwsResponseData - Data related PWS included in a N2 Information Transfer response
type PwsResponseData struct {

	// Unsigned Integer, i.e. only value 0 and integers above 0 are permissible.
	NgapMessageType int32 `json:"ngapMessageType"`

	// Integer where the allowed values correspond to the value range of an unsigned 16-bit integer.
	SerialNumber int32 `json:"serialNumber"`

	MessageIdentifier int32 `json:"messageIdentifier"`

	UnknownTaiList []Tai `json:"unknownTaiList,omitempty"`
}