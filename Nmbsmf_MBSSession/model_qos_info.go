/*
 * Nmbsmf-MBSSession
 *
 * MB-SMF MBSSession Service. © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved. 
 *
 * API version: 1.1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_Nmbsmf_MBSSession

// QosInfo - QoS Information
type QosInfo struct {

	QosFlowsAddModRequestList []QosFlowAddModifyRequestItem `json:"qosFlowsAddModRequestList,omitempty"`

	QosFlowsRelRequestList []int32 `json:"qosFlowsRelRequestList,omitempty"`
}