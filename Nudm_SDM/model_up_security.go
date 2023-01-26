/*
 * Nudm_SDM
 *
 * Nudm Subscriber Data Management Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 2.2.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_Nudm_SDM

// UpSecurity - Contains Userplain security information.
type UpSecurity struct {

	UpIntegr UpIntegrity `json:"upIntegr"`

	UpConfid UpConfidentiality `json:"upConfid"`
}