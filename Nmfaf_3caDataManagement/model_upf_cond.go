/*
 * Nmfaf_3caDataManagement
 *
 * MFAF 3GPP Consumer Adaptor (3CA) Data Management Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_Nmfaf_3caDataManagement

// UpfCond - Subscription to a set of NF Instances (UPFs), able to serve a certain service area (i.e. SMF serving area or TAI list) 
type UpfCond struct {

	ConditionType string `json:"conditionType"`

	SmfServingArea []string `json:"smfServingArea,omitempty"`

	TaiList []Tai `json:"taiList,omitempty"`
}
