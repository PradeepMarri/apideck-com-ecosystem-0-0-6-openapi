package tools

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/ecosystem-api/mcp-server/config"
	"github.com/ecosystem-api/mcp-server/models"
	"github.com/mark3labs/mcp-go/mcp"
)

func CategorylistingsallHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		ecosystem_idVal, ok := args["ecosystem_id"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: ecosystem_id"), nil
		}
		ecosystem_id, ok := ecosystem_idVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: ecosystem_id"), nil
		}
		idVal, ok := args["id"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: id"), nil
		}
		id, ok := idVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: id"), nil
		}
		queryParams := make([]string, 0)
		if val, ok := args["cursor"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("cursor=%v", val))
		}
		if val, ok := args["limit"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("limit=%v", val))
		}
		queryString := ""
		if len(queryParams) > 0 {
			queryString = "?" + strings.Join(queryParams, "&")
		}
		url := fmt.Sprintf("%s/ecosystems/%s/categories/%s/listings%s", cfg.BaseURL, ecosystem_id, id, queryString)
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to create request", err), nil
		}
		// No authentication required for this endpoint
		req.Header.Set("Accept", "application/json")

		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Request failed", err), nil
		}
		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to read response body", err), nil
		}

		if resp.StatusCode >= 400 {
			return mcp.NewToolResultError(fmt.Sprintf("API error: %s", body)), nil
		}
		// Use properly typed response
		var result models.GetListingsResponse
		if err := json.Unmarshal(body, &result); err != nil {
			// Fallback to raw text if unmarshaling fails
			return mcp.NewToolResultText(string(body)), nil
		}

		prettyJSON, err := json.MarshalIndent(result, "", "  ")
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to format JSON", err), nil
		}

		return mcp.NewToolResultText(string(prettyJSON)), nil
	}
}

func CreateCategorylistingsallTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("get_ecosystems_ecosystem_id_categories_id_listings",
		mcp.WithDescription("List category listings"),
		mcp.WithString("ecosystem_id", mcp.Required(), mcp.Description("")),
		mcp.WithString("id", mcp.Required(), mcp.Description("ID of the record you are acting upon.")),
		mcp.WithString("cursor", mcp.Description("Cursor to start from. You can find cursors for next/previous pages in the meta.cursors property of the response.")),
		mcp.WithNumber("limit", mcp.Description("Number of records to return")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    CategorylistingsallHandler(cfg),
	}
}
