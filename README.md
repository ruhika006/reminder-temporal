# Reminder Temporal

A simple drink water reminder application built with Temporal. It demonstrates workflow orchestration with activity execution, allowing you to schedule reminders with a configurable delay.

## How it Works

- **Workflow**: Accepts a message and delay duration, sleeps for the specified delay, then executes the reminder activity.
- **Activity**: Prints "Time to drink water!!!" when triggered.
- **Worker**: Registers and executes workflows and activities from the task queue.
- **Client**: Starts the workflow with a configurable delay (in seconds).


## How to Run

1. Start the Temporal server:
   ```bash
   temporal server start-dev
   ```

2. Start the worker (in a separate terminal):
   ```bash
   go run ./worker/worker.go
   ```

3. Execute the workflow (in another terminal):
   ```bash
   go run main.go <delay_in_seconds>
   ```
   Example: `go run main.go 5` will wait 5 seconds before triggering the reminder.

Output: Time to drink Water!!!