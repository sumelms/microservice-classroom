package endpoints

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/sumelms/microservice-classroom/internal/classroom/domain"

	"github.com/go-kit/kit/endpoint"
	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

type findClassroomRequest struct {
	UUID uuid.UUID `json:"uuid"`
}

type findClassroomResponse struct {
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

func NewFindClassroomHandler(s domain.ServiceInterface, opts ...kithttp.ServerOption) *kithttp.Server {
	return kithttp.NewServer(
		makeFindClassroomEndpoint(s),
		decodeFindClassroomRequest,
		encodeFindClassroomResponse,
		opts...,
	)
}

func makeFindClassroomEndpoint(s domain.ServiceInterface) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req, ok := request.(findClassroomRequest)
		if !ok {
			return nil, fmt.Errorf("invalid argument")
		}

		c, err := s.Classroom(ctx, req.UUID)
		if err != nil {
			return nil, err
		}

		return &findClassroomResponse{
			UUID:         c.UUID,
			Name:         c.Name,
			Description:  c.Description,
			Format:       c.Format,
			CanSubscribe: c.CanSubscribe,
			SubjectID:    c.SubjectID,
			CourseID:     c.CourseID,
			StartsAt:     c.StartsAt,
			EndsAt:       c.EndsAt,
			CreatedAt:    c.CreatedAt,
			UpdatedAt:    c.UpdatedAt,
		}, nil
	}
}

func decodeFindClassroomRequest(_ context.Context, r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)
	id, ok := vars["uuid"]
	if !ok {
		return nil, fmt.Errorf("invalid argument")
	}

	uid := uuid.MustParse(id)

	return findClassroomRequest{UUID: uid}, nil
}

func encodeFindClassroomResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	return kithttp.EncodeJSONResponse(ctx, w, response)
}
