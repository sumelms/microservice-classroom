package endpoints

import (
	"context"
	"fmt"
	"net/http"

	"github.com/sumelms/microservice-classroom/internal/classroom/domain"

	"github.com/go-kit/kit/endpoint"
	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

type deleteClassroomRequest struct {
	UUID uuid.UUID `json:"uuid" validate:"required"`
}

func NewDeleteClassroomHandler(s domain.ServiceInterface, opts ...kithttp.ServerOption) *kithttp.Server {
	return kithttp.NewServer(
		makeDeleteClassroomEndpoint(s),
		decodeDeleteClassroomRequest,
		encodeDeleteClassroomResponse,
		opts...,
	)
}

func makeDeleteClassroomEndpoint(s domain.ServiceInterface) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req, ok := request.(deleteClassroomRequest)
		if !ok {
			return nil, fmt.Errorf("invalid argument")
		}

		err := s.DeleteClassroom(ctx, req.UUID)
		if err != nil {
			return nil, err
		}

		return nil, nil
	}
}

func decodeDeleteClassroomRequest(_ context.Context, r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)
	id, ok := vars["uuid"]
	if !ok {
		return nil, fmt.Errorf("invalid argument")
	}

	uid := uuid.MustParse(id)

	return deleteClassroomRequest{UUID: uid}, nil
}

func encodeDeleteClassroomResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	return kithttp.EncodeJSONResponse(ctx, w, response)
}
