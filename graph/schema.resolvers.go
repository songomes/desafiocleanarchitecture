package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.47

import (
	"context"
	"fmt"
	"github.com/songomes/desafiocleanarchitecture/internal/usecase"

	"github.com/songomes/desafiocleanarchitecture/graph/model"
)

func (r *queryResolver) ListOrder(ctx context.Context) ([]*model.Order, error) {
	orders, err := r.GetAllOrdersUseCase.Execute()
	if err != nil {
		return nil, err
	}

	var result []*model.Order
	for _, order := range orders {
		result = append(result, &model.Order{
			ID:         order.ID,
			Price:      order.Price,
			Tax:        order.Tax,
			FinalPrice: order.FinalPrice,
		})
	}

	return result, nil
}

// CreateOrder is the resolver for the createOrder field.
func (r *mutationResolver) CreateOrder(ctx context.Context, input *model.OrderInput) (*model.Order, error) {
	dto := usecase.OrderInputDTO{
		ID:    input.ID,
		Price: float64(input.Price),
		Tax:   float64(input.Tax),
	}
	output, err := r.CreateOrderUseCase.Execute(dto)
	if err != nil {
		return nil, err
	}
	return &model.Order{
		ID:         output.ID,
		Price:      float64(output.Price),
		Tax:        float64(output.Tax),
		FinalPrice: float64(output.FinalPrice),
	}, nil
}

// Todos is the resolver for the todos field.
func (r *queryResolver) Todos(ctx context.Context) ([]*model.Todo, error) {
	panic(fmt.Errorf("not implemented: Todos - todos"))
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

// CreateTodo is the resolver for the createTodo field.
func (r *mutationResolver) CreateTodo(ctx context.Context, input model.NewTodo) (*model.Todo, error) {
	panic(fmt.Errorf("not implemented: CreateTodo - createTodo"))
}

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }