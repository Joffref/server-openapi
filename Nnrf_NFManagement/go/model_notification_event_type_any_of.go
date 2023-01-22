/*
 * NRF NFManagement Service
 *
 * NRF NFManagement Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.3.0-alpha.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package Nnrf_NFManagement

type NotificationEventTypeAnyOf string

// List of NotificationEventTypeAnyOf
const (
	REGISTERED NotificationEventTypeAnyOf = "NF_REGISTERED"
	DEREGISTERED NotificationEventTypeAnyOf = "NF_DEREGISTERED"
	PROFILE_CHANGED NotificationEventTypeAnyOf = "NF_PROFILE_CHANGED"
)