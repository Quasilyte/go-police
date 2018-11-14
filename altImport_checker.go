package police

import (
	"fmt"
	"go/ast"
	"strconv"

	"github.com/go-lintpack/lintpack"
)

func init() {
	var info lintpack.CheckerInfo
	info.Name = "altImport"
	info.Tags = []string{"style"}
	info.Summary = "Detect discouraged imports and propose an alternative"
	info.Before = `import "errors"`
	info.After = `import "github.com/foobar/errors"`

	lintpack.AddChecker(&info, func(ctx *lintpack.CheckerContext) lintpack.FileWalker {
		requireConfig()
		return &altImportChecker{
			ctx:   ctx,
			rules: config.AltImport,
		}
	})
}

type altImportChecker struct {
	ctx   *lintpack.CheckerContext
	rules map[string]string
}

func (c *altImportChecker) WalkFile(f *ast.File) {
	for _, imp := range f.Imports {
		quoted := imp.Path.Value
		path, err := strconv.Unquote(quoted)
		if err != nil { // Practically impossible
			panic(fmt.Sprintf("unquote import path: %v", err))
		}
		if suggestion := c.rules[path]; suggestion != "" {
			c.ctx.Warn(imp, "don't import %s, use %s instead", quoted, suggestion)
		}
	}
}
