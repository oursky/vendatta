# Oursky (Vibegear Rework)

Oursky is a developer-centric, single-binary dev environment manager. It abstracts complex infrastructure into a simple CLI, providing isolated, reproducible, and AI-agent-friendly codespaces using local providers (Docker, LXC).

## Key Features

- **Single Binary**: No complex dependencies on the host machine.
- **Git Worktree Managed**: Automatic isolation of code for every branch/session.
- **BYOA (Bring Your Own Agent)**: Built-in Model Context Protocol (MCP) server for Cursor, OpenCode, and Claude.
- **Service Discovery**: Automatic port mapping and environment variable injection (`OURSKY_SERVICE_WEB_URL`, etc.).
- **Docker-in-Docker**: Seamless support for `docker-compose` projects inside isolated environments.

## Quick Start

### 1. Installation
Build the binary using Go (requires Go 1.24+):
```bash
go build -o oursky cmd/oursky/main.go
```

### 2. Onboard a Project
Initialize the `.oursky` configuration in your repository:
```bash
./oursky init
```

### 3. Start Developing
Spin up an isolated environment for a feature branch:
```bash
./oursky dev feature-login
```

## Dogfooding (Developing Oursky with Oursky)

Oursky is designed to be self-hosting. To develop the `oursky` project itself using an isolated environment:

1.  **Build the CLI**:
    ```bash
    go build -o oursky cmd/oursky/main.go
    ```
2.  **Initialize Oursky on Oursky**:
    ```bash
    ./oursky init
    ```
3.  **Spin up a dev session**:
    ```bash
    ./oursky dev main
    ```
    *This creates a worktree at `.oursky/worktrees/main`, starts a container, and mounts it.*
4.  **Connect your Agent**:
    If using Cursor, open the folder `.oursky/worktrees/main`.
5.  **Verify Changes**:
    Since the Docker socket is bind-mounted (DinD), you can run `./oursky` commands *inside* the Oursky container to spawn further sub-environments or run tests.

## Project Structure

- `pkg/ctrl`: Control Plane - Orchestration and lifecycle logic.
- `pkg/provider`: Execution Plane - Environment abstraction (Docker, LXC).
- `pkg/worktree`: Filesystem isolation using Git Worktrees.
- `pkg/agent`: Agent Plane - MCP server and SSH gateway.
- `docs/`: Technical specifications and task-based planning.

## Roadmap

See [docs/planning/README.md](./docs/planning/README.md) for milestones and tasks.

- **M1: CLI MVP** (Current) - Docker + Worktree + MCP.
- **M2: Alpha** - LXC Provider, Advanced Port Forwarding.
- **M3: Beta** - QEMU/Virtualization for macOS/Windows.

---
*Powered by OhMyOpenCode.*
