/*
 * Unified Data Repository Service API file for subscription data
 *
 * Unified Data Repository Service (subscription data).   The API version is defined in 3GPP TS 29.504.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: -
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_Subscription_Data

type ScheduledCommunicationTypeAnyOf string

// List of ScheduledCommunicationTypeAnyOf
const (
	DOWNLINK_ONLY ScheduledCommunicationTypeAnyOf = "DOWNLINK_ONLY"
	UPLINK_ONLY ScheduledCommunicationTypeAnyOf = "UPLINK_ONLY"
	BIDIRECTIONAL ScheduledCommunicationTypeAnyOf = "BIDIRECTIONAL"
)
