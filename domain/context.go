package domain

import "context"

const (
	userContextKey = "user"
)

type UserContextStruct struct {
	NISN interface{}
	Name interface{}
}

func GetUserContext(ctx context.Context) UserContextStruct {
	return ctx.Value(userContextKey).(UserContextStruct)
}

func UserContext(ctx context.Context, value interface{}) context.Context {
	return context.WithValue(ctx, userContextKey, value)
}
