/*
 * Nucmf_Provisioning
 *
 * UCMF_Provisioning Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_Nucmf_Provisioning

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// CreateProvisioning - Create an Individual UE radio capability provisioning
func CreateProvisioning(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}