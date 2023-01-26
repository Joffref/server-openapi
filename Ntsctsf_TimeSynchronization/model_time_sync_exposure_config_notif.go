/*
 * Ntsctsf_TimeSynchronization Service API
 *
 * TSCTSF Time Synchronization Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_Ntsctsf_TimeSynchronization

// TimeSyncExposureConfigNotif - Contains the notification of time synchronization service state.
type TimeSyncExposureConfigNotif struct {

	// Notification Correlation ID assigned by the NF service consumer.
	ConfigNotifId string `json:"configNotifId"`

	StateOfConfig StateOfConfiguration `json:"stateOfConfig"`
}
