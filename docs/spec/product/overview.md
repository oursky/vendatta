# Product Overview: Project Oursky

## 1. Vision
Oursky is a developer-centric, local-first development environment manager. It aims to eliminate the "it works on my machine" problem by providing high-isolation, reproducible codespaces that are natively compatible with modern AI agents (Cursor, OpenCode).

## 2. Problem Statement
- **Environment Drift**: Developers struggle with conflicting node/python versions across projects.
- **Microservice Complexity**: Local development often requires complex `docker-compose` setups that are hard to isolate.
- **Agent Friction**: AI agents lack a standard, secure way to execute tools and understand environment-specific rules.
- **Bloat**: Existing solutions are often "heavy" (cloud-only) or "fragmented" (manual worktrees + manual docker).

## 3. Core Value Propositions

### **High Isolation (LXC/Docker + Worktrees)**
Unlike standard `docker-compose` which shares the host filesystem, Oursky uses **Git Worktrees** to provide a unique, branch-specific filesystem for every session. This prevents file-locking and allows parallel work on multiple branches without pollution.

### **BYOA (Bring Your Own Agent)**
Oursky provides a standardized **Model Context Protocol (MCP)** gateway. Any agent (Cursor, OpenCode, Claude) can connect to an Oursky session and immediately inherit:
- **Environment-aware Tools**: (e.g., `exec` inside the container).
- **Project-specific Rules**: Defined in `.oursky/agents/rules`.
- **Role-based Capabilities**: (e.g., `frontend` vs `backend` roles).

### **Single-Binary Portability**
Zero-setup installation. A single Go binary manages everything from worktree creation to container orchestration and port forwarding.

## 4. Target Personas

| Persona | Needs | Oursky Solution |
| :--- | :--- | :--- |
| **Senior Dev** | Complex orchestration, fast branch switching. | Automated Worktree + DinD orchestration. |
| **Agent User** | Secure tool execution for AI. | Built-in MCP Server with session boundaries. |
| **Team Lead** | Onboarding consistency. | Standardized `.oursky` config & lifecycle hooks. |

## 5. The "Oursky" Workflow
1.  **Onboard**: Run `oursky init`. Define ports and setup hooks.
2.  **Dev**: Run `oursky dev feature-x`. A clean worktree and container appear.
3.  **Code**: Open the worktree in Cursor. The agent connects to the MCP server.
4.  **Ship**: Verify changes in the isolated environment and commit from the worktree.
5.  **Clean**: Run `oursky kill`. All resources are wiped.
