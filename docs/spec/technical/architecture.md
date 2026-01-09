# Technical Architecture: Project Oursky

## 1. Overview
Oursky is a developer-centric, single-binary dev environment manager. It abstracts complex infrastructure (Docker, LXC, Worktrees) into a simple CLI interface, providing isolated, reproducible, and agent-friendly codespaces.

## 2. Component Diagram

```mermaid
graph TD
    CLI[Oursky CLI] --> CP[Control Plane]
    CP --> WM[Worktree Manager]
    CP --> P[Provider Interface]
    
    subgraph "Execution Plane"
        P --> Docker[Docker Provider]
        P --> LXC[LXC Provider]
        Docker --> DinD[Docker-in-Docker / Compose]
    end
    
    subgraph "Agent Plane"
        CLI --> MCP[MCP Server]
        MCP --> CP
        UserAgent[Cursor / OpenCode] <-->|JSON-RPC| MCP
    end

    subgraph "Filesystem"
        WM --> WT[.oursky/worktrees/]
        CP --> CFG[.oursky/config.yaml]
        CP --> AS[.oursky/agents/]
    end
```

## 3. Core Modules

### **Control Plane (`pkg/ctrl`)**
The central coordinator. It is responsible for:
- Parsing `.oursky/config.yaml`.
- Orchestrating the sequence: `Worktree Create` -> `Provider Create` -> `Setup Hook` -> `Agent Gateway`.
- Maintaining session state through Docker labels and filesystem markers.

### **Provider Interface (`pkg/provider`)**
An abstraction layer for environment lifecycles.
- **Methods**: `Create`, `Start`, `Stop`, `Destroy`, `Exec`, `List`.
- **MVP Implementation**: `DockerProvider` with DinD support and absolute path bind-mounting.

### **Worktree Manager (`pkg/worktree`)**
Automates the management of `git worktree`.
- Ensures that every dev session has a clean, isolated filesystem.
- Maps the worktree directory into the provider's execution context.

### **Agent Gateway (`pkg/agent`)**
Implements the **Model Context Protocol (MCP)**.
- **Tools**: `exec` (primary), `read_file`, `write_file` (via CP).
- **Interoperability**: Designed to be the standard interface for Cursor-agent, OpenCode, and Claude.

## 4. Environment Injection & Networking
Oursky solves the API discovery problem by injecting environment variables:
- **Port Discovery**: Host-mapped ports are discovered dynamically.
- **Injection**: Variables like `OURSKY_SERVICE_[NAME]_URL` are passed to the container, allowing seamless CORS and endpoint configuration.

## 5. Agent Scaffold
The `.oursky/agents/` directory acts as the **Single Source of Truth** for agent behavior.
- **Rules**: Markdown-based instructions.
- **Skills**: YAML tool definitions.
- **Sync**: CLI command `sync-agents` generates agent-specific configurations (e.g., `.cursorrules`).
