package endpoints

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/sumelms/microservice-classroom/internal/classroomlesson/domain"
	"net/http"
	"time"

	"github.com/go-kit/kit/endpoint"
	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	"github.com/sumelms/microservice-classroom/pkg/validator"
)

type updateClassroomLessonRequest struct {
	UUID        string `json:"uuid" validate:"required"`
	ClassroomID string `json:"classroom_id" validate:"required"`
	LessonID    string `json:"lesson_id" validate:"required"`
}

type updateClassroomLessonResponse struct {
	UUID        string     `json:"uuid"`
	ClassroomID string     `json:"classroom_id"`
	LessonID    string     `json:"lesson_id"`
	StartsAt    time.Time  `json:"starts_at"`
	EndsAt      *time.Time `json:"ends_at"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
}

func NewUpdateClassroomLessonHandler(s domain.ServiceInterface, opts ...kithttp.ServerOption) *kithttp.Server {
	return kithttp.NewServer(
		makeUpdateClassroomLessonEndpoint(s),
		decodeUpdateClassroomLessonRequest,
		encodeUpdateClassroomLessonResponse,
		opts...,
	)
}

func makeUpdateClassroomLessonEndpoint(s domain.ServiceInterface) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req, ok := request.(updateClassroomLessonRequest)
		if !ok {
			return nil, fmt.Errorf("invalid argument")
		}

		v := validator.NewValidator()
		if err := v.Validate(req); err != nil {
			return nil, err
		}

		c := domain.ClassroomLesson{}
		data, _ := json.Marshal(req)
		err := json.Unmarshal(data, &c)
		if err != nil {
			return nil, err
		}

		updated, err := s.UpdateClassroomLesson(ctx, &c)
		if err != nil {
			return nil, err
		}

		return updateClassroomLessonResponse{
			UUID:        updated.UUID,
			ClassroomID: updated.ClassroomID,
			LessonID:    updated.LessonID,
			StartsAt:    updated.StartsAt,
			EndsAt:      updated.EndsAt,
		}, nil
	}
}

func decodeUpdateClassroomLessonRequest(_ context.Context, r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)
	id, ok := vars["uuid"]
	if !ok {
		return nil, fmt.Errorf("invalid argument")
	}

	var req updateClassroomLessonRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, err
	}

	req.UUID = id

	return req, nil
}

func encodeUpdateClassroomLessonResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	return kithttp.EncodeJSONResponse(ctx, w, response)
}
