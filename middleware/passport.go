package middleware

import (
	"errors"
	"example.com/m/internal/models"
	"example.com/m/internal/repository"
	"example.com/m/tools"
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"time"
)

// IdentityKeyID is used to tell
// by what field we will identify user
const IdentityKeyID = "id"

// UserID struct
type UserID struct {
	ID string
}

var userRepo = repository.NewUserRepo()

// Passport is middleware for user authentication
func Passport() *jwt.GinJWTMiddleware {
	authMiddleware, _ := jwt.New(&jwt.GinJWTMiddleware{
		Realm:          "TastyOffice",
		Key:            []byte(os.Getenv("JWTSECRET")),
		Timeout:        time.Hour * 4,
		MaxRefresh:     time.Hour * 24,
		IdentityKey:    IdentityKeyID,
		SendCookie:     true,
		CookieMaxAge:   time.Hour * 24,
		CookieHTTPOnly: true,
		TokenHeadName:  "jwt",
		TokenLookup:    "cookie: jwt",
		LoginResponse: func(c *gin.Context, i int, s string, t time.Time) {
			value, _ := Passport().ParseTokenString(s)
			id := jwt.ExtractClaimsFromToken(value)["id"]
			result, err := userRepo.GetByID(id.(string))

			if err != nil {
				tools.CreateError(http.StatusUnauthorized, err, c)
				return
			}

			c.JSON(http.StatusOK, result)
		},
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if v, ok := data.(*UserID); ok {
				return jwt.MapClaims{
					IdentityKeyID: v.ID,
				}
			}
			return jwt.MapClaims{}
		},
		IdentityHandler: func(c *gin.Context) interface{} {
			claims := jwt.ExtractClaims(c)
			return &UserID{
				ID: claims[IdentityKeyID].(string),
			}
		},
		Authenticator: func(c *gin.Context) (interface{}, error) {
			var body models.LoginUserModel
			if err := c.ShouldBind(&body); err != nil {
				return "", errors.New("пароль или логин не введен")
			}

			result, err := userRepo.GetByKey("email", body.Username)
			if err != nil {
				return nil, errors.New("не верный логин или пароль")
			}
			equal := tools.CheckPasswordHash(body.Password, result.Password)
			if equal {
				return &UserID{
					ID: result.ID.String(),
				}, nil
			}
			return nil, errors.New("не верный логин или пароль")
		},
		Unauthorized: func(c *gin.Context, code int, message string) {
			c.JSON(code, gin.H{
				"code":    code,
				"message": message,
			})
		},
	})
	return authMiddleware
}
