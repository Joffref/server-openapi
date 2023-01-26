/*
 * EES ACR Management Event_API
 *
 * API for EES ACR Management Event.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_Eees_ACRManagementEvent

type NotificationMethodAnyOf string

// List of NotificationMethodAnyOf
const (
	PERIODIC NotificationMethodAnyOf = "PERIODIC"
	ONE_TIME NotificationMethodAnyOf = "ONE_TIME"
	ON_EVENT_DETECTION NotificationMethodAnyOf = "ON_EVENT_DETECTION"
)
