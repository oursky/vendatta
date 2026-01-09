# User Stories: Project Oursky

## 1. Feature Development (The Main Path)
> *As a full-stack developer, I want to start working on a new feature branch in an isolated environment so that my local machine stays clean and I can switch between branches instantly.*

- **Acceptance Criteria**:
    - Running `oursky dev branch-name` creates a dedicated worktree.
    - Dependencies are installed inside a container, not on the host.
    - My existing `docker-compose` setup for the database works inside the Oursky environment (DinD).

## 2. Agent Collaboration (BYOA)
> *As a developer using Cursor, I want my AI agent to be able to run tests and build commands inside my isolated environment so that it doesn't fail due to missing host dependencies.*

- **Acceptance Criteria**:
    - `oursky agent` provides an MCP endpoint.
    - The agent can use the `exec` tool to run `npm test` inside the container.
    - The agent follows the rules defined in `.oursky/agents/rules`.

## 3. Microservices & API Discovery
> *As a frontend developer, I want to automatically know the URL of my backend API running in the Oursky environment so that I don't have to manually update my `.env.local` every time.*

- **Acceptance Criteria**:
    - Oursky discovers the host-mapped port for the `api` service.
    - The environment variable `OURSKY_SERVICE_API_URL` is automatically injected into my frontend container.

## 4. Multi-Role Onboarding
> *As a team lead, I want to define different rules for 'Designer' and 'Developer' roles in my project so that the AI agent provides relevant suggestions based on who is using it.*

- **Acceptance Criteria**:
    - `.oursky/config.yaml` supports an `agent.role` field.
    - `oursky sync-agents` pulls role-specific rules from `.oursky/agents/rules/roles/`.

## 5. Teardown & Cleanup
> *As a developer, I want to completely remove all artifacts (containers, worktrees) related to a finished feature branch so that I don't waste disk space.*

- **Acceptance Criteria**:
    - `oursky kill session-id` destroys the container and the worktree.
    - No dangling Docker volumes or git worktrees remain.
