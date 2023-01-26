/*
 * Unified Data Repository Service API file for subscription data
 *
 * Unified Data Repository Service (subscription data).   The API version is defined in 3GPP TS 29.504.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: -
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_Subscription_Data

// SteeringInfo - Contains a combination of one PLMN identity and zero or more access technologies.
type SteeringInfo struct {

	PlmnId PlmnId `json:"plmnId"`

	AccessTechList []AccessTech `json:"accessTechList,omitempty"`
}