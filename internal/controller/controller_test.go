package controller

import (
	"OceanQA/boot"
	"OceanQA/internal/conf"
	"github.com/go-resty/resty/v2"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestShowAllQuestion(t *testing.T) {
	boot.Initialize()
	client := resty.New()
	resp, err := client.R().Get("http://localhost:" + conf.ServerConf.Port + "/v1/question/all")
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, http.StatusOK, resp.StatusCode())
	t.Log(resp.String())
}
