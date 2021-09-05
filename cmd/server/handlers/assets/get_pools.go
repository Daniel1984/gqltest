package assets

import (
	"context"
	"log"
	"net/http"

	"github.com/gqltest/cmd/server/models"
	"github.com/gqltest/internals/gql"
	"github.com/gqltest/internals/mdw"
	"github.com/gqltest/internals/servant"
	"github.com/julienschmidt/httprouter"
	"github.com/machinebox/graphql"
)

type GetPoolsHandler struct {
	gqlCli gql.GqlManager
}

func NewGetPoolsHandler(gqlCli gql.GqlManager) httprouter.Handle {
	return GetPoolsHandler{gqlCli}.Do()
}

func (gph GetPoolsHandler) getPools(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	defer r.Body.Close()
	assetId := p.ByName("id")

	gqlReq := graphql.NewRequest(`
		query ($id: String!) {
			token(id: $id) {
				id
				symbol
				name
				volume
				volumeUSD
				txCount
				poolCount
				whitelistPools {
					id
				}
			}
		}
	`)
	gqlReq.Var("id", assetId)

	resp := struct {
		Token models.Token `json:"token"`
	}{}

	if err := gph.gqlCli.Run(context.TODO(), gqlReq, &resp); err != nil {
		servant.RespondJSON(w, http.StatusBadRequest, "please try again later")
	}

	if err := servant.RespondJSON(w, 200, resp.Token.WhitelistPools); err != nil {
		log.Printf("getPools: failed responding: %s\n", err)
	}
}

func (gph GetPoolsHandler) Do() httprouter.Handle {
	return mdw.Chain(gph.getPools, mdw.GetCommonMdw()...)
}
