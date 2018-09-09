package fb

import (
	"encoding/json"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	fclient "github.com/madebyais/facebook-go-sdk"
	"github.com/vtomkiv/golang.api/util"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/facebook"
	"net/http"
	"time"
)

type FBUser struct {
	FirstName string `json:"first_name"`
}

type JwtFBClaims struct {
	Name string `json:"user_name"`
	jwt.StandardClaims
}

var FACEBOOK = &oauth2.Config{
	ClientID:     "1770367409743952",
	ClientSecret: "01400a8f25ead4d791275fb14b98535c",
	Scopes:       []string{},
	Endpoint:     facebook.Endpoint,
	RedirectURL:  "http://localhost:8088/auth/fb/callback",
}


// random string for oauth2 API calls to protect against CSRF
var	oauthStateString = "blablabla"

var logger = *util.GetLoggerInstance()


func Login(c echo.Context) error {
	url := FACEBOOK.AuthCodeURL(oauthStateString, oauth2.AccessTypeOffline)
	return c.Redirect(http.StatusTemporaryRedirect, url)
}

func HandleFBCallback(c echo.Context) error {
	state := c.FormValue("state")
	if state != oauthStateString {
		logger.Warnf("invalid oauth state, expected '%s', got '%s'", oauthStateString, state)
		return echo.ErrUnauthorized
		}

	code := c.FormValue("code")
	token, err := FACEBOOK.Exchange(oauth2.NoContext, code)
	if err != nil {
		logger.Warnf("oauthConf.Exchange() failed with '%v'", err)
		return echo.ErrUnauthorized
	}

	data , err := fclient.New().SetAccessToken(token.AccessToken).API(`/me`).Fields("first_name").Get()

	if err != nil {
		logger.Warnf("facebook.get failed with '%v'", err)
		return echo.ErrUnauthorized
	}

	var user FBUser

	json.Unmarshal([]byte(data.(string)), &user)

	// Set custom claims
	claims := JwtFBClaims{
		user.FirstName,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
		},
	}

	// Create token with claims
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token and send it as response.
	t, err := jwtToken.SignedString([]byte("secret"))
	if err != nil {
		logger.Errorf("failed to generate jwt token, error: %v ", err)
		return echo.ErrUnauthorized
	}

	return c.JSON(http.StatusOK, echo.Map{
		"token": t,
	})

	return echo.ErrUnauthorized

}

func Accessible(c echo.Context) error {
	return c.String(http.StatusOK, "Accessible")
}

func Restricted(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*JwtFBClaims)
	logger.Infof("user found with user name '%s'", claims.Name)
	name := claims.Name
	return c.String(http.StatusOK, "Welcome "+name+"!")
}
