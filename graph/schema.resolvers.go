package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/luisgarciaalanis/directiveBug/graph/generated"
	"github.com/luisgarciaalanis/directiveBug/graph/model"
)

func (r *queryResolver) Test(ctx context.Context) (*model.Test, error) {
	fmt.Println("Test resolver called!")
	return &model.Test{
		Text: "Hello",
	}, nil
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
