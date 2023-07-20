package authorization

import (
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"runedance/models/respond"
	"time"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/cloudwego/kitex/pkg/klog"

	"runedance/common/config"
	"runedance/common/constant"
	"runedance/common/errorCode"
)

var mw *jwt.GinJWTMiddleware

func Init() {
	var err error
	mw, err = jwt.New(&jwt.GinJWTMiddleware{
		Key:        []byte(config.JWT.Secret),
		Timeout:    config.JWT.Expires,
		MaxRefresh: time.Hour,
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if v, ok := data.(int64); ok {
				return jwt.MapClaims{
					constant.IdentityKey: v,
				}
			}
			return jwt.MapClaims{}
		},
		TokenLookup:   "query: token, form: token, header: Authorization, cookie: jwt",
		TokenHeadName: "Bearer",
		TimeFunc:      time.Now,
	})
	if err != nil {
		klog.Fatal(err)
	}
}

func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		userId, err := GetUserID(c)
		if err != nil {
			respond.Error(c, err)
			c.Abort()
			return
		}
		c.Set(constant.IdentityKey, userId)
		c.Next()
	}
}

func GetUserID(c *gin.Context) (int64, error) {
	claims, err := mw.GetClaimsFromJWT(c)
	if err != nil {
		return 0, errors.Wrap(errorCode.UnauthorizedErr, err.Error())
	}
	tempUserID, ok := claims[constant.IdentityKey].(float64)
	userId := int64(tempUserID)
	if !ok || userId <= 0 {
		return 0, errors.Wrap(errorCode.UnauthorizedErr, "jwt user_id error")
	}
	return userId, nil
}

func GenerateToken(userId int64) (string, error) {
	token, _, err := mw.TokenGenerator(userId)
	return token, err
}
