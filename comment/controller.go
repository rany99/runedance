package comment

import (
	"context"
	"runedance/comment/pack"
	"runedance/comment/service"
	"runedance/kitexGen/kitex_gen/commentproto"
	"runedance/pkg/errorCode"
)

// CommentServiceImpl implements the last service interface defined in the IDL.
type CommentServiceImpl struct{}

// CreateComment implements the CommentServiceImpl interface.
func (s *CommentServiceImpl) CreateComment(ctx context.Context, req *commentproto.CreateCommentReq) (resp *commentproto.CreateCommentResp, err error) {
	resp = new(commentproto.CreateCommentResp)

	if req.UserId < 0 || req.VideoId < 0 || len(req.Content) == 0 { // Empty comments are not allowed
		resp.BaseResp = pack.BuildBaseResp(errorCode.ParamErr)
		return resp, nil
	}

	commentInfo, err := service.NewCreateCommentService(ctx).CreateComment(req.UserId, req.VideoId, req.Content)
	if err != nil {
		resp.BaseResp = pack.BuildBaseResp(err)
		return resp, nil
	}
	resp.BaseResp = pack.BuildBaseResp(errorCode.Success)
	resp.CommentInfo = commentInfo
	return resp, nil
}

// DeleteComment implements the CommentServiceImpl interface.
func (s *CommentServiceImpl) DeleteComment(ctx context.Context, req *commentproto.DeleteCommentReq) (resp *commentproto.DeleteCommentResp, err error) {
	resp = new(commentproto.DeleteCommentResp)

	if req.CommentId < 0 || req.VideoId < 0 { // ensure the ID > 0
		resp.BaseResp = pack.BuildBaseResp(errorCode.ParamErr)
		return resp, nil
	}
	err = service.NewDeleteCommentService(ctx).DeleteComment(req.CommentId, req.VideoId)
	if err != nil {
		resp.BaseResp = pack.BuildBaseResp(err)
		return resp, nil
	}
	resp.BaseResp = pack.BuildBaseResp(errorCode.Success)
	return resp, nil
}

// GetCommentList implements the CommentServiceImpl interface.
func (s *CommentServiceImpl) GetCommentList(ctx context.Context, req *commentproto.GetCommentListReq) (resp *commentproto.GetCommentListResp, err error) {
	resp = new(commentproto.GetCommentListResp)

	if req.VideoId < 0 {
		resp.BaseResp = pack.BuildBaseResp(errorCode.ParamErr)
		return resp, nil
	}
	comments, err := service.NewGetCommentListService(ctx).GetCommentList(req.VideoId)
	if err != nil {
		resp.BaseResp = pack.BuildBaseResp(err)
		return resp, nil
	}
	resp.BaseResp = pack.BuildBaseResp(errorCode.Success)
	resp.CommentInfos = comments
	return resp, nil
}
