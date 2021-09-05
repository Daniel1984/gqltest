package blocks

import (
	"context"
	"log"
	"net/http"
	"strconv"

	"github.com/gqltest/cmd/server/models"
	"github.com/gqltest/internals/gql"
	"github.com/gqltest/internals/mdw"
	"github.com/gqltest/internals/servant"
	"github.com/julienschmidt/httprouter"
	"github.com/machinebox/graphql"
)

type GetSwapsHandler struct {
	gqlCli gql.GqlManager
}

func NewGetSwapsHandler(gqlCli gql.GqlManager) httprouter.Handle {
	return GetSwapsHandler{gqlCli}.Do()
}

func (gsh GetSwapsHandler) getSwaps(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	defer r.Body.Close()
	height := p.ByName("height")

	blockHeight, err := strconv.Atoi(height)
	if err != nil {
		servant.RespondJSON(w, http.StatusPreconditionFailed, "invalid height param")
		return
	}

	gqlReq := graphql.NewRequest(`
		query ($height: Int!) {
			swaps(Block_height: $height) {
				id
				timestamp
				token0 {
					id
					symbol
					name
					volume
					volumeUSD
					txCount
					poolCount
				}
				token1 {
					id
					symbol
					name
					volume
					volumeUSD
					txCount
					poolCount
				}
				amount0
				amount1
				amountUSD
			}
		}
	`)
	gqlReq.Var("height", blockHeight)

	resp := struct {
		Swaps []*models.Swap `json:"swaps"`
	}{}

	if err := gsh.gqlCli.Run(context.TODO(), gqlReq, &resp); err != nil {
		servant.RespondJSON(w, http.StatusBadRequest, "please try again later")
	}

	if err := servant.RespondJSON(w, 200, resp); err != nil {
		log.Printf("getPools: failed responding: %s\n", err)
	}
}

func (gsh GetSwapsHandler) Do() httprouter.Handle {
	return mdw.Chain(gsh.getSwaps, mdw.GetCommonMdw()...)
}
