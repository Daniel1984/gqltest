package gql

import (
	"context"

	"github.com/machinebox/graphql"
)

// GqlManager interface for graphql client for easy mocking
// in handlers when adding test coverage
type GqlManager interface {
	Run(context.Context, *graphql.Request, interface{}) error
}
