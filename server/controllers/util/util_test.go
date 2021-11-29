package util_test

import (
	"fmt"
	"os"
	"server/constant"
	"server/controllers/util"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
	"gopkg.in/h2non/gock.v1"
)

func TestMain(m *testing.M) {
	gin.SetMode(gin.TestMode)
	constant.ReadConfig("../../.env")

	exitCode := m.Run()

	os.Exit(exitCode)
}

func TestRequest(t *testing.T) {
	defer gock.Off()

	zendeskURL := fmt.Sprintf("https://%s", viper.GetString("ZENDESK_DOMAIN"))

	tests := []struct {
		testcase string
		data     map[string]interface{}
		err      error
	}{
		{
			testcase: "Ok",
			data: map[string]interface{}{
				"test": "success",
			},
			err: nil,
		},
		{
			testcase: "auth failed",
			data: map[string]interface{}{
				"error": "Auth failed",
			},
			err: fmt.Errorf("%d", 401),
		},
		{
			testcase: "server error",
			data: map[string]interface{}{
				"error": "server error",
			},
			err: fmt.Errorf("%d", 500),
		},
		{
			testcase: "other error",
			data:     nil,
			err:      fmt.Errorf("Error"),
		},
	}

	for _, test := range tests {
		mockRequest(test.testcase)

		data, err := util.Request("GET", fmt.Sprintf("%s/api/v2/mock", zendeskURL), nil)

		assert.Equal(t, test.data, data)
		if test.testcase != "other error" {
			assert.Equal(t, test.err, err)
		} else {
			assert.NotNil(t, err)
		}
	}
}

func mockRequest(testcase string) {
	if testcase == "Ok" {
		gock.New(fmt.Sprintf("https://%s", viper.GetString("ZENDESK_DOMAIN"))).
			Get("/api/v2/mock").
			Reply(200).
			JSON(map[string]interface{}{
				"test": "success",
			})
	} else if testcase == "auth failed" {
		gock.New(fmt.Sprintf("https://%s", viper.GetString("ZENDESK_DOMAIN"))).
			Get("/api/v2/mock").
			Reply(401).
			JSON(map[string]interface{}{
				"error": "Auth failed",
			})
	} else if testcase == "server error" {
		gock.New(fmt.Sprintf("https://%s", viper.GetString("ZENDESK_DOMAIN"))).
			Get("/api/v2/mock").
			Reply(500).
			JSON(map[string]interface{}{
				"error": "server error",
			})
	} else if testcase == "other error" {
		gock.New(fmt.Sprintf("https://%s", viper.GetString("ZENDESK_DOMAIN"))).
			Get("/api/v2/mock").
			Reply(500)
	}
}
