/*
 * 3gpp-bdt
 *
 * API for BDT resouce management.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.2.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_ResourceManagementOfBdt

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// CreateBDTSubscription - Creates a new background data transfer subscription resource.
func CreateBDTSubscription(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}

// FetchAllActiveBDTSubscriptions - Fetch all active background data transfer subscription resources for a given SCS/AS.
func FetchAllActiveBDTSubscriptions(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}