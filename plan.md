# Project Plan: Go REST Request Generator CLI

This plan outlines the steps to build a CLI tool in Go that can send configurable REST requests to a backend, supporting simple data generation and more complex traffic scenarios.

## 1. Foundational Setup

*   **Goal:** Initialize the Go project and set up the basic structure.
*   **Actions:**
    1.  Initialize a Go module: `go mod init tg`
    2.  Create a standard project layout:
        ```
        /cmd/tg/         # Main application entry point
        /internal/       # Core application logic
        /internal/cli/   # Command definitions (using Cobra)
        /internal/http/  # HTTP client helpers
        /internal/data/  # Data generation logic
        /configs/        # Example configuration templates
        ```
    3.  Install the `cobra` library, which is excellent for building powerful CLIs: `go get -u github.com/spf13/cobra@latest`

## 2. CLI Command Structure

*   **Goal:** Define the user-facing commands for the CLI.
*   **Actions:**
    *   **Root Command (`tg`):** The base command for the application.
    *   **`entity` command:**
        *   `create`: A subcommand to create a new entity on the backend.
            *   Example: `tg entity create --name "sensor-01" --type "temperature"`
    *   **`measurement` command:**
        *   `send`: A subcommand to send measurement data.
            *   **Flags:**
                *   `--entity-id`: The ID of the entity to associate data with.
                *   `--year`: The year to generate data for (e.g., 2020).
                *   `--value`: A fixed value to send for each day.
                *   `--random`: A boolean flag to use random values between 1-100.
                *   `--config`: Path to a template file defining the measurement parameters.
                *   `--backend-url`: The base URL of your backend API.
            *   Example: `tg measurement send --entity-id 123 --year 2020 --random --config ./configs/measurement.yaml --backend-url "http://localhost:8080"`

## 3. Configuration Templates

*   **Goal:** Allow users to define the structure of the data they want to send.
*   **Actions:**
    *   Use YAML for its readability. We will create a simple template that defines the static parts of a measurement.
    *   **Example `configs/measurement.yaml`:**
        ```yaml
        type: "temperature"
        unit: "celsius"
        description: "Daily average"
        ```
    *   The CLI will read this file, combine it with the generated daily data (timestamp, value), and send it as a JSON payload.

## 4. Core Implementation Logic

*   **Goal:** Write the Go code to power the CLI commands.
*   **Actions:**
    1.  **HTTP Client (`/internal/http`):** Create a reusable client to handle POST requests, including setting headers and handling responses and errors.
    2.  **Data Generation (`/internal/data`):**
        *   Implement a function to loop through each day of a given year.
        *   Implement a function to generate either a fixed value or a random integer in a specified range.
    3.  **Command Logic (`/internal/cli`):**
        *   For the `measurement send` command, orchestrate the process:
            1.  Parse command-line flags.
            2.  Read and parse the measurement config file.
            3.  Loop through each day of the specified year.
            4.  For each day, generate the value.
            5.  Construct the final JSON payload.
            6.  Use the HTTP client to send the request to the backend.

## 5. Supporting Complex Traffic Scenarios

*   **Goal:** Extend the tool to support sequences of different requests to simulate more realistic user traffic.
*   **Actions:**
    1.  **Introduce a `scenario` command:** `tg scenario run --file ./configs/scenario.yaml`
    2.  **Define a Scenario Template:** Create a YAML format to define a series of steps. This allows for defining a flow of API calls.
    3.  **Example `configs/scenario.yaml`:**
        ```yaml
        name: "Onboard New Sensor and Send Data"
        variables:
          entityName: "sensor-02"
          backend: "http://localhost:8080"
        
        steps:
          - name: "Create a new entity"
            request:
              method: "POST"
              url: "{{.backend}}/entities"
              body:
                name: "{{.entityName}}"
                type: "humidity"
            register: entity_id # Save the 'id' from the response
        
          - name: "Send initial measurement"
            request:
              method: "POST"
              url: "{{.backend}}/measurements"
              body:
                entity_id: "{{.entity_id}}" # Use the saved id
                value: 55
                unit: "percent"
        ```
    4.  **Scenario Runner:** Implement a runner that parses this YAML, executes steps sequentially, and uses Go's templating engine to substitute variables between steps.
