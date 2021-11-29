package v1

import (
	"server/controllers/util"

	"fmt"
	"log"
	"net/http"
	"net/url"
	"path"
	"server/constant"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

// @Summary Get Tickets
// @produce Json
// @Param page query string false "# of Page"
// @Param per_page query string false "# of Tickets per Page"
// @Success 200 {object} constant.Response
// @Failure 500 {object} constant.Response
// @Router /tickets [get]
func GetTickets(c *gin.Context) {
	zendeskURL, err := url.Parse(fmt.Sprintf("https://%s", viper.GetString("ZENDESK_DOMAIN")))
	if err != nil {
		constant.ResponseWithData(c, http.StatusOK, constant.ERROR, nil)
		return
	}

	zendeskURL.Path = path.Join("/api/v2/tickets")
	q := zendeskURL.Query()
	q.Set("per_page", c.Query("per_page"))
	q.Set("page", c.Query("page"))
	zendeskURL.RawQuery = q.Encode()

	data, err := util.Request("GET", zendeskURL.String(), nil)
	if err != nil {
		if data != nil {
			log.Printf("Error occures when fetching tickets from zendesk API, error code: %s", err)
			fmt.Println(data["error"])
			constant.ResponseWithData(c, http.StatusOK, constant.ERROR_API, data)
			return
		}
		log.Printf("Error occures when fetching tickets from zendesk API: %s", err)
		constant.ResponseWithData(c, http.StatusOK, constant.ERROR, data)
		return
	}

	constant.ResponseWithData(c, http.StatusOK, constant.SUCCESS, data)
}

// @Summary Get User Info
// @produce Json
// @Success 200 {object} constant.Response
// @Failure 500 {object} constant.Response
// @Router /users/{user_id} [get]
func GetUserInfo(c *gin.Context) {
	var params struct {
		ID string `uri:"user_id" binding:"required"`
	}
	if err := c.ShouldBindUri(&params); err != nil {
		constant.ResponseWithData(c, http.StatusOK, constant.ERROR, err)
		return
	}

	zendeskURL, err := url.Parse(fmt.Sprintf("https://%s", viper.GetString("ZENDESK_DOMAIN")))
	if err != nil {
		constant.ResponseWithData(c, http.StatusOK, constant.ERROR, nil)
		return
	}

	zendeskURL.Path = path.Join(fmt.Sprintf("/api/v2/users/%s", params.ID))

	data, err := util.Request("GET", zendeskURL.String(), nil)
	if err != nil {
		if data != nil {
			log.Printf("Error occures fetching user info from zendesk API, error code: %s", err)
			fmt.Println(data["error"])
			constant.ResponseWithData(c, http.StatusOK, constant.ERROR_API, data)
			return
		}
		log.Printf("Error occures fetching user info from zendesk API: %s", err)
		constant.ResponseWithData(c, http.StatusOK, constant.ERROR, data)
		return
	}

	constant.ResponseWithData(c, http.StatusOK, constant.SUCCESS, data)
}
