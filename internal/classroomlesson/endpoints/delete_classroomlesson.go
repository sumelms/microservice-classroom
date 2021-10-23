package endpoints

import (
	"context"
	"fmt"
	"github.com/sumelms/microservice-classroom/internal/classroomlesson/domain"
	"net/http"

	"github.com/go-kit/kit/endpoint"
	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
)

type deleteClassroomLessonRequest struct {
	UUID string `json:"uuid" validate:"required"`
}

func NewDeleteClassroomLessonHandler(s domain.ServiceInterface, opts ...kithttp.ServerOption) *kithttp.Server {
	return kithttp.NewServer(
		makeDeleteClassroomLessonEndpoint(s),
		decodeDeleteClassroomLessonRequest,
		encodeDeleteClassroomLessonResponse,
		opts...,
	)
}

func makeDeleteClassroomLessonEndpoint(s domain.ServiceInterface) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req, ok := request.(deleteClassroomLessonRequest)
		if !ok {
			return nil, fmt.Errorf("invalid argument")
		}

		err := s.DeleteClassroomLesson(ctx, req.UUID)
		if err != nil {
			return nil, err
		}

		return nil, nil
	}
}

func decodeDeleteClassroomLessonRequest(_ context.Context, r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)
	id, ok := vars["uuid"]
	if !ok {
		return nil, fmt.Errorf("invalid argument")
	}

	return deleteClassroomLessonRequest{UUID: id}, nil
}

func encodeDeleteClassroomLessonResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	return kithttp.EncodeJSONResponse(ctx, w, response)
}
