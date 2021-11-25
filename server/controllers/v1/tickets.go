package v1

import (
	"server/controllers/util"

	"encoding/json"
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

	rawData, err := util.Request("GET", zendeskURL.String(), nil)
	if err != nil {
		log.Printf("Error occures fetching tickets from zendesk API, error code: %s", err)
		constant.ResponseWithData(c, http.StatusOK, constant.ERROR, nil)
	}

	var data map[string]interface{}
	err = json.Unmarshal([]byte(rawData), &data)
	if err != nil {
		constant.ResponseWithData(c, http.StatusOK, constant.ERROR, nil)
	}

	constant.ResponseWithData(c, http.StatusOK, constant.SUCCESS, data)
}
