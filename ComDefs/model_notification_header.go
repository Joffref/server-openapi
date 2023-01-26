/*
 * Common Type Definitions
 *
 * OAS 3.0.1 specification of common type definitions in the Generic NRM © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.
 *
 * API version: 17.2.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_ComDefs

import (
	"time"
)

type NotificationHeader struct {

	Href string `json:"href"`

	NotificationId int32 `json:"notificationId"`

	NotificationType NotificationType `json:"notificationType"`

	EventTime time.Time `json:"eventTime"`

	SystemDN string `json:"systemDN"`
}
