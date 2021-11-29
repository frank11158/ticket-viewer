package v1_test

import (
	"server/constant"
	"server/routes"

	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"strconv"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
	"gopkg.in/h2non/gock.v1"
)

var router *gin.Engine

func TestMain(m *testing.M) {
	gin.SetMode(gin.TestMode)
	constant.ReadConfig("../../.env")
	router = routes.InitRouter()

	exitCode := m.Run()

	os.Exit(exitCode)
}

func TestGetTickets(t *testing.T) {
	defer gock.Off()

	tests := []struct {
		testcase string
		perPage  int
		page     int
		status   int
		msg      string
	}{
		{
			testcase: "Ok",
			status:   http.StatusOK,
			msg:      "Ok",
			page:     1,
			perPage:  25,
		},
		{
			testcase: "API Error",
			status:   http.StatusOK,
			msg:      "Zendesk API failed",
			page:     1,
			perPage:  25,
		},
		{
			testcase: "internal server error",
			status:   http.StatusOK,
			msg:      "Fail",
			page:     1,
			perPage:  25,
		},
	}

	for _, test := range tests {
		if test.testcase != "Ok" {
			mockTicketsError(test.testcase)
		} else {
			mockTickets(strconv.Itoa(test.perPage), strconv.Itoa(test.page))
		}

		q := url.Values{}
		q.Add("per_page", strconv.Itoa(test.perPage))
		q.Add("page", strconv.Itoa(test.page))
		w, m := request("GET", "/api/v1/tickets", nil, q.Encode())

		assert.Equal(t, test.status, w.Code)
		assert.Equal(t, test.msg, m["msg"])
	}
}

func TestGetUserInfo(t *testing.T) {
	defer gock.Off()

	tests := []struct {
		testcase string
		userID   string
		status   int
		msg      string
	}{
		{
			testcase: "OK",
			userID:   "0123456789",
			status:   http.StatusOK,
			msg:      "Ok",
		},
		{
			testcase: "invalid userID",
			userID:   "invalid-userID",
			status:   http.StatusOK,
			msg:      "Zendesk API failed",
		},
		{
			testcase: "internal server error",
			userID:   "test",
			status:   http.StatusOK,
			msg:      "Fail",
		},
	}

	for _, test := range tests {
		if test.testcase == "internal server error" {
			mockUserInfoError()
		}
		mockUserInfo(test.userID)

		w, m := request("GET", fmt.Sprintf("/api/v1/users/%s", test.userID), nil, "")

		assert.Equal(t, test.status, w.Code)
		assert.Equal(t, test.msg, m["msg"])
	}
}

func mockTickets(perPage string, page string) {
	gock.New(fmt.Sprintf("https://%s", viper.GetString("ZENDESK_DOMAIN"))).
		Get("/api/v2/tickets").
		MatchParam("per_page", perPage).
		MatchParam("page", page).
		Reply(200).
		JSON(map[string]interface{}{
			"code": 200,
			"tickets": []map[string]interface{}{
				{
					"id": 1,
				},
			},
		})
}

func mockTicketsError(errtype string) {
	if errtype == "API Error" {
		gock.New(fmt.Sprintf("https://%s", viper.GetString("ZENDESK_DOMAIN"))).
			Get("/api/v2/tickets").
			Reply(401).
			JSON(map[string]interface{}{
				"error": "Couldn't authenticate you",
			})
	} else if errtype == "internal srever error" {
		gock.New(fmt.Sprintf("https://%s", viper.GetString("ZENDESK_DOMAIN"))).
			Get(fmt.Sprintf("/api/v2/users/%s", "test")).
			Reply(200)
	}
}

func mockUserInfo(userID string) {
	if userID != "invalid-userID" {
		gock.New(fmt.Sprintf("https://%s", viper.GetString("ZENDESK_DOMAIN"))).
			Get(fmt.Sprintf("/api/v2/users/%s", userID)).
			Reply(200).
			JSON(map[string]interface{}{
				"code": 200,
			})
	} else {
		gock.New(fmt.Sprintf("https://%s", viper.GetString("ZENDESK_DOMAIN"))).
			Get(fmt.Sprintf("/api/v2/users/%s", userID)).
			Reply(404).
			JSON(map[string]interface{}{
				"error":       "RecordNotFound",
				"description": "Not found",
			})
	}
}

func mockUserInfoError() {
	gock.New(fmt.Sprintf("https://%s", viper.GetString("ZENDESK_DOMAIN"))).
		Get(fmt.Sprintf("/api/v2/users/%s", "test")).
		Reply(200)
}

func request(method string, path string, body []byte, query string) (*httptest.ResponseRecorder, map[string]interface{}) {
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(method, path, bytes.NewBuffer(body))
	req.URL.RawQuery = query
	router.ServeHTTP(w, req)

	m := make(map[string]interface{})
	_ = json.Unmarshal(w.Body.Bytes(), &m)
	return w, m
}
