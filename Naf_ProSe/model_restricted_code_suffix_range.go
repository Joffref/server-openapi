/*
 * Naf_ProSe API
 *
 * Naf_ProSe Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_Naf_ProSe

// RestrictedCodeSuffixRange - Contains a range of the Restricted Code Suffixes which are consecutive.
type RestrictedCodeSuffixRange struct {

	// Contains the ProSe Restricted Code Suffix.
	BeginningSuffix string `json:"beginningSuffix"`

	// Contains the ProSe Restricted Code Suffix.
	EndingSuffix string `json:"endingSuffix"`
}
