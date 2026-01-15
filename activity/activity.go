package activity

import (
	"context"
	"fmt"
)

// Activity must return an error
func Sendreminder(ctx context.Context, msg string) error {

	fmt.Println(msg)
	return nil
	/* Temporal will retry the activity a couple of times automatically (durable execution). 
	 return fmt.Errorf("activity failed intentionally")
	*/
}