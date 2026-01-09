# Task: CLI-01 CLI Scaffolding & Agent Sync

**Priority**: ⚡ Med
**Status**: [In Progress]

## 🎯 Objective
Refine the `init` command and implement the logic to sync generic agent rules to agent-specific configurations.

## 🛠 Implementation Details
1.  **Init Scaffolding**:
    - Create `.oursky/agents/{rules,skills,commands}`.
    - Create default `config.yaml` and `setup.sh`.
2.  **Agent Sync Logic**:
    - Read all `.md` files in `agents/rules/`.
    - Concatenate into a single string.
    - Write to `.cursorrules` in the workspace root.
3.  **Command UX**:
    - Implement `oursky sync-agents`.

## 🧪 Proof of Work
- `.cursorrules` is generated from the content of `.oursky/agents/rules/`.
- All scaffolded directories are present after `oursky init`.
