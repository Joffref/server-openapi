/*
 * SS_NetworkResourceMonitoring
 *
 * API for SEAL Network Resource Monitoring.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package SS_NetworkResourceMonitoring

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// ReadUnicastMonitoringSubscription - Read an existing individual unicast monitoring subscription resource according to the subscriptionId.
func ReadUnicastMonitoringSubscription(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}

// UnsubscribeUnicastMonitoring - Remove an existing individual unicast monitoring subscription resource according to the subscriptionId.
func UnsubscribeUnicastMonitoring(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}