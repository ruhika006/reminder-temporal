package main

import (
	"log"

	activity "github.com/ruhika006/reminder-temporal/activity"
	workflow "github.com/ruhika006/reminder-temporal/workflow"

	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"
)


func main(){

	c, err := client.Dial(client.Options{})
	if err != nil {
		log.Fatalln("Unable to create client", err)
	}
	defer c.Close()


	w := worker.New(c, "my-task-queue", worker.Options{})
	w.RegisterActivity(activity.Sendreminder)
	w.RegisterWorkflow(workflow.DrinkWaterReminderWorkflow)


	err = w.Run(worker.InterruptCh())
	if err!=nil{
		log.Fatalln("unable to run the workflow", err)
	}
}