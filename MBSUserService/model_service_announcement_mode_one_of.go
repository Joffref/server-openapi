/*
 * 3gpp-mbs-us
 *
 * API for MBS User Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_MBSUserService

type ServiceAnnouncementModeOneOf string

// List of ServiceAnnouncementModeOneOf
const (
	VIA_MBS_5 ServiceAnnouncementModeOneOf = "VIA_MBS_5"
	VIA_MBS_DISTRIBUTION_SESSION ServiceAnnouncementModeOneOf = "VIA_MBS_DISTRIBUTION_SESSION"
	PASSED_BACK ServiceAnnouncementModeOneOf = "PASSED_BACK"
)