package auth

import (
	"context"
)

func CreateContext(ctx context.Context, entity AuthEntity) context.Context {
	return context.WithValue(ctx, ContextAuthEntityKey, entity)
}

func GetContext(ctx context.Context) (AuthEntity, bool) {
	user, ok := ctx.Value(ContextAuthEntityKey).(AuthEntity)
	return user, ok
}

func GeneratedUUIDSession(userID string) string {
	return userID + "_SESSION"
}
func GeneratedUUIDCredential(userID string) string {
	return userID + "_CREDENTIAL"
}
