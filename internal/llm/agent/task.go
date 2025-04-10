package agent

import (
	"context"
	"errors"

	"github.com/kujtimiihoxha/termai/internal/app"
	"github.com/kujtimiihoxha/termai/internal/config"
	"github.com/kujtimiihoxha/termai/internal/llm/models"
	"github.com/kujtimiihoxha/termai/internal/llm/tools"
)

type taskAgent struct {
	*agent
}

func (c *taskAgent) Generate(ctx context.Context, sessionID string, content string) error {
	return c.generate(ctx, sessionID, content)
}

func NewTaskAgent(app *app.App) (Agent, error) {
	model, ok := models.SupportedModels[config.Get().Model.Coder]
	if !ok {
		return nil, errors.New("model not supported")
	}

	agentProvider, titleGenerator, err := getAgentProviders(app.Context, model)
	if err != nil {
		return nil, err
	}
	return &taskAgent{
		agent: &agent{
			App: app,
			tools: []tools.BaseTool{
				tools.NewGlobTool(),
				tools.NewGrepTool(),
				tools.NewLsTool(),
				tools.NewSourcegraphTool(),
				tools.NewViewTool(app.LSPClients),
			},
			model:          model,
			agent:          agentProvider,
			titleGenerator: titleGenerator,
		},
	}, nil
}
