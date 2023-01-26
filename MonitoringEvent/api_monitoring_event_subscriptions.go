/*
 * 3gpp-monitoring-event
 *
 * API for Monitoring Event.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.2.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_MonitoringEvent

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// CreateMonitoringEventSubscription - Creates a new subscription resource for monitoring event notification.
func CreateMonitoringEventSubscription(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}

// FetchAllMonitoringEventSubscriptions - Read all or queried active subscriptions for the SCS/AS.
func FetchAllMonitoringEventSubscriptions(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}