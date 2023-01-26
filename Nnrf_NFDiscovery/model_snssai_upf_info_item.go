/*
 * NRF NFDiscovery Service
 *
 * NRF NFDiscovery Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.3.0-alpha.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_Nnrf_NFDiscovery

// SnssaiUpfInfoItem - Set of parameters supported by UPF for a given S-NSSAI
type SnssaiUpfInfoItem struct {

	SNssai ExtSnssai `json:"sNssai"`

	DnnUpfInfoList []DnnUpfInfoItem `json:"dnnUpfInfoList"`

	RedundantTransport bool `json:"redundantTransport,omitempty"`
}
