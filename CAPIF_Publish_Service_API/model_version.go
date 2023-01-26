/*
 * CAPIF_Publish_Service_API
 *
 * API for publishing service APIs.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.2.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_CAPIF_Publish_Service_API

import (
	"time"
)

// Version - Represents the API version information.
type Version struct {

	// API major version in URI (e.g. v1)
	ApiVersion string `json:"apiVersion"`

	// string with format \"date-time\" as defined in OpenAPI.
	Expiry time.Time `json:"expiry,omitempty"`

	// Resources supported by the API.
	Resources []Resource `json:"resources,omitempty"`

	// Custom operations without resource association.
	CustOperations []CustomOperation `json:"custOperations,omitempty"`
}
