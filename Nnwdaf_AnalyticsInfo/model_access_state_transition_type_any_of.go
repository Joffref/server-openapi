/*
 * Nnwdaf_AnalyticsInfo
 *
 * Nnwdaf_AnalyticsInfo Service API.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.2.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_Nnwdaf_AnalyticsInfo

type AccessStateTransitionTypeAnyOf string

// List of AccessStateTransitionTypeAnyOf
const (
	ACCESS_TYPE_CHANGE_3_GPP AccessStateTransitionTypeAnyOf = "ACCESS_TYPE_CHANGE_3GPP"
	ACCESS_TYPE_CHANGE_N3_GPP AccessStateTransitionTypeAnyOf = "ACCESS_TYPE_CHANGE_N3GPP"
	RM_STATE_CHANGE_DEREGISTERED AccessStateTransitionTypeAnyOf = "RM_STATE_CHANGE_DEREGISTERED"
	RM_STATE_CHANGE_REGISTERED AccessStateTransitionTypeAnyOf = "RM_STATE_CHANGE_REGISTERED"
	CM_STATE_CHANGE_IDLE AccessStateTransitionTypeAnyOf = "CM_STATE_CHANGE_IDLE"
	CM_STATE_CHANGE_CONNECTED AccessStateTransitionTypeAnyOf = "CM_STATE_CHANGE_CONNECTED"
	HANDOVER AccessStateTransitionTypeAnyOf = "HANDOVER"
	MOBILITY_REGISTRATION_UPDATE AccessStateTransitionTypeAnyOf = "MOBILITY_REGISTRATION_UPDATE"
)
