/*
 * NR NRM
 *
 * OAS 3.0.1 specification of the NR NRM © 2020, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.
 *
 * API version: 17.7.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_NrNrm

type AlarmType string

// List of AlarmType
const (
	COMMUNICATIONS_ALARM AlarmType = "COMMUNICATIONS_ALARM"
	QUALITY_OF_SERVICE_ALARM AlarmType = "QUALITY_OF_SERVICE_ALARM"
	PROCESSING_ERROR_ALARM AlarmType = "PROCESSING_ERROR_ALARM"
	EQUIPMENT_ALARM AlarmType = "EQUIPMENT_ALARM"
	ENVIRONMENTAL_ALARM AlarmType = "ENVIRONMENTAL_ALARM"
	INTEGRITY_VIOLATION AlarmType = "INTEGRITY_VIOLATION"
	OPERATIONAL_VIOLATION AlarmType = "OPERATIONAL_VIOLATION"
	PHYSICAL_VIOLATION AlarmType = "PHYSICAL_VIOLATION"
	SECURITY_SERVICE_OR_MECHANISM_VIOLATION AlarmType = "SECURITY_SERVICE_OR_MECHANISM_VIOLATION"
	TIME_DOMAIN_VIOLATION AlarmType = "TIME_DOMAIN_VIOLATION"
)
