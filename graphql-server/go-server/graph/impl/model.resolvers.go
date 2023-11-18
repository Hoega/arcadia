package impl

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.40

import (
	"context"
	"strings"

	"github.com/kubeagi/arcadia/graphql-server/go-server/graph/generated"
	"github.com/kubeagi/arcadia/graphql-server/go-server/pkg/auth"
	"github.com/kubeagi/arcadia/graphql-server/go-server/pkg/client"
	md "github.com/kubeagi/arcadia/graphql-server/go-server/pkg/model"
)

// CreateModel is the resolver for the createModel field.
func (r *modelMutationResolver) CreateModel(ctx context.Context, obj *generated.ModelMutation, input generated.CreateModelInput) (*generated.Model, error) {
	token := auth.ForOIDCToken(ctx)
	c, err := client.GetClient(token)
	if err != nil {
		return nil, err
	}

	displayname, description, filed, modeltype := "", "", "", ""

	if input.DisplayName != "" {
		displayname = input.DisplayName
	}
	if input.Description != nil {
		description = *input.Description
	}
	if input.Field != "" {
		filed = input.Field
	}
	if input.Modeltype != "" {
		modeltype = input.Modeltype
	}
	return md.CreateModel(ctx, c, input.Name, input.Namespace, displayname, description, filed, modeltype)
}

// UpdateModel is the resolver for the updateModel field.
func (r *modelMutationResolver) UpdateModel(ctx context.Context, obj *generated.ModelMutation, input *generated.UpdateModelInput) (*generated.Model, error) {
	token := auth.ForOIDCToken(ctx)
	c, err := client.GetClient(token)
	if err != nil {
		return nil, err
	}
	name, displayname := "", ""
	if input.DisplayName != "" {
		displayname = input.DisplayName
	}
	if input.Name != "" {
		name = input.Name

	}
	return md.UpdateModel(ctx, c, name, input.Namespace, displayname)
}

// DeleteModel is the resolver for the deleteModel field.
func (r *modelMutationResolver) DeleteModel(ctx context.Context, obj *generated.ModelMutation, input *generated.DeleteModelInput) (*string, error) {
	token := auth.ForOIDCToken(ctx)
	c, err := client.GetClient(token)
	if err != nil {
		return nil, err
	}
	name := ""
	labelSelector, fieldSelector := "", ""
	if input.Name != nil {
		name = *input.Name
	}
	if input.FieldSelector != nil {
		fieldSelector = *input.FieldSelector
	}
	if input.LabelSelector != nil {
		labelSelector = *input.LabelSelector
	}
	return md.DeleteModel(ctx, c, name, input.Namespace, labelSelector, fieldSelector)
}

// GetModel is the resolver for the getModel field.
func (r *modelQueryResolver) GetModel(ctx context.Context, obj *generated.ModelQuery, name string, namespace string) (*generated.Model, error) {
	token := auth.ForOIDCToken(ctx)
	c, err := client.GetClient(token)
	if err != nil {
		return nil, err
	}
	return md.ReadModel(ctx, c, name, namespace)
}

// ListModels is the resolver for the listModels field.
func (r *modelQueryResolver) ListModels(ctx context.Context, obj *generated.ModelQuery, input generated.ListModelInput) (*generated.PaginatedResult, error) {
	token := auth.ForOIDCToken(ctx)
	c, err := client.GetClient(token)
	if err != nil {
		return nil, err
	}
	name, displayName, labelSelector, fieldSelector := "", "", "", ""
	page, pageSize := 1, 10
	if input.Name != nil {
		name = *input.Name
	}
	if input.DisplayName != nil {
		displayName = *input.DisplayName
	}
	if input.FieldSelector != nil {
		fieldSelector = *input.FieldSelector
	}
	if input.LabelSelector != nil {
		labelSelector = *input.LabelSelector
	}
	if input.Page != nil && *input.Page > 0 {
		page = *input.Page
	}
	if input.PageSize != nil && *input.PageSize > 0 {
		pageSize = *input.PageSize
	}
	result, err := md.ListModels(ctx, c, input.Namespace, labelSelector, fieldSelector)
	if err != nil {
		return nil, err
	}
	var filteredResult []generated.PageNode
	for idx, u := range result {
		if (name == "" || strings.Contains(u.Name, name)) && (displayName == "" || strings.Contains(u.DisplayName, displayName)) {
			filteredResult = append(filteredResult, result[idx])
		}
	}

	totalCount := len(filteredResult)
	end := page * pageSize
	if end > totalCount {
		end = totalCount
	}
	return &generated.PaginatedResult{
		TotalCount:  totalCount,
		HasNextPage: end < totalCount,
		Nodes:       filteredResult[(page-1)*pageSize : end],
	}, nil
}

// Model is the resolver for the Model field.
func (r *mutationResolver) Model(ctx context.Context) (*generated.ModelMutation, error) {
	return &generated.ModelMutation{}, nil
}

// Model is the resolver for the Model field.
func (r *queryResolver) Model(ctx context.Context) (*generated.ModelQuery, error) {
	return &generated.ModelQuery{}, nil
}

// ModelMutation returns generated.ModelMutationResolver implementation.
func (r *Resolver) ModelMutation() generated.ModelMutationResolver { return &modelMutationResolver{r} }

// ModelQuery returns generated.ModelQueryResolver implementation.
func (r *Resolver) ModelQuery() generated.ModelQueryResolver { return &modelQueryResolver{r} }

type modelMutationResolver struct{ *Resolver }
type modelQueryResolver struct{ *Resolver }