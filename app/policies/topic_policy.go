// Package policies 用户授权
package policies

import (
	"gohub/app/models/topic"
	"gohub/pkg/auth"

	"github.com/gin-gonic/gin"
)

func CanModifyTopic(c *gin.Context, topicModel topic.Topic) bool {
	return auth.CurrentUID(c) == topicModel.UserID
}
