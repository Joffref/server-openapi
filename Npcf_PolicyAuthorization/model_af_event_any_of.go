/*
 * Npcf_PolicyAuthorization Service API
 *
 * PCF Policy Authorization Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.2.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_Npcf_PolicyAuthorization

type AfEventAnyOf string

// List of AfEventAnyOf
const (
	ACCESS_TYPE_CHANGE AfEventAnyOf = "ACCESS_TYPE_CHANGE"
	ANI_REPORT AfEventAnyOf = "ANI_REPORT"
	APP_DETECTION AfEventAnyOf = "APP_DETECTION"
	CHARGING_CORRELATION AfEventAnyOf = "CHARGING_CORRELATION"
	EPS_FALLBACK AfEventAnyOf = "EPS_FALLBACK"
	FAILED_QOS_UPDATE AfEventAnyOf = "FAILED_QOS_UPDATE"
	FAILED_RESOURCES_ALLOCATION AfEventAnyOf = "FAILED_RESOURCES_ALLOCATION"
	OUT_OF_CREDIT AfEventAnyOf = "OUT_OF_CREDIT"
	PDU_SESSION_STATUS AfEventAnyOf = "PDU_SESSION_STATUS"
	PLMN_CHG AfEventAnyOf = "PLMN_CHG"
	QOS_MONITORING AfEventAnyOf = "QOS_MONITORING"
	QOS_NOTIF AfEventAnyOf = "QOS_NOTIF"
	RAN_NAS_CAUSE AfEventAnyOf = "RAN_NAS_CAUSE"
	REALLOCATION_OF_CREDIT AfEventAnyOf = "REALLOCATION_OF_CREDIT"
	SAT_CATEGORY_CHG AfEventAnyOf = "SAT_CATEGORY_CHG"
	SUCCESSFUL_QOS_UPDATE AfEventAnyOf = "SUCCESSFUL_QOS_UPDATE"
	SUCCESSFUL_RESOURCES_ALLOCATION AfEventAnyOf = "SUCCESSFUL_RESOURCES_ALLOCATION"
	TSN_BRIDGE_INFO AfEventAnyOf = "TSN_BRIDGE_INFO"
	UP_PATH_CHG_FAILURE AfEventAnyOf = "UP_PATH_CHG_FAILURE"
	USAGE_REPORT AfEventAnyOf = "USAGE_REPORT"
)
