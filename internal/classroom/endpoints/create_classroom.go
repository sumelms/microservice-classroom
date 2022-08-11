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

	"github.com/sumelms/microservice-classroom/pkg/validator"
)

type createClassroomRequest struct {
	Code         string     `json:"code" validate:"required,max=15"`
	Name         string     `json:"name" validate:"required,max=100"`
	Description  string     `json:"description" validate:"max=255"`
	Format       string     `json:"format" validate:"classroom_format"`
	CabSubscribe bool       `json:"cab_subscribe"`
	SubjectID    *uuid.UUID `json:"subject_id"`
	CourseID     uuid.UUID  `json:"course_id" validate:"required"`
	StartsAt     time.Time  `json:"starts_at" validate:"required"`
	EndsAt       *time.Time `json:"ends_at"`
}

type createClassroomResponse struct {
	UUID         uuid.UUID  `json:"uuid"`
	Code         string     `json:"code"`
	Name         string     `json:"name"`
	Description  string     `json:"description"`
	Format       string     `json:"format"`
	CanSubscribe bool       `json:"can_subscribe"`
	SubjectID    *uuid.UUID `json:"subject_id,omitempty"`
	CourseID     uuid.UUID  `json:"course_id"`
	StartsAt     time.Time  `json:"starts_at" validate:"required"`
	EndsAt       *time.Time `json:"ends_at,omitempty"`
	CreatedAt    time.Time  `json:"created_at"`
	UpdatedAt    time.Time  `json:"updated_at"`
}

func NewCreateClassroomHandler(s domain.ServiceInterface, opts ...kithttp.ServerOption) *kithttp.Server {
	return kithttp.NewServer(
		makeCreateClassroomEndpoint(s),
		decodeCreateClassroomRequest,
		encodeCreateClassroomResponse,
		opts...,
	)
}

func makeCreateClassroomEndpoint(s domain.ServiceInterface) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req, ok := request.(createClassroomRequest)
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

		created, err := s.CreateClassroom(ctx, &c)
		if err != nil {
			return nil, err
		}

		return createClassroomResponse{
			UUID:         created.UUID,
			Code:         created.Code,
			Name:         created.Name,
			Description:  created.Description,
			Format:       created.Format,
			CanSubscribe: created.CanSubscribe,
			SubjectID:    created.SubjectID,
			CourseID:     created.CourseID,
			StartsAt:     created.StartsAt,
			EndsAt:       created.EndsAt,
			CreatedAt:    created.CreatedAt,
			UpdatedAt:    created.UpdatedAt,
		}, err
	}
}

func decodeCreateClassroomRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req createClassroomRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, err
	}
	return req, nil
}

func encodeCreateClassroomResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	return kithttp.EncodeJSONResponse(ctx, w, response)
}
