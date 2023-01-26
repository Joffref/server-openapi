/*
 * NRF NFManagement Service
 *
 * NRF NFManagement Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.3.0-alpha.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_Nnrf_NFManagement

// NsacfCapability - NSACF service capabilities (e.g. to monitor and control the number of registered UEs or established PDU sessions per network slice) 
type NsacfCapability struct {

	// Indicates the service capability of the NSACF to monitor and control the number of registered UEs per network slice for the network slice that is subject to NSAC   true: Supported   false (default): Not Supported 
	SupportUeSAC bool `json:"supportUeSAC,omitempty"`

	// Indicates the service capability of the NSACF to monitor and control the number of established PDU sessions per network slice for the network slice that is subject to NSAC   true: Supported   false (default): Not Supported 
	SupportPduSAC bool `json:"supportPduSAC,omitempty"`
}
