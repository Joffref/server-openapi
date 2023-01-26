/*
 * Npcf_PolicyAuthorization Service API
 *
 * PCF Policy Authorization Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.2.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_Npcf_PolicyAuthorization

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Route is the information for every URI.
type Route struct {
	// Name is the name of this Route.
	Name        string
	// Method is the string for the HTTP method. ex) GET, POST etc..
	Method      string
	// Pattern is the pattern of the URI.
	Pattern     string
	// HandlerFunc is the handler function of this route.
	HandlerFunc gin.HandlerFunc
}

// Routes is the list of the generated Route.
type Routes []Route

// NewRouter returns a new router.
func NewRouter() *gin.Engine {
	router := gin.Default()
	for _, route := range routes {
		switch route.Method {
		case http.MethodGet:
			router.GET(route.Pattern, route.HandlerFunc)
		case http.MethodPost:
			router.POST(route.Pattern, route.HandlerFunc)
		case http.MethodPut:
			router.PUT(route.Pattern, route.HandlerFunc)
		case http.MethodPatch:
			router.PATCH(route.Pattern, route.HandlerFunc)
		case http.MethodDelete:
			router.DELETE(route.Pattern, route.HandlerFunc)
		}
	}

	return router
}

// Index is the index handler.
func Index(c *gin.Context) {
	c.String(http.StatusOK, "Hello World!")
}

var routes = Routes{
	{
		"Index",
		http.MethodGet,
		"/npcf-policyauthorization/v1/",
		Index,
	},

	{
		"PostAppSessions",
		http.MethodPost,
		"/npcf-policyauthorization/v1/app-sessions",
		PostAppSessions,
	},

	{
		"DeleteEventsSubsc",
		http.MethodDelete,
		"/npcf-policyauthorization/v1/app-sessions/:appSessionId/events-subscription",
		DeleteEventsSubsc,
	},

	{
		"UpdateEventsSubsc",
		http.MethodPut,
		"/npcf-policyauthorization/v1/app-sessions/:appSessionId/events-subscription",
		UpdateEventsSubsc,
	},

	{
		"DeleteAppSession",
		http.MethodPost,
		"/npcf-policyauthorization/v1/app-sessions/:appSessionId/delete",
		DeleteAppSession,
	},

	{
		"GetAppSession",
		http.MethodGet,
		"/npcf-policyauthorization/v1/app-sessions/:appSessionId",
		GetAppSession,
	},

	{
		"ModAppSession",
		http.MethodPatch,
		"/npcf-policyauthorization/v1/app-sessions/:appSessionId",
		ModAppSession,
	},

	{
		"PcscfRestoration",
		http.MethodPost,
		"/npcf-policyauthorization/v1/app-sessions/pcscf-restoration",
		PcscfRestoration,
	},
}
