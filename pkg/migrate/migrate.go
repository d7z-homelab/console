package migrate

import (
	v1 "github.com/d7z-homelab/console/pkg/migrate/v1"
)

func init() {
	add(v1.InitSystem)
}
