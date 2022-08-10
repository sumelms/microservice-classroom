package endpoints

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"

	domain2 "github.com/sumelms/microservice-classroom/internal/classroom/domain"
	"github.com/sumelms/microservice-classroom/pkg/validator"

	"github.com/go-kit/kit/endpoint"
	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/google/uuid"
)

type updateSubscriptionRequest struct {
	UUID        uuid.UUID `json:"uuid" validate:"required"`
	UserID      uuid.UUID `json:"user_id" validate:"required"`
	ClassroomID uuid.UUID `json:"classroom_id" validate:"required"`
	Role        string    `json:"role"`
}

type updateSubscriptionResponse struct {
	UUID        uuid.UUID `json:"uuid"`
	UserID      uuid.UUID `json:"user_id"`
	ClassroomID uuid.UUID `json:"classroom_id"`
	Role        string    `json:"role"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func NewUpdateSubscriptionHandler(s domain2.ServiceInterface, opts ...kithttp.ServerOption) *kithttp.Server {
	return kithttp.NewServer(
		makeUpdateSubscriptionEndpoint(s),
		decodeUpdateSubscriptionRequest,
		encodeUpdateSubscriptionResponse,
		opts...,
	)
}

func makeUpdateSubscriptionEndpoint(s domain2.ServiceInterface) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req, ok := request.(updateSubscriptionRequest)
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

		updated, err := s.UpdateSubscription(ctx, &sub)

		return updateSubscriptionResponse{
			UUID:        updated.UUID,
			UserID:      updated.UserID,
			ClassroomID: updated.ClassroomID,
			Role:        updated.Role,
			CreatedAt:   updated.CreatedAt,
			UpdatedAt:   updated.UpdatedAt,
		}, nil
	}
}

func decodeUpdateSubscriptionRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)
	id, ok := vars["uuid"]
	if !ok {
		return nil, fmt.Errorf("invalid argument")
	}

	var req updateSubscriptionRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, err
	}

	req.UUID = uuid.MustParse(id)

	return req, nil
}

func encodeUpdateSubscriptionResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	return kithttp.EncodeJSONResponse(ctx, w, response)
}
