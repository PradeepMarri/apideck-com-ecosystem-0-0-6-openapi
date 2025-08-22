package main

import (
	"github.com/ecosystem-api/mcp-server/config"
	"github.com/ecosystem-api/mcp-server/models"
	tools_listing "github.com/ecosystem-api/mcp-server/tools/listing"
	tools_product "github.com/ecosystem-api/mcp-server/tools/product"
	tools_category "github.com/ecosystem-api/mcp-server/tools/category"
	tools_collection "github.com/ecosystem-api/mcp-server/tools/collection"
	tools_ecosystem "github.com/ecosystem-api/mcp-server/tools/ecosystem"
)

func GetAll(cfg *config.APIConfig) []models.Tool {
	return []models.Tool{
		tools_listing.CreateListingsoneTool(cfg),
		tools_product.CreateProductlistingsallTool(cfg),
		tools_category.CreateCategoriesallTool(cfg),
		tools_category.CreateCategoriesoneTool(cfg),
		tools_category.CreateCategorylistingsallTool(cfg),
		tools_collection.CreateCollectionlistingsallTool(cfg),
		tools_product.CreateProductsoneTool(cfg),
		tools_ecosystem.CreateEcosystemsoneTool(cfg),
		tools_collection.CreateCollectionsallTool(cfg),
		tools_collection.CreateCollectionsoneTool(cfg),
		tools_listing.CreateListingsallTool(cfg),
		tools_product.CreateProductsallTool(cfg),
	}
}
