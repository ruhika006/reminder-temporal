# Reminder Temporal

A simple drink water slack channel reminder application built with Temporal. It demonstrates workflow orchestration with activity execution, allowing you to schedule reminders with a configurable delay.

## How it Works

- **Workflow**: Accepts a message and delay duration, sleeps for the specified delay, then executes the reminder activity.
- **Activity**: Prints "Time to drink water!!!" when triggered.
- **Worker**: Registers and executes workflows and activities from the task queue.
- **Client**: Starts the workflow with a configurable delay (in seconds).

## Prerequisite:
- Create a Slack incoming webhook for your desired channel
- Add the webhook URL to `.env` as `SLACK_WEBHOOK_URL=<your-url>`

## How to Run

1. Start the Temporal server:
   ```bash
   temporal server start-dev
   ```

2. Start the worker (in a separate terminal):
   ```bash
   go run ./worker/worker.go
   ```

3. Execute the workflow and all files(notification.go) (in another terminal):
   ```bash
   go run . 2 <delay_in_seconds>
   ```
   Example: `go run . 5` will wait 5 seconds before triggering the reminder.

Slack channel message: Reminder: Time to drink Water!!!