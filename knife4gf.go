package knife4gf

import (
	"fmt"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcache"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/util/gconv"
	"net/http"
	"time"
)

// Knife4gf is the struct for swagger feature management.
type Knife4gf struct {
	Info          SwaggerInfo // Swagger information.
	Schemes       []string    // Supported schemes of the swagger API like "http", "https".
	Host          string      // The host of the swagger APi like "127.0.0.1", "www.mydomain.com"
	BasicPath     string      // The URI for the swagger API like "/", "v1", "v2".
	BasicAuthUser string      `c:"user"` // HTTP basic authentication username.
	BasicAuthPass string      `c:"pass"` // HTTP basic authentication password.
}

// SwaggerInfo is the information field for swagger.
type SwaggerInfo struct {
	Title          string // Title of the swagger API.
	Version        string // Version of the swagger API.
	TermsOfService string // As the attribute name.
	Description    string // Detail description of the swagger API.
}

const (
	Name               = "knife4gf"
	Author             = "sqiu_li@163.com"
	Version            = "v1.0.0"
	Description        = "knife4gf is knife4j for GoFrame GoFrame project. https://github.com/iasuma/knife4gf"
	MaxAuthAttempts    = 10          // Max authentication count for failure try.
	AuthFailedInterval = time.Minute // Authentication retry interval after last failed.
)

const (
	docPath = "/kdoc"
)

// Name returns the name of the plugin.
func (kf *Knife4gf) Name() string {
	return Name
}

// Author returns the author of the plugin.
func (kf *Knife4gf) Author() string {
	return Author
}

// Version returns the version of the plugin.
func (kf *Knife4gf) Version() string {
	return Version
}

// Description returns the description of the plugin.
func (kf *Knife4gf) Description() string {
	return Description
}

// Install installs the swagger to server as a plugin.
// It implements the interface ghttp.Plugin.
func (kf *Knife4gf) Install(s *ghttp.Server) error {
	var (
		ctx = gctx.New()
		//oai = s.GetOpenApi()
	)

	// Retrieve the configuration map and assign it to swagger object.
	m := g.Cfg().MustGet(ctx, "swagger").Map()
	if m != nil {
		if err := gconv.Struct(m, kf); err != nil {
			s.Logger().Fatal(ctx, err)
		}
	}

	var kdocPath string
	kdocPath = g.Cfg().MustGet(ctx, "server.docPath").String()
	if kdocPath == "" {
		kdocPath = docPath
	}

	// The swagger resource files are served as static file service.
	s.AddStaticPath(kdocPath, "resource/swagger")
	// It here uses HOOK feature handling basic auth authentication and swagger.json modification.
	s.Group("/swagger", func(group *ghttp.RouterGroup) {
		group.Hook("/*", ghttp.HookBeforeServe, func(r *ghttp.Request) {
			if kf.BasicAuthUser != "" {
				// Authentication security checks.
				var (
					authCacheKey = fmt.Sprintf(`swagger_auth_failed_%s`, r.GetClientIp())
					authCount    = gcache.MustGet(ctx, authCacheKey).Int()
				)
				if authCount > MaxAuthAttempts {
					r.Response.WriteStatus(
						http.StatusForbidden,
						"max authentication count exceeds, please try again in one minute!",
					)
					r.ExitAll()
				}
				// Basic authentication.
				if !r.BasicAuth(kf.BasicAuthUser, kf.BasicAuthPass) {
					_ = gcache.Set(ctx, authCacheKey, authCount+1, AuthFailedInterval)
					r.ExitAll()
				}
			}
		})
	})
	return nil
}

// Remove uninstalls swagger feature from server.
func (kf Knife4gf) Remove() error {
	return nil
}
