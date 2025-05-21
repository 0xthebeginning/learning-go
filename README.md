![Build Go Project](https://github.com/0xthebeginning/learning-go/actions/workflows/go.yml/badge.svg)

Practicing Go to deepen my systems and tooling skills.


# learning-go

A minimal Go application that runs a basic HTTP server and is containerized and deployed to Kubernetes.

---

## 🧠 About

This is a personal learning project to explore Go, Docker, Kubernetes, and CI pipelines.  
The app listens on port `8080` and responds with a simple message at the root path (`/`).

---

## 🛠️ Stack

- **Go** – `hello.go`
- **Docker** – Containerized the Go app
- **Kubernetes** – Deployed locally using Minikube
- **YAML** – `deployment.yaml` and `service.yaml` for Kubernetes

---

## 🚀 How to Run Locally (with Minikube)

1. Start Minikube:
   ```bash
   minikube start --driver=docker
