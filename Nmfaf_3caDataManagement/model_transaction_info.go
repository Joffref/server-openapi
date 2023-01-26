/*
 * Nmfaf_3caDataManagement
 *
 * MFAF 3GPP Consumer Adaptor (3CA) Data Management Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_Nmfaf_3caDataManagement

// TransactionInfo - Represents SMF Transaction Information.
type TransactionInfo struct {

	// Unsigned Integer, i.e. only value 0 and integers above 0 are permissible.
	Transaction int32 `json:"transaction"`

	Snssai Snssai `json:"snssai,omitempty"`

	AppIds []string `json:"appIds,omitempty"`

	TransacMetrics []TransactionMetric `json:"transacMetrics,omitempty"`
}
