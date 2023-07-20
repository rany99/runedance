package pack

import (
	"runedance/kitexGen/kitex_gen/videoproto"
	"runedance/video/dao/dal/entity"
)

// Video pack video info : video to videoproto.VideoInfo
func Video(m *entity.Video) *videoproto.VideoInfo {
	if m == nil {
		return nil
	}
	return &videoproto.VideoInfo{
		VideoBaseInfo: &videoproto.VideoBaseInfo{
			UserId:   int64(m.UserId),
			PlayUrl:  m.PlayUrl,
			CoverUrl: m.CoverUrl,
			Title:    m.Title,
		},
		VideoId:      int64(m.ID),
		LikeCount:    m.FavoriteCount,
		CommentCount: m.CommentCount,
	}
}

func Videos(ms []*entity.Video) []*videoproto.VideoInfo {
	videos := make([]*videoproto.VideoInfo, len(ms))
	for i, m := range ms {
		if n := Video(m); n != nil {
			videos[i] = n
		}
	}
	return videos
}
