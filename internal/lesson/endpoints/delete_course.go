package endpoints

import (
	"context"
	"fmt"
	"net/http"

	"github.com/go-kit/kit/endpoint"
	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	"github.com/sumelms/microservice-classroom/internal/lesson/domain"
)

type deleteLessonRequest struct {
	UUID string `json:"uuid" validate:"required"`
}

func NewDeleteLessonHandler(s domain.ServiceInterface, opts ...kithttp.ServerOption) *kithttp.Server {
	return kithttp.NewServer(
		makeDeleteLessonEndpoint(s),
		decodeDeleteLessonRequest,
		encodeDeleteLessonResponse,
		opts...,
	)
}

func makeDeleteLessonEndpoint(s domain.ServiceInterface) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req, ok := request.(deleteLessonRequest)
		if !ok {
			return nil, fmt.Errorf("invalid argument")
		}

		err := s.DeleteLesson(ctx, req.UUID)
		if err != nil {
			return nil, err
		}

		return nil, nil
	}
}

func decodeDeleteLessonRequest(_ context.Context, r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)
	id, ok := vars["uuid"]
	if !ok {
		return nil, fmt.Errorf("invalid argument")
	}

	return deleteLessonRequest{UUID: id}, nil
}

func encodeDeleteLessonResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	return kithttp.EncodeJSONResponse(ctx, w, response)
}
