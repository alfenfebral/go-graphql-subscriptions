package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"graphql-subscriptions/graph/generated"

	"github.com/nats-io/nats.go"
)

func (r *queryResolver) Payload(ctx context.Context) (*string, error) {
	value := "hello world"
	return &value, nil
}

func (r *subscriptionResolver) Payload(ctx context.Context) (<-chan *string, error) {
	ch := make(chan *string)

	r.Nats.Subscribe("payload-subject", func(msg *nats.Msg) {
		payload := string(msg.Data)
		ch <- &payload
	})

	return ch, nil
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

// Subscription returns generated.SubscriptionResolver implementation.
func (r *Resolver) Subscription() generated.SubscriptionResolver { return &subscriptionResolver{r} }

type queryResolver struct{ *Resolver }
type subscriptionResolver struct{ *Resolver }

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//  - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//    it when you're done.
//  - You have helper methods in this file. Move them out to keep these resolver files clean.
