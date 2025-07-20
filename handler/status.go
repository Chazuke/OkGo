package handler

import (
	"net/http"
	"runtime/debug"
	"time"

	"github.com/gofiber/fiber/v2"
)

var info *debug.BuildInfo

func init() {
	info, _ = debug.ReadBuildInfo()
}

type StatusResponse struct {
	Version    string           `json:"version,omitempty"`
	CompiledAt string           `json:"compiled_at,omitempty"`
	ExecutedAt string           `json:"executed_at"`
	Uptime     string           `json:"uptime"`
	BuildInfo  *debug.BuildInfo `json:"build_info"`
}

func (h *handler) getStatus() fiber.Handler {
	return func(c *fiber.Ctx) error {
		return c.Status(http.StatusOK).JSON(StatusResponse{
			Version:    h.info.Version,
			CompiledAt: h.info.CompiledAt,
			ExecutedAt: h.info.ExecutedAt.Format(time.RFC3339),
			Uptime:     time.Now().UTC().Sub(h.info.ExecutedAt).String(),
			BuildInfo:  info,
		})
	}
}
