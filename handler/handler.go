package handler

import (
	"net/http"
	"okgo/logger"
	"regexp"
	"time"

	"github.com/gofiber/fiber/v2"
)

type (
	handler struct {
		log  *logger.OkGoLogger
		info *HandlerInfo
	}

	HandlerInfo struct {
		CompiledAt string
		ExecutedAt time.Time
		GitBranch  string
		GitCommit  string
		Version    string
	}
)

const (
	statusEndpointURI = "/status"
)

var skipRegex *regexp.Regexp

func init() {
	skipRegex = regexp.MustCompile("^(kube-probe|GoogleHC/)")
}

func New(info *HandlerInfo, log *logger.OkGoLogger) *handler {
	return &handler{
		log:  log,
		info: info,
	}
}

func (h *handler) AddStatusEndpoint(f *fiber.App, basePath string) {
	router := f.Group(basePath)

	router.Add(http.MethodGet, statusEndpointURI, h.getStatus())
}

// Skipper is used for speciying which route(s) should be opted out by the OTEL collector
func Skipper(c *fiber.Ctx) bool {
	if c.Path() == statusEndpointURI {
		return true
	}
	return skipRegex.MatchString(c.Get("User-Agent"))
}
