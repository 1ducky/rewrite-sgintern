package auth

import (
	"context"

	"github.com/google/uuid"
)

func CreateContext(ctx context.Context, entity AuthEntity) context.Context {
	return context.WithValue(ctx, ContextAuthEntityKey, entity)
}

func GetContext(ctx context.Context) (AuthEntity, bool) {
	user, ok := ctx.Value(ContextAuthEntityKey).(AuthEntity)
	return user, ok
}

func GeneratedUUIDSession(userID string) string {
	return "SESSION_" + userID + "_" + uuid.New().String()
}
func GeneratedUUIDCredential(userID string) string {
	return "CREDENTIAL_" + userID + "_" + uuid.New().String()
}
