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
	"github.com/gorilla/mux"
	"github.com/sumelms/microservice-classroom/pkg/validator"
)

type updateClassroomRequest struct {
	UUID        string `json:"uuid" validate:"required"`
	Title       string `json:"title" validate:"required,max=100"`
	Description string `json:"description" validate:"required,max=255"`
	SubjectID   string `json:"subject_id" validate:"required"`
	CourseID    string `json:"course_id" validate:"required"`
}

type updateClassroomResponse struct {
	UUID        string    `json:"uuid"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	SubjectID   string    `json:"subject_id"`
	CourseID    string    `json:"course_id"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
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
		c.SubjectID = req.SubjectID
		c.CourseID = req.CourseID

		updated, err := s.UpdateClassroom(ctx, &c)
		if err != nil {
			return nil, err
		}

		return updateClassroomResponse{
			UUID:        updated.UUID,
			Title:       updated.Title,
			SubjectID:   updated.SubjectID,
			CourseID:    updated.CourseID,
			Description: updated.Description,
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

	req.UUID = id

	return req, nil
}

func encodeUpdateClassroomResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	return kithttp.EncodeJSONResponse(ctx, w, response)
}
