/*
 * Npcf_PolicyAuthorization Service API
 *
 * PCF Policy Authorization Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.2.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_Npcf_PolicyAuthorization

// AfEventNotification - Describes the event information delivered in the notification.
type AfEventNotification struct {

	Event AfEvent `json:"event"`

	Flows []Flows `json:"flows,omitempty"`
}
