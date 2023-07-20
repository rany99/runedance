package api

import (
	"github.com/gin-gonic/gin"
	auth "runedance/gateway/api/authorization"
	"runedance/gateway/application"
	"runedance/models/pojo"
	"runedance/models/respond"
	"runedance/pkg/errorCode"
)

func GetMessageList(c *gin.Context) {
	appUserID, err := auth.GetUserID(c)
	if err != nil {
		respond.Error(c, err)
		return
	}

	param := new(pojo.MessageListReq)
	if err := c.ShouldBind(param); err != nil {
		respond.Error(c, err)
		return
	}
	messageList, err := application.MessageAppIns.GetMessageList(c, appUserID, param.ToUserId)
	if err != nil {
		respond.Error(c, err)
	}

	response := pojo.MessageListResp{
		BaseResp:    respond.Success,
		MessageList: messageList,
	}
	respond.Send(c, &response)
}

// Handle the POST request of /message/action/, currently only support message sending
func HandleMessageAction(c *gin.Context) {
	appUserID, err := auth.GetUserID(c)
	if err != nil {
		respond.Error(c, err)
		return
	}

	param := new(pojo.MessageOperationReq)
	if err := c.ShouldBind(param); err != nil {
		respond.Error(c, err)
		return
	}

	switch param.ActionType {
	case 1:
		err := application.MessageAppIns.CreateMessage(c, appUserID, param.ToUserId, param.Content)
		if err != nil {
			respond.Error(c, err)
		} else {
			respond.OK(c)
		}
	default:
		respond.Error(c, errorCode.ParamErr)
	}
}
