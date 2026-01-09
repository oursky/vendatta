# Testing Strategy: Project Oursky

## 1. Overview
Project Oursky is a systems-level CLI tool. Testing must cover both high-level orchestration logic and low-level provider interactions (Docker, Git).

## 2. Test Tiers

### **Tier 1: Unit Tests (Fast, Hermetic)**
- **Scope**: Configuration parsing, URL calculation, Rule aggregation.
- **Mocking**: Use interfaces for `Provider` and `WorktreeManager`.
- **Target**: `pkg/config`, `pkg/agent`, logic in `pkg/ctrl`.

### **Tier 2: Integration Tests (Medium, Requires Environment)**
- **Scope**: Real Docker container creation, Real Git Worktree commands.
- **Environment**: Requires a local Docker daemon and Git installed.
- **Target**: `pkg/provider/docker`, `pkg/worktree`.

### **Tier 3: E2E Verification (Slow, Full Flow)**
- **Scope**: The entire `oursky init` -> `oursky dev` -> `oursky agent` lifecycle.
- **Verification**: Use a test repository (fixture) and verify that:
    - Files are created.
    - Environment variables are present in the container.
    - MCP tools return successful responses.

## 3. Tooling
- **Go Test**: Primary test runner.
- **Docker-in-Docker (GitHub Actions)**: For CI/CD verification.
- **Test Fixtures**: Located in `internal/testfixtures/`.

## 4. Quality Gates
- **Zero Lint Errors**: Verified via `golangci-lint`.
- **Diagnostics Clean**: `lsp_diagnostics` must pass on all changed files.
- **E2E Pass**: Critical for every PR affecting `pkg/ctrl` or `pkg/provider`.

## 5. Mocking Strategy
```go
// Example Mock Provider for Ctrl Tests
type MockProvider struct {
    CreatedSessions []string
}

func (m *MockProvider) Create(...) (*provider.Session, error) {
    m.CreatedSessions = append(m.CreatedSessions, sessionID)
    return &provider.Session{ID: sessionID}, nil
}
```
