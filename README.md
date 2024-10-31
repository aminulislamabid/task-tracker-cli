# Task Tracker CLI

`task-tracker-cli` is a command-line application for managing tasks with functionality for adding, listing, and tracking the progress of various tasks. This project is built using Go and is easy to run and test through the Makefile commands provided. **_Use positional arguments in command line to accept user inputs._**

## Features

- Add Task, Update, Delete.
- Track the status of tasks (`todo`, `in-progress`, `done`).
- All List, List by status.
- Test coverage and testing utilities to ensure quality.
- Easy build, run, and test automation via Makefile.

## Project Structure

- **`cmd/main.go`**: The entry point for the application.
- **`internal/`**: Contains internal package files:
  - **`constants`**: Stores constants used across the application.
  - **`models`**: Defines the `Task` model.
  - **`storage`**: Provides functions for saving and loading tasks.
  - **`task`**: Contains functionalities for task-specific operations.
  - **`utils`**: Contains utilities for setup and file operations.

## Requirements

- **Go**: Ensure that Go is installed on your system.
- **Make**: The project uses Makefile to manage build, run, and test commands.

## Usage - Makefile Commands

**_You can run this application as you wish_** but using a `Makefile` simplifies common commands.

- **Build the Application**

  To compile the project:

  ```bash
  make build
  ```

- **Run the Application**

  To run the application with optional arguments:

  ```bash
  make run ARGS="your-arguments-here"
  ```

  **_Example:_**

  ```bash
  make run ARGS="add coffee-at-9AM"
  ```

- **Run Test**

  To run all tests in the project:

  ```bash
  make test
  ```

  Generate coverage.out file:

  ```bash
  go test ./... -coverprofile=your_absolute_path\task-tracker-cli\coverage.out
  ```

  Generate coverage.html file

  ```bash
  go tool cover -html=your_absolute_path\task-tracker-cli\coverage.out -o your_absolute_path\task-tracker-cli\coverage.html
  ```
