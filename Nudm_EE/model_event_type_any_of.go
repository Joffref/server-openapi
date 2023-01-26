/*
 * Nudm_EE
 *
 * Nudm Event Exposure Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.2.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_Nudm_EE

type EventTypeAnyOf string

// List of EventTypeAnyOf
const (
	LOSS_OF_CONNECTIVITY EventTypeAnyOf = "LOSS_OF_CONNECTIVITY"
	UE_REACHABILITY_FOR_DATA EventTypeAnyOf = "UE_REACHABILITY_FOR_DATA"
	UE_REACHABILITY_FOR_SMS EventTypeAnyOf = "UE_REACHABILITY_FOR_SMS"
	LOCATION_REPORTING EventTypeAnyOf = "LOCATION_REPORTING"
	CHANGE_OF_SUPI_PEI_ASSOCIATION EventTypeAnyOf = "CHANGE_OF_SUPI_PEI_ASSOCIATION"
	ROAMING_STATUS EventTypeAnyOf = "ROAMING_STATUS"
	COMMUNICATION_FAILURE EventTypeAnyOf = "COMMUNICATION_FAILURE"
	AVAILABILITY_AFTER_DDN_FAILURE EventTypeAnyOf = "AVAILABILITY_AFTER_DDN_FAILURE"
	CN_TYPE_CHANGE EventTypeAnyOf = "CN_TYPE_CHANGE"
	DL_DATA_DELIVERY_STATUS EventTypeAnyOf = "DL_DATA_DELIVERY_STATUS"
	PDN_CONNECTIVITY_STATUS EventTypeAnyOf = "PDN_CONNECTIVITY_STATUS"
	UE_CONNECTION_MANAGEMENT_STATE EventTypeAnyOf = "UE_CONNECTION_MANAGEMENT_STATE"
	ACCESS_TYPE_REPORT EventTypeAnyOf = "ACCESS_TYPE_REPORT"
	REGISTRATION_STATE_REPORT EventTypeAnyOf = "REGISTRATION_STATE_REPORT"
	CONNECTIVITY_STATE_REPORT EventTypeAnyOf = "CONNECTIVITY_STATE_REPORT"
	TYPE_ALLOCATION_CODE_REPORT EventTypeAnyOf = "TYPE_ALLOCATION_CODE_REPORT"
	FREQUENT_MOBILITY_REGISTRATION_REPORT EventTypeAnyOf = "FREQUENT_MOBILITY_REGISTRATION_REPORT"
	PDU_SES_REL EventTypeAnyOf = "PDU_SES_REL"
	PDU_SES_EST EventTypeAnyOf = "PDU_SES_EST"
	UE_MEMORY_AVAILABLE_FOR_SMS EventTypeAnyOf = "UE_MEMORY_AVAILABLE_FOR_SMS"
)
