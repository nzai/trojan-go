//go:build api || full
// +build api full

package build

import (
	_ "github.com/nzai/trojan-go/api/control"
	_ "github.com/nzai/trojan-go/api/service"
)
