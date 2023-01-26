/*
 * Fault Supervision MnS
 *
 * OAS 3.0.1 definition of the Fault Supervision MnS © 2021, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.
 *
 * API version: 17.1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_FaultMnS

type PerceivedSeverity string

// List of PerceivedSeverity
const (
	INDETERMINATE PerceivedSeverity = "INDETERMINATE"
	CRITICAL PerceivedSeverity = "CRITICAL"
	MAJOR PerceivedSeverity = "MAJOR"
	MINOR PerceivedSeverity = "MINOR"
	WARNING PerceivedSeverity = "WARNING"
	CLEARED PerceivedSeverity = "CLEARED"
)
