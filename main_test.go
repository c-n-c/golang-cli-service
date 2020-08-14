package main

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"github.com/foolin/goview/supports/ginview"
	"github.com/gin-gonic/gin"
	"net/http/httptest"
	"testing"
)

func TestHomePage(t *testing.T) {


	t.Run("Homepage rendering test - 200 ok", func(t *testing.T){
		w := httptest.NewRecorder()
		r := gin.Default()
		r.HTMLRender = ginview.Default()
		r.GET("/", HomePage)

		ctx, _ := gin.CreateTestContext(w)

		req, _ := http.NewRequestWithContext(ctx, "GET", "/", nil)

		r.ServeHTTP(w, req)
		assert.Equal(t, 200, w.Code)
	})

}

func TestRunCmd(t *testing.T) {

	w := httptest.NewRecorder()
	r := gin.Default()
	r.HTMLRender = ginview.Default()
	r.GET("/runcmd", RunCmd)

	ctx, _ := gin.CreateTestContext(w)

	req, _ := http.NewRequestWithContext(ctx, "GET", "/runcmd", nil)

	r.ServeHTTP(w, req)
	assert.Equal(t, 500, w.Code)


	w1 := httptest.NewRecorder()
	ctx1, _ := gin.CreateTestContext(w1)
	req2, _ := http.NewRequestWithContext(ctx1, "GET", "/runcmd?text=uname", nil )

	r.ServeHTTP(w1, req2)
	assert.Equal(t, 200, w1.Code)
}