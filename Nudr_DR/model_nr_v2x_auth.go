/*
 * Nudr_DataRepository API OpenAPI file
 *
 * Unified Data Repository Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 2.2.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_Nudr_DR

// NrV2xAuth - Contains NR V2X services authorized information.
type NrV2xAuth struct {

	VehicleUeAuth UeAuth `json:"vehicleUeAuth,omitempty"`

	PedestrianUeAuth UeAuth `json:"pedestrianUeAuth,omitempty"`
}
