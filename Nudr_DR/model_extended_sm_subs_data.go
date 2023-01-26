/*
 * Nudr_DataRepository API OpenAPI file
 *
 * Unified Data Repository Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 2.2.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_Nudr_DR

// ExtendedSmSubsData - Contains identifiers of shared Session Management Subscription Data and optionally individual Session Management Subscription Data.
type ExtendedSmSubsData struct {

	SharedSmSubsDataIds []string `json:"sharedSmSubsDataIds"`

	IndividualSmSubsData []SessionManagementSubscriptionData `json:"individualSmSubsData,omitempty"`
}
