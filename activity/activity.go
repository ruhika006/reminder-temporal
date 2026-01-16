package activity

import (
	"context"
)

// Activity must return an error
func Sendreminder(ctx context.Context, url string, msg string) error{

	return  Notification(url, msg)
}