package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary Get Tickets
// @produce Json
// @Param page query string false "# of Page"
// @Param per_page query string false "# of Tickets per Page"
// @Success 200 {object} constant.Response
// @Failure 500 {object} constant.Response
// @Router /tickets [get]
func GetTickets(c *gin.Context) {
	c.String(http.StatusOK, "Alive")
}
