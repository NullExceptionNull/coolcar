package profile

import (
	"context"
	"coolcar/server/shared/id"
)

type Manager struct {
}

func (m Manager) Verify(ctx context.Context, accountID string) (id.IdentityId, error) {
	return id.IdentityId("0702"), nil
}
