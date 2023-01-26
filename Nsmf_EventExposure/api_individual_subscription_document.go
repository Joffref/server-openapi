/*
 * Nsmf_EventExposure
 *
 * Session Management Event Exposure Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.2.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_Nsmf_EventExposure

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// DeleteIndividualSubcription - Delete an individual subscription for event notifications from the SMF
func DeleteIndividualSubcription(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}

// GetIndividualSubcription - Read an individual subscription for event notifications from the SMF
func GetIndividualSubcription(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}

// ReplaceIndividualSubcription - Replace an individual subscription for event notifications from the SMF
func ReplaceIndividualSubcription(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}
