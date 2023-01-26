/*
 * NRF NFManagement Service
 *
 * NRF NFManagement Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.3.0-alpha.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_Nnrf_NFManagement

// SnssaiMbSmfInfoItem - Parameters supported by an MB-SMF for a given S-NSSAI
type SnssaiMbSmfInfoItem struct {

	SNssai ExtSnssai `json:"sNssai"`

	DnnInfoList []DnnMbSmfInfoItem `json:"dnnInfoList"`
}
