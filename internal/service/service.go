package service

import (
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/text/gstr"
)

const (
	apiDocName           = `{ApiDocName}`
	apiDocURL            = `{ApiDocURL}`
	servicesJsonTemplate = `[
    {
        "name": "{ApiDocName}",
        "url": "{ApiDocURL}",
        "swaggerVersion": "2.0",
        "location": "{ApiDocURL}"
    },
]
`
)

func ApiServices(s *ghttp.Server) string {
	var (
		oai     = s.GetOpenApi()
		content = servicesJsonTemplate
	)

	content = gstr.ReplaceByMap(servicesJsonTemplate, map[string]string{
		"{ApiDocName}": oai.Info.Title,
		"{ApiDocURL}":  "/api.json",
	})
	return content
}
