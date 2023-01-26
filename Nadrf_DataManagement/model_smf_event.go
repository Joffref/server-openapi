/*
 * Nadrf_DataManagement
 *
 * ADRF Data Management Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_Nadrf_DataManagement

// SmfEvent - Possible values are: - AC_TY_CH: Access Type Change - UP_PATH_CH: UP Path Change - PDU_SES_REL: PDU Session Release - PLMN_CH: PLMN Change - UE_IP_CH: UE IP address change - RAT_TY_CH: RAT Type Change - DDDS: Downlink data delivery status - COMM_FAIL: Communication Failure - PDU_SES_EST: PDU Session Establishment - QFI_ALLOC: QFI allocation - QOS_MON: QoS Monitoring - SMCC_EXP: SM congestion control experience for PDU Session - DISPERSION: Session Management transaction dispersion - RED_TRANS_EXP: Redundant transmission experience for PDU Session - WLAN_INFO: WLAN information on PDU session for which Access Type is NON_3GPP_ACCESS and   RAT Type is TRUSTED_WLAN - UPF_INFO: The UPF information, including the UPF ID/address/FQDN information. 
type SmfEvent struct {
}
