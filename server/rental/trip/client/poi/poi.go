package poi

import (
	"context"
	rentalpb "coolcar/proto/rental/gen/go"
)

type Manager struct {
}

func (m *Manager) Resolve(context.Context, *rentalpb.Location) (string, error) {
	return "陆家嘴", nil
}
