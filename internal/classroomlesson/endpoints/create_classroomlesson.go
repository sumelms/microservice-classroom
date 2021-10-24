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
	"github.com/sumelms/microservice-classroom/pkg/validator"
)

type createClassroomLessonRequest struct {
	ClassroomID string     `json:"classroom_id" validate:"required"`
	LessonID    string     `json:"lesson_id" validate:"required"`
	StartsAt    time.Time  `json:"starts_at" validate:"required"`
	EndsAt      *time.Time `json:"ends_at"`
}

type createClassroomLessonResponse struct {
	UUID        string     `json:"uuid"`
	ClassroomID string     `json:"classroom_id"`
	LessonID    string     `json:"lesson_id"`
	StartsAt    time.Time  `json:"starts_at"`
	EndsAt      *time.Time `json:"ends_at"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
}

func NewCreateClassroomLessonHandler(s domain.ServiceInterface, opts ...kithttp.ServerOption) *kithttp.Server {
	return kithttp.NewServer(
		makeCreateClassroomLessonEndpoint(s),
		decodeCreateClassroomLessonRequest,
		encodeCreateClassroomLessonResponse,
		opts...,
	)
}

func makeCreateClassroomLessonEndpoint(s domain.ServiceInterface) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req, ok := request.(createClassroomLessonRequest)
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
		c.ClassroomID = req.ClassroomID
		c.LessonID = req.LessonID

		created, err := s.CreateClassroomLesson(ctx, &c)
		if err != nil {
			return nil, err
		}

		return createClassroomLessonResponse{
			UUID:        created.UUID,
			ClassroomID: created.ClassroomID,
			LessonID:    created.LessonID,
			StartsAt:    created.StartsAt,
			EndsAt:      created.EndsAt,
			CreatedAt:   created.CreatedAt,
			UpdatedAt:   created.UpdatedAt,
		}, err
	}
}

func decodeCreateClassroomLessonRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req createClassroomLessonRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, err
	}
	return req, nil
}

func encodeCreateClassroomLessonResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	return kithttp.EncodeJSONResponse(ctx, w, response)
}
