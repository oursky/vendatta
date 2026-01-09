package agent

import (
	"context"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
	"github.com/vibegear/oursky/pkg/provider"
)

type AgentServer struct {
	SessionID string
	Provider  provider.Provider
}

func NewAgentServer(sessionID string, p provider.Provider) *AgentServer {
	return &AgentServer{SessionID: sessionID, Provider: p}
}

func (s *AgentServer) Serve() error {
	mcpServer := server.NewMCPServer("oursky-agent", "1.0.0")

	mcpServer.AddTool(mcp.NewTool("exec",
		mcp.WithDescription("Execute a command in the dev environment"),
		mcp.WithString("cmd", mcp.Description("The command to execute"), mcp.Required()),
	), s.handleExec)

	return server.ServeStdio(mcpServer)
}

func (s *AgentServer) handleExec(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	cmd, err := request.RequireString("cmd")
	if err != nil {
		return nil, err
	}

	err = s.Provider.Exec(ctx, s.SessionID, provider.ExecOptions{
		Cmd:    []string{"/bin/bash", "-c", cmd},
		Stdout: true,
		Stderr: true,
	})
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return mcp.NewToolResultText("Command executed successfully"), nil
}
