/*
 * Ndccf_DataManagement
 *
 * DCCF Data Management Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_Ndccf_DataManagement

// ParameterProcessingInstruction - Contains an event parameter name and the respective event parameter values and sets of attributes to be used in summarized reports. 
type ParameterProcessingInstruction struct {

	// A JSON pointer value that references an attribute within the notification object to which the processing instruction is applied. 
	Name string `json:"name"`

	// A list of values for the attribute identified by the name attribute.
	Values []interface{} `json:"values"`

	// Attributes requested to be used in the summarized reports.
	SumAttrs []SummarizationAttribute `json:"sumAttrs"`

	AggrLevel AggregationLevel `json:"aggrLevel,omitempty"`

	// Indicates the UEs for which processed reports are requested.
	Supis []string `json:"supis,omitempty"`

	// Indicates the Areas of Interest for which processed reports are requested.
	Areas []NetworkAreaInfo `json:"areas,omitempty"`
}
