/*
 * Nsmsf_SMService Service API
 *
 * SMSF SMService.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 2.2.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_Nsmsf_SMService

// CellGlobalId - Contains a Cell Global Identification as defined in 3GPP TS 23.003, clause 4.3.1.
type CellGlobalId struct {

	PlmnId PlmnId `json:"plmnId"`

	Lac string `json:"lac"`

	CellId string `json:"cellId"`
}
