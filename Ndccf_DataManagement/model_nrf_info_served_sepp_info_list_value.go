/*
 * Ndccf_DataManagement
 *
 * DCCF Data Management Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_Ndccf_DataManagement

type NrfInfoServedSeppInfoListValue struct {

	SeppPrefix string `json:"seppPrefix,omitempty"`

	// Port numbers for HTTP and HTTPS. The key of the map shall be \"http\" or \"https\". 
	SeppPorts map[string]int32 `json:"seppPorts,omitempty"`

	RemotePlmnList []PlmnId `json:"remotePlmnList,omitempty"`

	RemoteSnpnList []PlmnIdNid `json:"remoteSnpnList,omitempty"`
}
