package activity

import (
	"context"
)

// Activity must return an error
func Sendreminder(ctx context.Context, msg string) (string, error) {

	
	return msg, nil
	/* Temporal will retry the activity a couple of times automatically (durable execution). 
	 return fmt.Errorf("activity failed intentionally")
	*/
}