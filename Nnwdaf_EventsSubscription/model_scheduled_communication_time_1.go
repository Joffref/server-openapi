/*
 * Nnwdaf_EventsSubscription
 *
 * Nnwdaf_EventsSubscription Service API.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.2.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_Nnwdaf_EventsSubscription

// ScheduledCommunicationTime1 - Identifies time and day of the week when the UE is available for communication.
type ScheduledCommunicationTime1 struct {

	// Identifies the day(s) of the week. If absent, it indicates every day of the week. 
	DaysOfWeek []int32 `json:"daysOfWeek,omitempty"`

	// String with format partial-time or full-time as defined in clause 5.6 of IETF RFC 3339. Examples, 20:15:00, 20:15:00-08:00 (for 8 hours behind UTC).  
	TimeOfDayStart string `json:"timeOfDayStart,omitempty"`

	// String with format partial-time or full-time as defined in clause 5.6 of IETF RFC 3339. Examples, 20:15:00, 20:15:00-08:00 (for 8 hours behind UTC).  
	TimeOfDayEnd string `json:"timeOfDayEnd,omitempty"`
}
