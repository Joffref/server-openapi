/*
 * Unified Data Repository Service API file for subscription data
 *
 * Unified Data Repository Service (subscription data).   The API version is defined in 3GPP TS 29.504.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: -
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package Subscription_Data

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// RemoveAmfGroupSubscriptions - Deletes AMF Subscription Info for an eeSubscription for a group of UEs or any UE
func RemoveAmfGroupSubscriptions(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}

// RemoveAmfSubscriptionsInfo - Deletes AMF Subscription Info for an eeSubscription
func RemoveAmfSubscriptionsInfo(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}