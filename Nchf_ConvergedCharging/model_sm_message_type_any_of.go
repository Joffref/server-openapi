/*
 * Nchf_ConvergedCharging
 *
 * ConvergedCharging Service    © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved. 
 *
 * API version: 3.1.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_Nchf_ConvergedCharging

type SmMessageTypeAnyOf string

// List of SmMessageTypeAnyOf
const (
	SUBMISSION SmMessageTypeAnyOf = "SUBMISSION"
	DELIVERY_REPORT SmMessageTypeAnyOf = "DELIVERY_REPORT"
	SM_SERVICE_REQUEST SmMessageTypeAnyOf = "SM_SERVICE_REQUEST"
	DELIVERY SmMessageTypeAnyOf = "DELIVERY"
)