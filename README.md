# tg - Go REST Request Generator CLI

`tg` is a command-line tool built with Go for generating and sending REST requests to a backend. It's designed for testing, data seeding, and simulating traffic patterns.

## Features

-   Create entities on your backend.
-   Send measurement data for specific entities.
-   Generate data for a full year with either fixed or random values.
-   Use YAML templates to define request bodies.
-   Run complex scenarios involving multiple API calls in sequence.
-   Substitute variables between scenario steps.

## Installation

To get started, you need to have Go installed on your system.

1.  **Clone the repository:**
    ```bash
    git clone https://github.com/marasonic/tg.git
    cd tg
    ```

2.  **Build the executable:**
    ```bash
    go build -o tg ./cmd/tg
    ```
    This will create an executable file named `tg` in the project root. You can also use `go run` to execute commands directly without building.

## Usage

The CLI is structured around a few main commands: `entity`, `measurement`, and `scenario`.

### `entity`

Manage entities on the backend.

#### `entity create`

Create a new entity.

```bash
./tg entity create --name <entity-name> --type <entity-type> [--backend-url <url>]
```

**Example:**

```bash
./tg entity create --name "sensor-main-hall" --type "temperature" --backend-url "http://api.example.com"
```

### `measurement`

Send measurement data.

#### `measurement send`

Send measurement data for a specific entity over a period of time.

```bash
./tg measurement send --entity-id <id> --year <year> [--value <val> | --random] [--config <path>] [--backend-url <url>]
```

**Flags:**

-   `--entity-id` (required): The ID of the entity to associate the data with.
-   `--year`: The year to generate data for (default: 2020).
-   `--value`: A fixed numeric value to send for each day.
-   `--random`: If set, sends a random value between 1 and 100 for each day.
-   `--config`: Path to a YAML file to use as a template for the request body.
-   `--backend-url`: The base URL of your backend API (default: `http://localhost:8080`).

**Examples:**

1.  **Send a fixed value for each day in 2021:**
    ```bash
    ./tg measurement send --entity-id 123 --year 2021 --value 42
    ```

2.  **Send random values using a config template:**
    ```bash
    ./tg measurement send --entity-id 456 --year 2020 --random --config ./configs/measurement.yaml
    ```

### `scenario`

Run complex, multi-step traffic scenarios.

#### `scenario run`

Execute a scenario defined in a YAML file.

```bash
./tg scenario run --file <path-to-scenario.yaml>
```

**Example:**

```bash
./tg scenario run --file ./configs/scenario.yaml
```

## Configuration

### Measurement Template

You can define the static parts of a measurement's JSON payload in a YAML file. The CLI reads this file and merges it with the dynamically generated data (like `timestamp` and `value`).

**Example `configs/measurement.yaml`:**

```yaml
type: "temperature"
unit: "celsius"
description: "Daily average reading"
```

When you run the `measurement send` command with this config, the final JSON payload for each request will look something like this:

```json
{
  "description": "Daily average reading",
  "entity_id": "456",
  "timestamp": "2020-01-01T00:00:00Z",
  "type": "temperature",
  "unit": "celsius",
  "value": 87
}
```

### Scenario Template

Scenarios allow you to define a sequence of API calls to simulate a workflow. This is useful for integration testing or simulating user behavior.

**Example `configs/scenario.yaml`:**

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
    register: entity_id # Saves the 'id' from the response body

  - name: "Send initial measurement"
    request:
      method: "POST"
      url: "{{.backend}}/measurements"
      body:
        entity_id: "{{.entity_id}}" # Uses the value saved from the previous step
        value: 55
        unit: "percent"
```

**How it works:**

-   **`variables`**: Define static variables that can be used in any step.
-   **`steps`**: An array of requests to be executed in order.
-   **`register`**: After a step runs, you can save a value from the JSON response body into a temporary variable. In the example, it saves the `id` field from the "Create a new entity" response.
-   **Templating**: The `url` and `body` fields support Go's templating syntax (`{{.variableName}}`) to substitute values from the `variables` block or from values registered by previous steps.
