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

// QueryLcsPrivacyData - Retrieves the LCS Privacy subscription data of a UE
func QueryLcsPrivacyData(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}
