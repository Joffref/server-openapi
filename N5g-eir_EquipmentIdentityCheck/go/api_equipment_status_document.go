/*
 * 5G-EIR Equipment Identity Check
 *
 * 5G-EIR Equipment Identity Check Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.2.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package N5g-eir_EquipmentIdentityCheck

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetEquipmentStatus - Retrieves the status of the UE
func GetEquipmentStatus(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}