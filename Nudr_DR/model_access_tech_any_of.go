/*
 * Nudr_DataRepository API OpenAPI file
 *
 * Unified Data Repository Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 2.2.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_Nudr_DR

type AccessTechAnyOf string

// List of AccessTechAnyOf
const (
	NR AccessTechAnyOf = "NR"
	EUTRAN_IN_WBS1_MODE_AND_NBS1_MODE AccessTechAnyOf = "EUTRAN_IN_WBS1_MODE_AND_NBS1_MODE"
	EUTRAN_IN_NBS1_MODE_ONLY AccessTechAnyOf = "EUTRAN_IN_NBS1_MODE_ONLY"
	EUTRAN_IN_WBS1_MODE_ONLY AccessTechAnyOf = "EUTRAN_IN_WBS1_MODE_ONLY"
	UTRAN AccessTechAnyOf = "UTRAN"
	GSM_AND_ECGSM_IO_T AccessTechAnyOf = "GSM_AND_ECGSM_IoT"
	GSM_WITHOUT_ECGSM_IO_T AccessTechAnyOf = "GSM_WITHOUT_ECGSM_IoT"
	ECGSM_IO_T_ONLY AccessTechAnyOf = "ECGSM_IoT_ONLY"
	CDMA_1X_RTT AccessTechAnyOf = "CDMA_1xRTT"
	CDMA_HRPD AccessTechAnyOf = "CDMA_HRPD"
	GSM_COMPACT AccessTechAnyOf = "GSM_COMPACT"
)
