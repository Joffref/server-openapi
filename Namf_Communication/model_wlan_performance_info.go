/*
 * Namf_Communication
 *
 * AMF Communication Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.2.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_Namf_Communication

// WlanPerformanceInfo - The WLAN performance related information.
type WlanPerformanceInfo struct {

	NetworkArea NetworkAreaInfo `json:"networkArea,omitempty"`

	WlanPerSsidInfos []WlanPerSsIdPerformanceInfo `json:"wlanPerSsidInfos"`
}