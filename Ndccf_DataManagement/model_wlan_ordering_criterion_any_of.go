/*
 * Ndccf_DataManagement
 *
 * DCCF Data Management Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_Ndccf_DataManagement

type WlanOrderingCriterionAnyOf string

// List of WlanOrderingCriterionAnyOf
const (
	TIME_SLOT_START WlanOrderingCriterionAnyOf = "TIME_SLOT_START"
	NUMBER_OF_UES WlanOrderingCriterionAnyOf = "NUMBER_OF_UES"
	RSSI WlanOrderingCriterionAnyOf = "RSSI"
	RTT WlanOrderingCriterionAnyOf = "RTT"
	TRAFFIC_INFO WlanOrderingCriterionAnyOf = "TRAFFIC_INFO"
)
