package blocks

import (
	"net/http"

	"github.com/gqltest/internals/gql"
	"github.com/gqltest/internals/mdw"
	"github.com/julienschmidt/httprouter"
)

type GetAssetsHandler struct {
	gqlCli gql.GqlManager
}

func NewGetAssetsHandler(gqlCli gql.GqlManager) httprouter.Handle {
	return GetAssetsHandler{gqlCli}.Do()
}

func (gah GetAssetsHandler) getAssets(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	defer r.Body.Close()
	height := p.ByName("height")
	w.Write([]byte("TODO: implement get assets from block height: " + height))
}

func (gah GetAssetsHandler) Do() httprouter.Handle {
	return mdw.Chain(gah.getAssets, mdw.GetCommonMdw()...)
}
