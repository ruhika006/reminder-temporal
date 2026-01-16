package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
	"github.com/ruhika006/reminder-temporal/workflow"
	"go.temporal.io/sdk/client"
)

func main() {
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Println("Warning: .env file not found, using system environment variables")
	}

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

	// Slack
	url := os.Getenv("SLACK_WEBHOOK_URL")
	if url == "" {
		log.Fatalln("SLACK_WEBHOOK_URL environment variable not set")
	}

	wo, err := c.ExecuteWorkflow(context.Background(), options, workflow.DrinkWaterReminderWorkflow, url, "Time to drink Water!!!", delay)
	if err != nil {
		log.Fatalln("error while executing workflow: ", err)
	}

	var result string

	err = wo.Get(context.Background(), &result)
	if err != nil {
		log.Fatalln("unable to get workflow result", err)
	}
	
	fmt.Println(result)
}
