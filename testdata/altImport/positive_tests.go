package checker_test

import (
	/// don't import "errors", use fmt.Errorf instead
	"errors"

	/// don't import "path", use path/filepath package instead
	"path"
)

var (
	_ = errors.New
	_ = path.Join
)
