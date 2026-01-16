package workflow

import (
	"log"
	"time"

	"go.temporal.io/sdk/workflow"

	"github.com/ruhika006/reminder-temporal/activity"
)


func DrinkWaterReminderWorkflow(ctx workflow.Context, message string, delay time.Duration) (string, error){

	
	ao := workflow.ActivityOptions{
		StartToCloseTimeout: time.Minute,
	}

	ctx = workflow.WithActivityOptions(ctx, ao)

	err := workflow.Sleep(ctx, delay)
	if err!=nil{
		log.Fatalln("unable to sleep", err)
	}
	var result string
	err = workflow.ExecuteActivity(ctx, activity.Sendreminder, message).Get(ctx, &result )
	if err!=nil{
		log.Fatalln(err)
	}
	return result, nil
}