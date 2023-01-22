/*
 * Nnwdaf_AnalyticsInfo
 *
 * Nnwdaf_AnalyticsInfo Service API.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.2.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package Nnwdaf_AnalyticsInfo

// ServiceNameListCond - Subscription to a set of NFs based on their support for a Service Name in the Servic Name list 
type ServiceNameListCond struct {

	ConditionType string `json:"conditionType"`

	ServiceNameList []ServiceName `json:"serviceNameList"`
}