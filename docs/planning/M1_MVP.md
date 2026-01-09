# Milestone: M1_MVP

**Objective**: A single-binary CLI that can initialize a project, spin up a Docker-based isolated codespace with a dedicated worktree, and expose it to an AI agent via MCP.

## 🎯 Success Criteria
- [ ] `./oursky init` scaffolds all necessary files.
- [ ] `./oursky dev <branch>` results in a running container with the worktree mounted.
- [ ] Service ports are discovered and injected as environment variables.
- [ ] Cursor-agent can connect to the environment via the MCP server.

## 🛠 Tasks

| ID | Title | Priority | Status |
| :--- | :--- | :--- | :--- |
| **[INF-01](./tasks/INF-01.md)** | Docker Provider implementation (DinD, Ports) | 🔥 High | [Completed] |
| **[COR-01](./tasks/COR-01.md)** | Orchestration Engine (Worktree + Hooks) | 🔥 High | [Completed] |
| **[AGT-01](./tasks/AGT-01.md)** | MCP Server implementation | ⚡ Med | [Completed] |
| **[CLI-01](./tasks/CLI-01.md)** | CLI Scaffolding & Agent Sync | ⚡ Med | [In Progress] |
| **[VFY-01](./tasks/V2Y-01.md)** | E2E Verification with Docker-Compose | 🔥 High | [Pending] |
