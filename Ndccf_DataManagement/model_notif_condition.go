/*
 * Ndccf_DataManagement
 *
 * DCCF Data Management Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_Ndccf_DataManagement

// NotifCondition - Condition (list of attributes in the NF Profile) to determine whether a notification must be sent by NRF 
type NotifCondition struct {

	MonitoredAttributes []string `json:"monitoredAttributes,omitempty"`

	UnmonitoredAttributes []string `json:"unmonitoredAttributes,omitempty"`
}
