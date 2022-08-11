package endpoints

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/sumelms/microservice-classroom/internal/classroom/domain"

	"github.com/go-kit/kit/endpoint"
	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/google/uuid"
	"github.com/gorilla/mux"

	"github.com/sumelms/microservice-classroom/pkg/validator"
)

type updateClassroomRequest struct {
	UUID         uuid.UUID  `json:"uuid" validate:"required"`
	Code         string     `json:"code" validate:"required,max=15"`
	Name         string     `json:"name" validate:"required,max=100"`
	Description  string     `json:"description" validate:"required,max=255"`
	Format       string     `json:"format" validate:"classroom_format"`
	CanSubscribe bool       `json:"can_subscribe" validate:"required"`
	SubjectID    *uuid.UUID `json:"subject_id" validate:"required"`
	CourseID     uuid.UUID  `json:"course_id" validate:"required"`
	StartsAt     time.Time  `json:"starts_at" validate:"required"`
	EndsAt       *time.Time `json:"ends_at"`
}

type updateClassroomResponse struct {
	UUID         uuid.UUID  `json:"uuid"`
	Code         string     `json:"code"`
	Name         string     `json:"name"`
	Description  string     `json:"description"`
	Format       string     `json:"format"`
	CanSubscribe bool       `json:"can_subscribe"`
	SubjectID    *uuid.UUID `json:"subject_id,omitempty"`
	CourseID     uuid.UUID  `json:"course_id"`
	StartsAt     time.Time  `json:"starts_at"`
	EndsAt       *time.Time `json:"ends_at,omitempty"`
	CreatedAt    time.Time  `json:"created_at"`
	UpdatedAt    time.Time  `json:"updated_at"`
}

func NewUpdateClassroomHandler(s domain.ServiceInterface, opts ...kithttp.ServerOption) *kithttp.Server {
	return kithttp.NewServer(
		makeUpdateClassroomEndpoint(s),
		decodeUpdateClassroomRequest,
		encodeUpdateClassroomResponse,
		opts...,
	)
}

func makeUpdateClassroomEndpoint(s domain.ServiceInterface) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req, ok := request.(updateClassroomRequest)
		if !ok {
			return nil, fmt.Errorf("invalid argument")
		}

		v := validator.NewValidator()
		if err := v.Validate(req); err != nil {
			return nil, err
		}

		c := domain.Classroom{}
		data, _ := json.Marshal(req)
		err := json.Unmarshal(data, &c)
		if err != nil {
			return nil, err
		}
		if req.SubjectID != nil {
			c.SubjectID = req.SubjectID
		}
		c.CourseID = req.CourseID

		updated, err := s.UpdateClassroom(ctx, &c)
		if err != nil {
			return nil, err
		}

		return updateClassroomResponse{
			UUID:         updated.UUID,
			Code:         updated.Code,
			Name:         updated.Name,
			Description:  updated.Description,
			Format:       updated.Format,
			CanSubscribe: updated.CanSubscribe,
			SubjectID:    updated.SubjectID,
			CourseID:     updated.CourseID,
			StartsAt:     c.StartsAt,
			EndsAt:       c.EndsAt,
		}, nil
	}
}

func decodeUpdateClassroomRequest(_ context.Context, r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)
	id, ok := vars["uuid"]
	if !ok {
		return nil, fmt.Errorf("invalid argument")
	}

	var req updateClassroomRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, err
	}

	req.UUID = uuid.MustParse(id)

	return req, nil
}

func encodeUpdateClassroomResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	return kithttp.EncodeJSONResponse(ctx, w, response)
}
