# Custom CI/CD (From Scratch)

A minimal CI/CD system built from scratch to understand how modern pipelines work internally.

---

## 🚀 Overview

This project implements a simple CI/CD pipeline:

Code → Build → Test → Deploy

It is designed for learning purposes and focuses on core concepts like orchestration, execution, and automation.

---

## ⚙️ Features

- Webhook-based trigger (on code push)
- YAML-based pipeline configuration
- Sequential job execution
- Basic logging system
- Simple deployment support

---

## 🏗️ Architecture

The system consists of:

- **Webhook Server** – Listens for repository events  
- **Orchestrator** – Parses and controls pipeline execution  
- **Runner** – Executes commands/scripts  
- **Logger** – Captures logs for each stage  

---

## 📂 Project Structure

```
.
├── main.go
├── pipeline.yaml
├── webhook/
├── orchestrator/
├── runner/
└── logs/
```

---

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
🛠️ Getting Started

-Clone the repository

-Configure pipeline.yaml

Run the server:
```go run main.go```

Trigger the pipeline via webhook or manually

## 🎯 Goal

This project is intended for learning purposes only.
It helps in understanding how CI/CD systems like Jenkins, GitHub Actions, or GitLab CI work internally.
