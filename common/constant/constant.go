package constant

// rpc service name
const (
	IdentityKey              = "id"
	VideoDomainServiceName   = "video_domain_service"
	UserDomainServiceName    = "user_domain_service"
	CommentDomainServiceName = "comment_domain_service"
	MessageDomainServiceName = "message_domain_service"
)

// redis key prefix
const (
	UserInfoRedisPrefix          = "user_info:"
	UserInfoCntRedisPrefix       = "user_info_cnt:"
	FollowCountRedisPrefix       = "follow_cnt:"
	FanCountRedisPrefix          = "fan_cnt:"
	WorkCountRedisPrefix         = "work_cnt:"
	FavoriteCountRedisPrefix     = "favorite_cnt:"
	FollowRedisPrefix            = "follow:"
	FanRedisPrefix               = "fan:"
	VideoInfoRedisPrefix         = "video_info:"
	LikeRedisPrefix              = "like:"
	PublishRedisPrefix           = "publish:"
	VideoInfoCntRedisPrefix      = "video_info_cnt:"
	CommentCountRedisPrefix      = "comment_cnt:"
	LikeCountRedisPrefix         = "like_cnt:"
	CommentRedisPrefix           = "comment:"
	CommentInfoRedisPrefix       = "comment_info:"
	MessageRedisPrefix           = "message:"
	MessageLatestTimeRedisPrefix = "message_latest_time:"
	LoginFailCounterRedisPrefix  = "user_login_fail_counter:"
	LoginFailLockRedisPrefix     = "user_login_fail_lock:"
)

// pulsar topic name
const (
	LikeVideoTopic     = "persistent://public/douyin_prod/like_video"
	FollowUserTopic    = "persistent://public/douyin_prod/follow_user"
	CreateMessageTopic = "persistent://public/douyin_prod/create_message"
	CommentTopic       = "persistent://public/douyin_prod/comment"
)

// Action type
const (
	LikeVideo   = 1
	UnLikeVideo = 2
)

const (
	FollowUser   = 1
	UnFollowUser = 2
)

const (
	CreateComment = 1
	DeleteComment = 2
)
