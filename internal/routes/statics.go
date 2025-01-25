package routes

import (
	"net/http"
	"thor/internal/embeded"

	"github.com/gofiber/fiber/v2/middleware/filesystem"
)

var staticsHandler = filesystem.New(filesystem.Config{
	Root:       http.FS(embeded.Statics),
	PathPrefix: "statics",
	MaxAge:     60 * 60,
})
