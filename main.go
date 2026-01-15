package main

import (
	"context"
	"log"
	"os"
	"strconv"
	"time"

	workflow "github.com/ruhika006/reminder-temporal/workflow"
	"go.temporal.io/sdk/client"
)


func main(){

	c, err := client.Dial(client.Options{})
	if err != nil {
		log.Fatalln("Unable to create client", err)
	}
	defer c.Close()

	options := client.StartWorkflowOptions{
		TaskQueue: "my-task-queue",
	}

	// Parse the delay duration from command-line argument (in seconds)
	delaySeconds, err := strconv.Atoi(os.Args[1])
	if err != nil {
		log.Fatalln("Invalid delay argument:", err)
	}
	delay := time.Duration(delaySeconds) * time.Second

	wo, err := c.ExecuteWorkflow(context.Background(), options, workflow.DrinkWaterReminderWorkflow, "drink water", delay)
	if err!=nil{
		log.Fatalln("error while executing workflow: ", err)
	}

	var result string

	err = wo.Get(context.Background(), &result)
	if err!=nil{
		log.Fatalln("unable to get workflow result", err)
	}

	log.Println("Workflow started: ", result)

}