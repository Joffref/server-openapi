/*
 * Ndccf_ContextManagement
 *
 * DCCF Context Management Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_Ndccf_ContextManagement

type ScheduledCommunicationTypeAnyOf string

// List of ScheduledCommunicationTypeAnyOf
const (
	DOWNLINK_ONLY ScheduledCommunicationTypeAnyOf = "DOWNLINK_ONLY"
	UPLINK_ONLY ScheduledCommunicationTypeAnyOf = "UPLINK_ONLY"
	BIDIRECTIONAL ScheduledCommunicationTypeAnyOf = "BIDIRECTIONAL"
)