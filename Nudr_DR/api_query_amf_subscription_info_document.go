/*
 * Nudr_DataRepository API OpenAPI file
 *
 * Unified Data Repository Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 2.2.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_Nudr_DR

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetAmfGroupSubscriptions - Retrieve AMF subscription Info for a group of UEs or any UE
func GetAmfGroupSubscriptions(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}

// GetAmfSubscriptionInfo - Retrieve AMF subscription Info
func GetAmfSubscriptionInfo(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}