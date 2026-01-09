# CLI UX Specification: Project Oursky

## 1. Design Philosophy
- **Speed**: Commands should provide immediate feedback (sub-second for information, clear progress for IO).
- **Clarity**: Status updates should use colors and clear labels.
- **Actionability**: Error messages must suggest a solution (e.g., "Docker not running. Start Docker and try again.").
- **Scriptability**: Global `--json` flag for all commands to allow integration into other tools.

## 2. Command Structure

| Command | Usage | Feedback Pattern |
| :--- | :--- | :--- |
| `init` | `oursky init` | Interactive prompts + Success checklist. |
| `dev` | `oursky dev <branch>` | Progress bars for Image Pull & Worktree creation. |
| `list` | `oursky list` | Tabular data with color-coded status (Active=Green). |
| `agent` | `oursky agent <id>` | Silent (Stdio mode) but logs to `.oursky/logs/agent.log`. |
| `kill` | `oursky kill <id>` | Explicit confirmation before destructive actions. |

## 3. Feedback Elements

### **Progress Indicators**
For long-running tasks like `Image Pull` or `Worktree Setup`:
```text
[1/3] Creating worktree 'feature-x'... OK
[2/3] Pulling docker image 'node:20'... [=====>    ] 60%
[3/3] Running setup hook... 
```

### **Error Handling**
Errors should follow the **Context-Problem-Solution** pattern:
```text
Error: Failed to bind-mount worktree.
Problem: The directory '/home/user/repo' is not shared with Docker.
Solution: Add this directory to Docker Desktop > Settings > Resources > File Sharing.
```

## 4. Visual Language
- **Accent Color**: Sky Blue (`#00BFFF`).
- **Success**: Green Checkmark (`✔`).
- **Warning**: Yellow Triangle (`⚠`).
- **Error**: Red Cross (`✖`).

## 5. Agent Interoperability UX
- `oursky agent` should be robust. If the connection drops, it should auto-recover or provide a clean restart capability.
- `oursky sync-agents` should show a diff of what rules are being updated in `.cursorrules`.
