package baggage

import "context"

type Baggage struct {
	Err       error
	ProfileID string
}

type baggageKey struct{}

func WithContext(ctx context.Context, b *Baggage) context.Context {
	return context.WithValue(ctx, baggageKey{}, b)
}

func PutError(ctx context.Context, err error) {
	b, ok := ctx.Value(baggageKey{}).(*Baggage)
	if ok {
		b.Err = err
	}
}

func PutProfileID(ctx context.Context, profileID string) {
	b, ok := ctx.Value(baggageKey{}).(*Baggage)
	if ok {
		b.ProfileID = profileID
	}
}

type postIDKey struct{}

func PutPostID(ctx context.Context, postID string) context.Context {
	return context.WithValue(ctx, postIDKey{}, postID)
}

func GetPostID(ctx context.Context) string {
	if id, ok := ctx.Value(postIDKey{}).(string); ok {
		return id
	}
	return ""
}
