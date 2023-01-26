/*
 * Nmbsmf-MBSSession
 *
 * MB-SMF MBSSession Service. © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved. 
 *
 * API version: 1.1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_Nmbsmf_MBSSession

// StatusSubscribeRspData - Data within StatusSubscribe Response
type StatusSubscribeRspData struct {

	Subscription MbsSessionSubscription `json:"subscription"`

	EventList MbsSessionEventReportList `json:"eventList,omitempty"`
}
