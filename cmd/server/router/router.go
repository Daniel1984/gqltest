package router

import (
	"github.com/gqltest/cmd/server/handlers/assets"
	"github.com/gqltest/cmd/server/handlers/blocks"
	"github.com/gqltest/internals/gql"
	"github.com/julienschmidt/httprouter"
)

func New(gqlCli gql.GqlManager) *httprouter.Router {
	mux := httprouter.New()
	mux.GET("/assets/:id/pools", assets.NewGetPoolsHandler(gqlCli))
	mux.GET("/assets/:id/volume", assets.NewGetVolumeHandler(gqlCli))
	mux.GET("/blocks/:height/swaps", blocks.NewGetSwapsHandler(gqlCli))
	mux.GET("/blocks/:height/assets", blocks.NewGetAssetsHandler(gqlCli))
	return mux
}
