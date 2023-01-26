/*
 * Nnwdaf_DataManagement
 *
 * Nnwdaf_DataManagement API Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_Nnwdaf_DataManagement

// SnssaiInfoItem - Parameters supported by an NF for a given S-NSSAI Set of parameters supported by NF for a given S-NSSAI 
type SnssaiInfoItem struct {

	SNssai ExtSnssai `json:"sNssai"`

	DnnInfoList []DnnInfoItem `json:"dnnInfoList"`
}
