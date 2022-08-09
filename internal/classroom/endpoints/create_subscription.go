package endpoints

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	domain2 "github.com/sumelms/microservice-classroom/internal/classroom/domain"
	"github.com/sumelms/microservice-classroom/pkg/validator"

	"github.com/go-kit/kit/endpoint"
	kithttp "github.com/go-kit/kit/transport/http"
)

type createSubscriptionRequest struct {
	UserID      string `json:"user_id" validate:"required"`
	ClassroomID string `json:"classroom_id" validate:"required"`
	Role        string `json:"role"`
}

type createSubscriptionResponse struct {
	UUID        string `json:"uuid"`
	UserID      string `json:"user_id"`
	ClassroomID string `json:"classroom_id"`
	Role        string `json:"role"`
}

func NewCreateSubscriptionHandler(s domain2.ServiceInterface, opts ...kithttp.ServerOption) *kithttp.Server {
	return kithttp.NewServer(
		makeCreateSubscriptionEndpoint(s),
		decodeCreateSubscriptionRequest,
		encodeCreateSubscriptionResponse,
		opts...,
	)
}

func makeCreateSubscriptionEndpoint(s domain2.ServiceInterface) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req, ok := request.(createSubscriptionRequest)
		if !ok {
			return nil, fmt.Errorf("invalid argument")
		}

		v := validator.NewValidator()
		if err := v.Validate(req); err != nil {
			return nil, err
		}

		var sub domain2.Subscription
		data, _ := json.Marshal(req)
		err := json.Unmarshal(data, &sub)
		if err != nil {
			return nil, err
		}

		created, err := s.CreateSubscription(ctx, &sub)
		if err != nil {
			return nil, err
		}

		return createSubscriptionResponse{
			UUID:        created.UUID,
			UserID:      created.UserID,
			ClassroomID: created.ClassroomID,
			Role:        created.Role,
		}, nil
	}
}

func decodeCreateSubscriptionRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var req createSubscriptionRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, err
	}
	return req, nil
}

func encodeCreateSubscriptionResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	return kithttp.EncodeJSONResponse(ctx, w, response)
}
