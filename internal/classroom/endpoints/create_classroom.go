package endpoints

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/sumelms/microservice-classroom/internal/classroom/domain"
	"net/http"
	"time"

	"github.com/go-kit/kit/endpoint"
	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/sumelms/microservice-classroom/pkg/validator"
)

type createClassroomRequest struct {
	Title       string `json:"title" validate:"required,max=100"`
	Description string `json:"description" validate:"required,max=255"`
	SubjectID   string `json:"subject_id" validate:"required"`
	CourseID    string `json:"course_id" validate:"required"`
}

type createClassroomResponse struct {
	UUID        string    `json:"uuid"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	SubjectID   string    `json:"subject_id"`
	CourseID    string    `json:"course_id"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
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
		c.SubjectID = req.SubjectID
		c.CourseID =  req.CourseID

		created, err := s.CreateClassroom(ctx, &c)
		if err != nil {
			return nil, err
		}

		return createClassroomResponse{
			UUID:        created.UUID,
			Title:       created.Title,
			Description: created.Description,
			SubjectID:   created.SubjectID,
			CourseID:    created.CourseID,
			CreatedAt:   created.CreatedAt,
			UpdatedAt:   created.UpdatedAt,
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
