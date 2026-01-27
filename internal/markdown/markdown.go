package markdown

import (
	"bytes"
	"fmt"
	"regexp"
	"strings"

	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/extension"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/renderer/html"
)

func ToHTML(source []byte, urls ...string) ([]byte, error) {
	var buf bytes.Buffer

	md := goldmark.New(
		goldmark.WithExtensions(
			extension.GFM,
			extension.Typographer,
			extension.Strikethrough,
			extension.Linkify,
		),
		goldmark.WithParserOptions(
			parser.WithAutoHeadingID(),
		),
		goldmark.WithRendererOptions(
			html.WithUnsafe(),
		),
	)

	if err := md.Convert(source, &buf); err != nil {
		return []byte{}, fmt.Errorf("error to convert markdwon to html: %w", err)
	}

	return buf.Bytes(), nil
}

func rewriteURL(src, baseURL string) string {
	if strings.HasPrefix(src, "http") || strings.HasPrefix(src, "/") {
		return src
	}
	cleanPath := src
	for strings.HasPrefix(cleanPath, "../") || strings.HasPrefix(cleanPath, "./") {
		cleanPath = strings.TrimPrefix(cleanPath, "../")
		cleanPath = strings.TrimPrefix(cleanPath, "./")
	}
	return baseURL + cleanPath
}

func ProcessFinalHTML(htmlContent string, baseURL string) string {
	reImg := regexp.MustCompile(`(<img\s+[^>]*src=["'])([^"']*)(["'][^>]*>)`)

	content := reImg.ReplaceAllStringFunc(htmlContent, func(match string) string {
		submatches := reImg.FindStringSubmatch(match)
		if len(submatches) < 4 {
			return match
		}

		oldSrc := submatches[2]
		newSrc := rewriteURL(oldSrc, baseURL)

		return submatches[1] + newSrc + submatches[3]
	})

	reSrc := regexp.MustCompile(`(<source\s+[^>]*srcset=["'])([^"']*)(["'][^>]*>)`)

	content = reSrc.ReplaceAllStringFunc(content, func(match string) string {
		submatches := reSrc.FindStringSubmatch(match)
		if len(submatches) < 4 {
			return match
		}

		oldSrc := submatches[2]
		newSrc := rewriteURL(oldSrc, baseURL)

		return submatches[1] + newSrc + submatches[3]
	})

	return content
}
