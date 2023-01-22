/*
 * NRF NFManagement Service
 *
 * NRF NFManagement Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.3.0-alpha.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package Nnrf_NFManagement

type NotificationTypeAnyOf string

// List of NotificationTypeAnyOf
const (
	N1_MESSAGES NotificationTypeAnyOf = "N1_MESSAGES"
	N2_INFORMATION NotificationTypeAnyOf = "N2_INFORMATION"
	LOCATION_NOTIFICATION NotificationTypeAnyOf = "LOCATION_NOTIFICATION"
	DATA_REMOVAL_NOTIFICATION NotificationTypeAnyOf = "DATA_REMOVAL_NOTIFICATION"
	DATA_CHANGE_NOTIFICATION NotificationTypeAnyOf = "DATA_CHANGE_NOTIFICATION"
	LOCATION_UPDATE_NOTIFICATION NotificationTypeAnyOf = "LOCATION_UPDATE_NOTIFICATION"
	NSSAA_REAUTH_NOTIFICATION NotificationTypeAnyOf = "NSSAA_REAUTH_NOTIFICATION"
	NSSAA_REVOC_NOTIFICATION NotificationTypeAnyOf = "NSSAA_REVOC_NOTIFICATION"
	MATCH_INFO_NOTIFICATION NotificationTypeAnyOf = "MATCH_INFO_NOTIFICATION"
	DATA_RESTORATION_NOTIFICATION NotificationTypeAnyOf = "DATA_RESTORATION_NOTIFICATION"
	TSCTS_NOTIFICATION NotificationTypeAnyOf = "TSCTS_NOTIFICATION"
)