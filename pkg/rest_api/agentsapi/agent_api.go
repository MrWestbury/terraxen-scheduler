package agentsapi

import "github.com/gin-gonic/gin"

type AgentPoolApi struct{}

func NewAgentPoolApi(parentGroup *gin.RouterGroup) *AgentPoolApi {
	aga := &AgentPoolApi{}

	agentpools := parentGroup.Group("agentpools")

	agentpools.GET("/")  // List agentpools
	agentpools.POST("/") // Create agentpool

	agentpools.DELETE("/:agentpool") // Delete an agentpool
	agentpools.GET("/:agentpool")    // Get details of an agentpool
	agentpools.PUT("/:agentpool")    // Update agentpool

	agentpools.GET("/:agentpool/agents") // List agents

	return aga
}
