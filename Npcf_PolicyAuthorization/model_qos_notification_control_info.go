/*
 * Npcf_PolicyAuthorization Service API
 *
 * PCF Policy Authorization Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.2.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_Npcf_PolicyAuthorization

// QosNotificationControlInfo - Indicates whether the QoS targets for a GRB flow are not guaranteed or guaranteed again.
type QosNotificationControlInfo struct {

	NotifType QosNotifType `json:"notifType"`

	Flows []Flows `json:"flows,omitempty"`

	AltSerReq string `json:"altSerReq,omitempty"`
}