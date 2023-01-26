/*
 * Nmfaf_3daDataManagement
 *
 * MFAF 3GPP DCCF Adaptor (3DA) Data Management Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_Nmfaf_3daDataManagement

type NotificationEventTypeAnyOf string

// List of NotificationEventTypeAnyOf
const (
	REGISTERED NotificationEventTypeAnyOf = "NF_REGISTERED"
	DEREGISTERED NotificationEventTypeAnyOf = "NF_DEREGISTERED"
	PROFILE_CHANGED NotificationEventTypeAnyOf = "NF_PROFILE_CHANGED"
)
