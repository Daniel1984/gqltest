package assets

import (
	"log"
	"net/http"

	"github.com/gqltest/internals/gql"
	"github.com/gqltest/internals/mdw"
	"github.com/julienschmidt/httprouter"
)

type GetVolumeHandler struct {
	gqlCli gql.GqlManager
}

func NewGetVolumeHandler(gqlCli gql.GqlManager) httprouter.Handle {
	return GetVolumeHandler{gqlCli}.Do()
}

func (gvh GetVolumeHandler) getVolume(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	defer r.Body.Close()
	assetId := p.ByName("id")
	log.Println("assetId: ", assetId)
	w.Write([]byte("TODO: implement total volume for given timerange and asset id: " + assetId))
}

func (gvh GetVolumeHandler) Do() httprouter.Handle {
	return mdw.Chain(gvh.getVolume, mdw.GetCommonMdw()...)
}
