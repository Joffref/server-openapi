/*
 * Ndccf_DataManagement
 *
 * DCCF Data Management Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_Ndccf_DataManagement

// RedundantTransmissionExpInfo - The redundant transmission experience related information. When subscribed event is  \"RED_TRANS_EXP\", the \"redTransInfos\" attribute shall be included. 
type RedundantTransmissionExpInfo struct {

	SpatialValidCon NetworkAreaInfo `json:"spatialValidCon,omitempty"`

	// String representing a Data Network as defined in clause 9A of 3GPP TS 23.003;  it shall contain either a DNN Network Identifier, or a full DNN with both the Network  Identifier and Operator Identifier, as specified in 3GPP TS 23.003 clause 9.1.1 and 9.1.2. It shall be coded as string in which the labels are separated by dots  (e.g. \"Label1.Label2.Label3\"). 
	Dnn string `json:"dnn,omitempty"`

	RedTransExps []RedundantTransmissionExpPerTs `json:"redTransExps"`
}
