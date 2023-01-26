/*
 * Npcf_AMPolicyControl
 *
 * Access and Mobility Policy Control Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.2.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_Npcf_AMPolicyControl

// RequestTrigger - Possible values are: - LOC_CH: Location change (tracking area). The tracking area of the UE has changed. - PRA_CH: Change of UE presence in PRA. The AMF reports the current presence status   of the UE in a Presence Reporting Area, and notifies that the UE enters/leaves the   Presence Reporting Area. - SERV_AREA_CH: Service Area Restriction change. The UDM notifies the AMF that the   subscribed service area restriction information has changed. - RFSP_CH: RFSP index change. The UDM notifies the AMF that the subscribed RFSP index has   changed. - ALLOWED_NSSAI_CH: Allowed NSSAI change. The AMF notifies that the set of UE allowed   S-NSSAIs has changed. - UE_AMBR_CH: UE-AMBR change. The UDM notifies the AMF that the subscribed UE-AMBR has   changed. - SMF_SELECT_CH: SMF selection information change. The UE requested for an unsupported   DNN or UE requested for a DNN within the list of DNN candidates for replacement per   S-NSSAI. - ACCESS_TYPE_CH: Access Type change. The AMF notifies that the access type and the RAT   type combinations available in the AMF for a UE with simultaneous 3GPP and non-3GPP   connectivity has changed.  - UE_SLICE_MBR_CH: UE-Slice-MBR change. The NF service consumer notifies any changes    in the subscribed UE-Slice-MBR for each subscribed S-NSSAI of the home PLMN mapping    to a S-NSSAI of the serving PLMN. - NWDAF_DATA_CH: NDWAF DATA CHANGE. The AMF notifies that the NWDAF instance IDs used   for the UE and/or associated Analytics IDs used for the UE and available in the AMF   have changed. - TARGET_NSSAI: Generation of Target NSSAI. The NF service consumer notifies that the   Target NSSAI was generated. 
type RequestTrigger struct {
}
