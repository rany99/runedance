package api

import (
	"github.com/gin-gonic/gin"
	"runedance/common/constant"
	"runedance/gateway/application"
	"runedance/models/pojo"
	"runedance/models/respond"
	"runedance/pkg/errorCode"
)

func CommentAction(c *gin.Context) {
	appUserID := c.GetInt64(constant.IdentityKey)
	param := new(pojo.CommentOperationReq)
	if err := c.ShouldBind(param); err != nil {
		respond.Error(c, err)
		return
	}
	switch param.ActionType {
	case constant.CreateComment: // create comment on a video
		comment, err := application.CommentAppIns.CreateComment(c, appUserID, param.VideoId, param.CommentText)
		if err != nil {
			respond.Error(c, err)
			return
		}
		resp := &pojo.CreateCommentResp{
			BaseResp: respond.Success,
			Comment:  comment,
		}
		respond.Send(c, resp)
	case constant.DeleteComment: // delete one comment
		if err := application.CommentAppIns.DeleteComment(c, param.CommentId, param.VideoId); err != nil {
			respond.Error(c, err)
			return
		}
		respond.OK(c)
	default:
		respond.Error(c, errorCode.ParamErr)
	}
}

// CommentList (GET): get comment list of one video
func CommentList(c *gin.Context) {
	appUserID := c.GetInt64(constant.IdentityKey)
	param := new(pojo.CommentListReq)
	if err := c.ShouldBind(param); err != nil {
		respond.Error(c, err)
		return
	}
	comments, err := application.CommentAppIns.GetCommentList(c, appUserID, param.VideoId)
	if err != nil {
		respond.Error(c, err)
		return
	}
	resp := &pojo.CommentListResp{
		BaseResp:    respond.Success,
		CommentList: comments,
	}
	respond.Send(c, resp)
}
