/*
 * Unified Data Repository Service API file for subscription data
 *
 * Unified Data Repository Service (subscription data).   The API version is defined in 3GPP TS 29.504.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: -
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_Subscription_Data

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// ModifyEesubscription - Modify an individual ee subscription of a UE
func ModifyEesubscription(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}

// QueryeeSubscription - Retrieve a eeSubscription
func QueryeeSubscription(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}

// RemoveeeSubscriptions - Deletes a eeSubscription
func RemoveeeSubscriptions(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}

// UpdateEesubscriptions - Update an individual ee subscriptions of a UE
func UpdateEesubscriptions(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}