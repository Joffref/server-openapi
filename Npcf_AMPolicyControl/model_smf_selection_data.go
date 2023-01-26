/*
 * Npcf_AMPolicyControl
 *
 * Access and Mobility Policy Control Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.2.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_Npcf_AMPolicyControl

// SmfSelectionData - Represents the SMF Selection information that may be replaced by the PCF.
type SmfSelectionData struct {

	UnsuppDnn bool `json:"unsuppDnn,omitempty"`

	// Contains the list of DNNs per S-NSSAI that are candidates for replacement. The snssai attribute within the CandidateForReplacement data type is the key of the map. 
	Candidates *map[string]CandidateForReplacement `json:"candidates,omitempty"`

	Snssai Snssai `json:"snssai,omitempty"`

	MappingSnssai Snssai `json:"mappingSnssai,omitempty"`

	// String representing a Data Network as defined in clause 9A of 3GPP TS 23.003;  it shall contain either a DNN Network Identifier, or a full DNN with both the Network  Identifier and Operator Identifier, as specified in 3GPP TS 23.003 clause 9.1.1 and 9.1.2. It shall be coded as string in which the labels are separated by dots  (e.g. \"Label1.Label2.Label3\"). 
	Dnn string `json:"dnn,omitempty"`
}
