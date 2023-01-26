/*
 * Namf_Communication
 *
 * AMF Communication Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.2.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_Namf_Communication

// AnalyticsMetadataIndication - Contains analytics metadata information requested to be used during analytics generation. 
type AnalyticsMetadataIndication struct {

	DataWindow TimeWindow `json:"dataWindow,omitempty"`

	DataStatProps []DatasetStatisticalProperty `json:"dataStatProps,omitempty"`

	Strategy OutputStrategy `json:"strategy,omitempty"`

	AggrNwdafIds []string `json:"aggrNwdafIds,omitempty"`
}
