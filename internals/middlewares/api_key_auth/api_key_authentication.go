package apiAuth

import (
	"strings"
	"github.com/gin-gonic/gin"
	"github.com/swastiijain24/psp/internals/services"
)


type APIMiddleware struct {
	generator     *APIKeyGenerator
	hasher        *APIKeyHasher
	apiKeyService services.ApiKeyService
}

func NewApiAuthMiddleware(generator *APIKeyGenerator, hasher *APIKeyHasher, apiKeyService services.ApiKeyService) *APIMiddleware {
	return &APIMiddleware{
		generator:     generator,
		hasher:        hasher,
		apiKeyService: apiKeyService,
	}
}

func (m *APIMiddleware) ApiAuthentication() gin.HandlerFunc {
	return func(c *gin.Context) {

		pspID := c.GetHeader("X-PSP-ID")
		apiKey := m.extractAPIKey(c)
		if apiKey == "" {
			c.AbortWithStatusJSON(401, gin.H{"error": "API Key required"})
			return

		}

		if !m.generator.ValidateFormat(apiKey) {
			c.AbortWithStatusJSON(401, gin.H{"error": "Invalid API key format"})
			return
		}

		_, _, randomPart, err := m.generator.ParseKey(apiKey)
		if err != nil || len(randomPart) < 8 {
			c.AbortWithStatusJSON(401, gin.H{"error": "Invalid API key"})
			return
		}

		pspRegistration, err := m.apiKeyService.GetAPIKeyByPspId(c, pspID)
		if err != nil {
			c.AbortWithStatusJSON(401, gin.H{"error": "Invalid API key"})
			return
		}

		valid, err := m.hasher.Verify(apiKey, pspRegistration.HashedApiKey)
		if err != nil || !valid {
			c.AbortWithStatusJSON(401, gin.H{"error": "Invalid API key"})
			return
		}

		IsActive, err := m.apiKeyService.IsValid(c, pspID)
		if err != nil {
			c.AbortWithStatusJSON(500, gin.H{"error": "Internal server error"})
			return
		}

		if !IsActive {
			c.AbortWithStatusJSON(401, gin.H{"error": "API key is no longer valid"})
			return

		}


		c.Next()
	}
}

func (m *APIMiddleware) extractAPIKey(c *gin.Context) string {
	auth := c.GetHeader("Authorization")
	if strings.HasPrefix(auth, "Bearer ") {
		return strings.TrimPrefix(auth, "Bearer ")
	}

	if key := c.GetHeader("X-API-Key"); key != "" {
		return key
	}
	return ""
}
