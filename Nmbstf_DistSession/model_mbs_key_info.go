/*
 * Nmbstf-distsession
 *
 * MBSTF Distribution Session Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_Nmbstf_DistSession

import (
	"time"
)

// MbsKeyInfo - MBS Security Key Data Structure
type MbsKeyInfo struct {

	// string with format 'bytes' as defined in OpenAPI
	KeyDomainId string `json:"keyDomainId"`

	// string with format 'bytes' as defined in OpenAPI
	MskId string `json:"mskId"`

	// string with format 'bytes' as defined in OpenAPI
	Msk string `json:"msk,omitempty"`

	// string with format 'date-time' as defined in OpenAPI.
	MskLifetime time.Time `json:"mskLifetime,omitempty"`

	// string with format 'bytes' as defined in OpenAPI
	MtkId string `json:"mtkId,omitempty"`

	// string with format 'bytes' as defined in OpenAPI
	Mtk string `json:"mtk,omitempty"`
}
