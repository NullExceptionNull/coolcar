package car

import (
	"context"
	"coolcar/server/shared/id"
)

type Manager struct {
}

func (m *Manager) Verify(ctx context.Context, carID id.CarID) error {
	return nil
}

func (m *Manager) UnLock(ctx context.Context, carID id.CarID) error {
	return nil
}
