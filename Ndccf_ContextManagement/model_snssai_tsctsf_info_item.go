/*
 * Ndccf_ContextManagement
 *
 * DCCF Context Management Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_Ndccf_ContextManagement

// SnssaiTsctsfInfoItem - Set of parameters supported by TSCTSF for a given S-NSSAI
type SnssaiTsctsfInfoItem struct {

	SNssai ExtSnssai `json:"sNssai"`

	DnnInfoList []DnnTsctsfInfoItem `json:"dnnInfoList"`
}
