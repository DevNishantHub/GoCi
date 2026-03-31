# Custom CI/CD (From Scratch)

A minimal, high-level CI/CD system built from scratch to understand how modern pipelines work under the hood.

## 🚀 Overview

This project implements a simple CI/CD pipeline:

```
Code → Build → Test → Deploy
```

It is designed for learning purposes, focusing on core concepts like orchestration, execution, and automation.

## ⚙️ Features

* Webhook-based trigger (on code push)
* Pipeline configuration using YAML
* Sequential job execution
* Basic logging of pipeline steps
* Simple deployment script support

## 🏗️ Architecture

* **Webhook Server** – Listens for repository events
* **Pipeline Orchestrator** – Reads and manages pipeline steps
* **Runner** – Executes commands/scripts
* **Logger** – Captures output of each stage

## 📂 Project Structure

```
.
├── main.go          # Entry point
├── pipeline.yaml    # Pipeline configuration
├── runner/          # Execution logic
├── orchestrator/    # Pipeline control logic
├── webhook/         # Webhook handler
└── logs/            # Execution logs
```

## 📌 Example Pipeline

```yaml
stages:
  - build
  - test
  - deploy

build:
  script: go build ./...

test:
  script: go test ./...

deploy:
  script: ./deploy.sh
```

## 🛠️ Getting Started

1. Clone the repo
2. Configure `pipeline.yaml`
3. Run the server:

   ```bash
   go run main.go
   ```
4. Trigger via webhook or manually

## 🎯 Goal

This project is not meant to replace existing CI/CD tools, but to help understand how they work internally.

## 📄 License

MIT
