/*
 * Npcf_BDTPolicyControl Service API
 *
 * PCF BDT Policy Control Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.2.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_Npcf_BDTPolicyControl

// BdtPolicyDataPatch - A JSON Merge Patch body schema containing modification instruction to be performed on the bdtPolData attribute of the BdtPolicy data structure to select a transfer policy. Adds selTransPolicyId to BdtPolicyData data structure. 
type BdtPolicyDataPatch struct {

	// Contains an identity (i.e. transPolicyId value) of the selected transfer policy. If the BdtNotification_5G feature is supported value 0 indicates that no transfer policy is selected. 
	SelTransPolicyId int32 `json:"selTransPolicyId"`
}
