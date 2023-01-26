/*
 * Unified Data Repository Service API file for structured data for exposure
 *
 * The API version is defined in 3GPP TS 29.504   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: -
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_Exposure_Data

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// CreateIndividualExposureDataSubscription - Create a subscription to receive notification of exposure data changes
func CreateIndividualExposureDataSubscription(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}
