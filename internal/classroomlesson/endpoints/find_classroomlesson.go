package endpoints

import (
	"context"
	"fmt"
	"github.com/sumelms/microservice-classroom/internal/classroomlesson/domain"
	"net/http"
	"time"

	"github.com/go-kit/kit/endpoint"
	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
)

type findClassroomLessonRequest struct {
	UUID string `json:"uuid"`
}

type findClassroomLessonResponse struct {
	UUID        string     `json:"uuid"`
	ClassroomID string     `json:"classroom_id"`
	LessonID    string     `json:"lesson_id"`
	StartsAt    time.Time  `json:"starts_at"`
	EndsAt      *time.Time `json:"ends_at"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
}

func NewFindClassroomLessonHandler(s domain.ServiceInterface, opts ...kithttp.ServerOption) *kithttp.Server {
	return kithttp.NewServer(
		makeFindClassroomLessonEndpoint(s),
		decodeFindClassroomLessonRequest,
		encodeFindClassroomLessonResponse,
		opts...,
	)
}

func makeFindClassroomLessonEndpoint(s domain.ServiceInterface) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req, ok := request.(findClassroomLessonRequest)
		if !ok {
			return nil, fmt.Errorf("invalid argument")
		}

		c, err := s.FindClassroomLesson(ctx, req.UUID)
		if err != nil {
			return nil, err
		}

		return &findClassroomLessonResponse{
			UUID:        c.UUID,
			ClassroomID: c.ClassroomID,
			LessonID:    c.LessonID,
			StartsAt:    c.StartsAt,
			EndsAt:      c.EndsAt,
		}, nil
	}
}

func decodeFindClassroomLessonRequest(_ context.Context, r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)
	id, ok := vars["uuid"]
	if !ok {
		return nil, fmt.Errorf("invalid argument")
	}

	return findClassroomLessonRequest{UUID: id}, nil
}

func encodeFindClassroomLessonResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	return kithttp.EncodeJSONResponse(ctx, w, response)
}
