/*
 * Nmbsmf-MBSSession
 *
 * MB-SMF MBSSession Service. © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved. 
 *
 * API version: 1.1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_Nmbsmf_MBSSession

import (
	"time"
)

// MbsSessionEventReport - MBS session event report
type MbsSessionEventReport struct {

	EventType MbsSessionEventType `json:"eventType"`

	// string with format 'date-time' as defined in OpenAPI.
	TimeStamp time.Time `json:"timeStamp,omitempty"`

	IngressTunAddrInfo IngressTunAddrInfo `json:"ingressTunAddrInfo,omitempty"`

	BroadcastDelStatus BroadcastDeliveryStatus `json:"broadcastDelStatus,omitempty"`
}