package markdown

import (
	"fmt"
	"strings"

	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/ast"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/text"
	"github.com/yuin/goldmark/util"
)

type pathRewriter struct {
	BaseURL string
}

func (p *pathRewriter) Transform(node *ast.Document, reader text.Reader, pc parser.Context) {
	// En Goldmark, ast.Walk recibe el nodo y una función que retorna (ast.WalkStatus, error)
	err := ast.Walk(node, func(n ast.Node, entering bool) (ast.WalkStatus, error) {
		if !entering {
			return ast.WalkContinue, nil
		}

		// Verificamos si el nodo es una imagen
		if n.Kind() == ast.KindImage {
			img := n.(*ast.Image)
			src := string(img.Destination)

			if !strings.HasPrefix(src, "http") && !strings.HasPrefix(src, "/") {
				// Limpiamos la ruta y concatenamos
				cleanPath := src
				for strings.HasPrefix(cleanPath, "../") || strings.HasPrefix(cleanPath, "./") {
					cleanPath = strings.TrimPrefix(cleanPath, "../")
					cleanPath = strings.TrimPrefix(cleanPath, "./")
				}
				img.Destination = []byte(p.BaseURL + cleanPath)
			}
		}

		return ast.WalkContinue, nil
	})

	if err != nil {
		fmt.Printf("Error caminando el AST: %v\n", err)
	}
}

// Implementación de la extensión
type rewriteExtension struct {
	BaseURL string
}

func ImgSrc(baseURL string) *rewriteExtension {
	return &rewriteExtension{
		BaseURL: baseURL,
	}
}

func (e *rewriteExtension) Extend(m goldmark.Markdown) {
	m.Parser().AddOptions(
		parser.WithASTTransformers(
			util.Prioritized(&pathRewriter{BaseURL: e.BaseURL}, 100),
		),
	)
}
