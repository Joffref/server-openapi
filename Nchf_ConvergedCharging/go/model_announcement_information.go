/*
 * Nchf_ConvergedCharging
 *
 * ConvergedCharging Service    © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved. 
 *
 * API version: 3.1.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package Nchf_ConvergedCharging

type AnnouncementInformation struct {

	// Integer where the allowed values correspond to the value range of an unsigned 32-bit integer. 
	AnnouncementIdentifier int32 `json:"announcementIdentifier,omitempty"`

	// String providing an URI formatted according to RFC 3986.
	AnnouncementReference string `json:"announcementReference,omitempty"`

	VariableParts []VariablePart `json:"variableParts,omitempty"`

	// indicating a time in seconds.
	TimeToPlay int32 `json:"timeToPlay,omitempty"`

	QuotaConsumptionIndicator QuotaConsumptionIndicator `json:"quotaConsumptionIndicator,omitempty"`

	// Integer where the allowed values correspond to the value range of an unsigned 32-bit integer. 
	AnnouncementPriority int32 `json:"announcementPriority,omitempty"`

	PlayToParty PlayToParty `json:"playToParty,omitempty"`

	AnnouncementPrivacyIndicator AnnouncementPrivacyIndicator `json:"announcementPrivacyIndicator,omitempty"`

	Language string `json:"Language,omitempty"`
}