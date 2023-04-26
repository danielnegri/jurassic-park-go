// Copyright 2023 The Jurassic Park Authors
//
// Licensed under the AGPL, Version 3.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.gnu.org/licenses/agpl-3.0.en.html
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package server

import (
	"fmt"
	"net/http"
	"regexp"
	"runtime"
	"strings"
	"time"

	"github.com/danielnegri/jurassic-park-go/pkg/app"
	"github.com/danielnegri/jurassic-park-go/pkg/errors"
	"github.com/danielnegri/jurassic-park-go/pkg/version"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

const Prefix = "/api/v1"

var root = gin.H{
	"service":         app.Description,
	"arch":            runtime.GOARCH,
	"build_time":      version.BuildTime,
	"commit":          version.CommitHash,
	"os":              runtime.GOOS,
	"runtime_version": runtime.Version(),
	"version":         version.Version,
}

func (s *service) newHandler() http.Handler {
	if s.cfg.ReleaseMode {
		gin.SetMode(gin.ReleaseMode)
	}

	router := gin.New()
	router.Use(gin.Recovery())
	router.Use(cors.Default())
	router.Use(LoggerMiddleware(s.logger, s.now, time.RFC3339, true))
	router.NoRoute(s.handleNotFound)

	router.GET("/", s.handleRoot)
	router.GET("/health", s.handleHealth)
	router.GET("/ping", s.handlePing)

	return router
}

func (s *service) handleRoot(c *gin.Context) {
	c.JSON(http.StatusOK, root)
}

func (s *service) handleNotFound(c *gin.Context) {
	s.abortWithStatus(c, http.StatusNotFound, http.StatusText(http.StatusNotFound))
}

func (s *service) handleHealth(c *gin.Context) {
	if s.health.IsHealthy() {
		s.abortWithStatus(c, http.StatusInternalServerError, "Health check failed.")
		return
	}

	s.logger.Infof("Health check passed")
}

func (s *service) handlePing(c *gin.Context) {
	c.Status(http.StatusOK)
}

type HTTPErrorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func (er *HTTPErrorResponse) Error() string {
	return fmt.Sprintf("%d - %s", er.Code, er.Message)
}

func (s *service) abortWithStatus(c *gin.Context, code int, message string) {
	c.AbortWithStatusJSON(code, &HTTPErrorResponse{
		Code:    code,
		Message: message,
	})
}

var newLine = regexp.MustCompile(`\r?\n?\t`)

func (s *service) abortWithError(ctx *gin.Context, err error) {
	code := errors.KindUnexpected
	msg := newLine.ReplaceAllString(err.Error(), " ")
	e, ok := err.(*errors.Error)
	if ok {
		code = e.Kind
		if index := strings.Index(msg, ":"); len(msg) > index+1 {
			msg = strings.TrimSpace(msg[index+1:])
		}
	}

	ctx.AbortWithStatusJSON(code, &HTTPErrorResponse{
		Code:    code,
		Message: msg,
	})
}
