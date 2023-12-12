package impl

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.40

import (
	"context"

	"github.com/kubeagi/arcadia/graphql-server/go-server/graph/generated"
	md "github.com/kubeagi/arcadia/graphql-server/go-server/pkg/model"
)

// Files is the resolver for the files field.
func (r *modelResolver) Files(ctx context.Context, obj *generated.Model, input *generated.FileFilter) (*generated.PaginatedResult, error) {
	c, err := getClientFromCtx(ctx)
	if err != nil {
		return nil, err
	}
	return md.ModelFiles(ctx, c, obj.Name, obj.Namespace, input)
}

// CreateModel is the resolver for the createModel field.
func (r *modelMutationResolver) CreateModel(ctx context.Context, obj *generated.ModelMutation, input generated.CreateModelInput) (*generated.Model, error) {
	c, err := getClientFromCtx(ctx)
	if err != nil {
		return nil, err
	}

	return md.CreateModel(ctx, c, input)
}

// UpdateModel is the resolver for the updateModel field.
func (r *modelMutationResolver) UpdateModel(ctx context.Context, obj *generated.ModelMutation, input *generated.UpdateModelInput) (*generated.Model, error) {
	c, err := getClientFromCtx(ctx)
	if err != nil {
		return nil, err
	}

	return md.UpdateModel(ctx, c, input)
}

// DeleteModels is the resolver for the deleteModels field.
func (r *modelMutationResolver) DeleteModels(ctx context.Context, obj *generated.ModelMutation, input *generated.DeleteCommonInput) (*string, error) {
	c, err := getClientFromCtx(ctx)
	if err != nil {
		return nil, err
	}

	return md.DeleteModels(ctx, c, input)
}

// GetModel is the resolver for the getModel field.
func (r *modelQueryResolver) GetModel(ctx context.Context, obj *generated.ModelQuery, name string, namespace string) (*generated.Model, error) {
	c, err := getClientFromCtx(ctx)
	if err != nil {
		return nil, err
	}
	return md.ReadModel(ctx, c, name, namespace)
}

// ListModels is the resolver for the listModels field.
func (r *modelQueryResolver) ListModels(ctx context.Context, obj *generated.ModelQuery, input generated.ListModelInput) (*generated.PaginatedResult, error) {
	c, err := getClientFromCtx(ctx)
	if err != nil {
		return nil, err
	}

	return md.ListModels(ctx, c, input)
}

// Model is the resolver for the Model field.
func (r *mutationResolver) Model(ctx context.Context) (*generated.ModelMutation, error) {
	return &generated.ModelMutation{}, nil
}

// Model is the resolver for the Model field.
func (r *queryResolver) Model(ctx context.Context) (*generated.ModelQuery, error) {
	return &generated.ModelQuery{}, nil
}

// Model returns generated.ModelResolver implementation.
func (r *Resolver) Model() generated.ModelResolver { return &modelResolver{r} }

// ModelMutation returns generated.ModelMutationResolver implementation.
func (r *Resolver) ModelMutation() generated.ModelMutationResolver { return &modelMutationResolver{r} }

// ModelQuery returns generated.ModelQueryResolver implementation.
func (r *Resolver) ModelQuery() generated.ModelQueryResolver { return &modelQueryResolver{r} }

type modelResolver struct{ *Resolver }
type modelMutationResolver struct{ *Resolver }
type modelQueryResolver struct{ *Resolver }
