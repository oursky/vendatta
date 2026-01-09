# Technical Specification: Agent Gateway (MCP)

## 1. Overview
The Agent Gateway is a Model Context Protocol (MCP) server built directly into the `oursky` binary. It acts as the bridge between high-level AI agents (Cursor, OpenCode) and the isolated development environment.

## 2. Protocol Details
- **Transport**: JSON-RPC over Standard Input/Output (STDIO).
- **Session Context**: The gateway is started for a specific session: `./oursky agent [session-id]`.

## 3. Capabilities

### **Tool: `exec`**
Executes a command inside the session's environment.
- **Input**: `cmd` (string).
- **Execution**: Routed via `Provider.Exec`.
- **Output**: Combined Stdout/Stderr.

### **Tool Discovery (Future)**
Dynamically loads skills from `.oursky/agents/skills/` and registers them as MCP tools.

## 4. Agent Configuration
Agents are configured to use Oursky via their respective setup files:

### **Cursor**
Cursor uses `.cursorrules` located in the root of the worktree. Oursky generates this file by aggregating rules from `.oursky/agents/rules/`.

### **OpenCode**
OpenCode uses `.opencode/` for skill definitions. Oursky's `sync-agents` command will export YAML definitions from `agents/skills/` to the OpenCode format.
