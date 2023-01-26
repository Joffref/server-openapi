/*
 * Nhss_EE
 *
 * HSS Event Exposure.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_Nhss_EE

// PatchResult - The execution report result on failed modification.
type PatchResult struct {

	// The execution report contains an array of report items. Each report item indicates one  failed modification. 
	Report []ReportItem `json:"report"`
}
