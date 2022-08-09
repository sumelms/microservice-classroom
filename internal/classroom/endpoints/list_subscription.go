package endpoints

import (
	"context"
	"fmt"
	"net/http"

	"github.com/go-kit/kit/endpoint"
	kithttp "github.com/go-kit/kit/transport/http"

	"github.com/sumelms/microservice-classroom/internal/classroom/domain"
)

type listSubscriptionRequest struct {
	ClassroomID string `json:"classroom_id"`
	UserID      string `json:"user_id"`
	Role        string `json:"role"`
}

type listSubscriptionResponse struct {
	Subscriptions []findSubscriptionResponse `json:"subscriptions"`
}

func NewListSubscriptionHandler(s domain.ServiceInterface, opts ...kithttp.ServerOption) *kithttp.Server {
	return kithttp.NewServer(
		makeListSubscriptionEndpoint(s),
		decodeListSubscriptionRequest,
		encodeListSubscriptionResponse,
		opts...,
	)
}

func makeListSubscriptionEndpoint(s domain.ServiceInterface) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req, ok := request.(listSubscriptionRequest)
		if !ok {
			return nil, fmt.Errorf("invalid argument")
		}

		filters := make(map[string]interface{})
		if len(req.ClassroomID) > 0 {
			filters["classroom_id"] = req.ClassroomID
		}
		if len(req.UserID) > 0 {
			filters["user_id"] = req.UserID
		}
		if len(req.Role) > 0 {
			filters["role"] = req.Role
		}

		subscriptions, err := s.ListSubscription(ctx, filters)
		if err != nil {
			return nil, err
		}

		var list []findSubscriptionResponse
		for _, sub := range subscriptions {
			list = append(list, findSubscriptionResponse{
				UUID:        sub.UUID,
				UserID:      sub.UserID,
				ClassroomID: sub.ClassroomID,
				Role:        sub.Role,
				CreatedAt:   sub.CreatedAt,
				UpdatedAt:   sub.UpdatedAt,
			})
		}

		return &listSubscriptionResponse{Subscriptions: list}, nil
	}
}

func decodeListSubscriptionRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	classroomID := r.FormValue("classroom_id")
	userID := r.FormValue("user_id")
	role := r.FormValue("role")
	return listSubscriptionRequest{
		ClassroomID: classroomID,
		UserID:      userID,
		Role:        role,
	}, nil
}

func encodeListSubscriptionResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	return kithttp.EncodeJSONResponse(ctx, w, response)
}
