/*
 * NRF NFManagement Service
 *
 * NRF NFManagement Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.3.0-alpha.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_Nnrf_NFManagement

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// CreateSubscription - Create a new subscription
func CreateSubscription(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}
