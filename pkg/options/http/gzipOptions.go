package http

import (
	"github.com/baoyxing/micro-extend/pkg/config/hertz_conf"
	hgzip "github.com/hertz-contrib/gzip"
	"strings"
)

func GzipOptions(server hertz_conf.Server) []hgzip.Option {
	var options []hgzip.Option

	if server.Gzip.Excluded.Enable {

		if server.Gzip.Excluded.ExcludedExtensions.Enable {
			options = append(options, hgzip.WithExcludedExtensions(strings.Split(server.Gzip.Excluded.ExcludedExtensions.Extensions, ",")))
		}
		if server.Gzip.Excluded.ExcludedPaths.Enable {
			options = append(options, hgzip.WithExcludedExtensions(strings.Split(server.Gzip.Excluded.ExcludedPaths.Paths, ",")))
		}
		if server.Gzip.Excluded.ExcludedPathRegexes.Enable {
			options = append(options, hgzip.WithExcludedExtensions(strings.Split(server.Gzip.Excluded.ExcludedPathRegexes.Regexes, ",")))
		}
	}
	return options
}
