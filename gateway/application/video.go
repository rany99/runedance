package application

import (
	"context"
	"github.com/pkg/errors"
	"mime/multipart"
	"runedance/common/config"
	"runedance/gateway/rpc"
	"runedance/kitexGen/kitex_gen/userproto"
	"runedance/kitexGen/kitex_gen/videoproto"
	"runedance/models/pojo"
	"runedance/pkg/minio"
)

var VideoAppIns *VideoAppService

type VideoAppService struct {
}

func NewVideoAppService() *VideoAppService {
	return &VideoAppService{}
}

func (v VideoAppService) PublishVideo(ctx context.Context, appUserID int64, title string, fileHeader *multipart.FileHeader) (err error) {
	file, err := fileHeader.Open()
	if err != nil {
		return err
	}
	defer file.Close()
	cosUploadReq := &minio.Video{
		Title:    title,
		Filename: "/douyin/" + fileHeader.Filename,
		File:     file,
		UserID:   appUserID,
	}
	videoUrl, err := minio.UploadVideo(ctx, cosUploadReq)
	if err != nil {
		return err
	}
	req := &videoproto.CreateVideoReq{
		VideoBaseInfo: &videoproto.VideoBaseInfo{
			UserId:   appUserID,
			PlayUrl:  videoUrl.PlayUrl,
			CoverUrl: videoUrl.CoverUrl,
			Title:    title,
		},
	}
	if err := rpc.CreateVideo(ctx, req); err != nil {
		return errors.Wrapf(err, "CreateVideo rpc failed, appUserID: %v, title: %v", appUserID, title)
	}
	return nil
}

func (v VideoAppService) LikeVideo(ctx context.Context, appUserID int64, videoID int64) (err error) {
	req := &videoproto.LikeVideoReq{
		UserId:  appUserID,
		VideoId: videoID,
	}
	if err := rpc.LikeVideo(ctx, req); err != nil {
		return errors.Wrapf(err, "LikeVideo rpc failed, appUserID: %v, videoID: %v", appUserID, videoID)
	}
	return nil
}

func (v VideoAppService) UnLikeVideo(ctx context.Context, appUserID int64, videoID int64) (err error) {
	req := &videoproto.UnLikeVideoReq{
		UserId:  appUserID,
		VideoId: videoID,
	}
	if err := rpc.UnLikeVideo(ctx, req); err != nil {
		return errors.Wrapf(err, "UnLikeVideo rpc failed, appUserID: %v, videoID: %v", appUserID, videoID)
	}
	return nil
}

func (v VideoAppService) GetVideoList(ctx context.Context, appUserID int64, userId int64) (videoList []*pojo.Video, err error) {
	req := &videoproto.GetVideoListByUserIdReq{
		AppUserId: appUserID,
		UserId:    userId,
	}
	videos, err := rpc.GetVideoListByUserId(ctx, req)
	if err != nil {
		return nil, errors.Wrapf(err, "GetVideoListByUserId rpc failed, appUserID: %v, userId: %v", appUserID, userId)
	}
	// get authors
	n := len(videos)
	authors := make([]*userproto.UserInfo, n)
	for i := 0; i < n; i++ {
		subReq := &userproto.GetUserReq{
			AppUserId: appUserID,
			UserId:    videos[i].VideoBaseInfo.UserId,
		}
		authors[i], err = rpc.GetUser(ctx, subReq)
		if err != nil {
			return nil, errors.Wrapf(err, "GetUser rpc failed, appUserID: %v, userId: %v", appUserID, videos[i].VideoBaseInfo.UserId)
		}
	}
	// pack videos and authors
	packedVideos := toVideoDTOs(videos, authors)
	return packedVideos, nil
}

func (v VideoAppService) GetLikeVideoList(ctx context.Context, appUserID int64, userId int64) (userList []*pojo.Video, err error) {
	req := &videoproto.GetLikeVideoListReq{
		AppUserId: appUserID,
		UserId:    userId,
	}
	videos, err := rpc.GetLikeVideoList(ctx, req)
	if err != nil {
		return nil, errors.Wrapf(err, "GetLikeVideoList rpc failed, appUserID: %v, userId: %v", appUserID, userId)
	}
	n := len(videos)
	authors := make([]*userproto.UserInfo, n)
	for i := 0; i < n; i++ {
		subReq := &userproto.GetUserReq{
			AppUserId: appUserID,
			UserId:    videos[i].VideoBaseInfo.UserId,
		}
		authors[i], err = rpc.GetUser(ctx, subReq)
		if err != nil {
			return nil, errors.Wrapf(err, "GetUser rpc failed, appUserID: %v, userId: %v", appUserID, videos[i].VideoBaseInfo.UserId)
		}
	}
	packedVideos := toVideoDTOs(videos, authors)
	return packedVideos, nil
}

func (v VideoAppService) Feed(ctx context.Context, appUserID int64, latestTime int64) (videoList []*pojo.Video, nextTime int64, err error) {
	req := &videoproto.GetVideoListByTimeReq{
		AppUserId:  appUserID,
		LatestTime: latestTime,
		Count:      config.Server.FeedCount,
	}
	videos, nextTime, err := rpc.GetVideoListByTime(ctx, req)
	if err != nil {
		return nil, 0, errors.Wrapf(err, "GetVideoListByTime rpc failed, appUserID: %v, latestTime: %v", appUserID, latestTime)
	}
	n := len(videos)
	authors := make([]*userproto.UserInfo, n)
	for i := 0; i < n; i++ {
		subReq := &userproto.GetUserReq{
			AppUserId: appUserID,
			UserId:    videos[i].VideoBaseInfo.UserId,
		}
		authors[i], err = rpc.GetUser(ctx, subReq)
		if err != nil {
			return nil, 0, errors.Wrapf(err, "GetUser rpc failed, appUserID: %v, userId: %v", appUserID, videos[i].VideoBaseInfo.UserId)
		}
	}
	packedVideos := toVideoDTOs(videos, authors)
	return packedVideos, nextTime, nil
}

// toVideoDTO
// transform one videoproto.VideoInfo into one pojo.Video with author information
func toVideoDTO(v *videoproto.VideoInfo, author *userproto.UserInfo) *pojo.Video {
	if v == nil {
		return nil
	}
	return &pojo.Video{
		ID:           v.VideoId,
		Author:       toUserDTO(author),
		PlayAddr:     v.VideoBaseInfo.PlayUrl,
		CoverAddr:    v.VideoBaseInfo.CoverUrl,
		LikeCount:    v.LikeCount,
		CommentCount: v.CommentCount,
		IsFavorite:   v.IsFavorite,
		Title:        v.VideoBaseInfo.Title,
	}
}

// toVideoDTOs
// apply toVideoDTO to an array of videoproto.VideoInfo
func toVideoDTOs(vs []*videoproto.VideoInfo, authors []*userproto.UserInfo) []*pojo.Video {
	videos := make([]*pojo.Video, len(vs))
	for i, v := range vs {
		videos[i] = toVideoDTO(v, authors[i])
	}
	return videos
}
