package endpoints

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/go-kit/kit/endpoint"
	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/sumelms/microservice-classroom/internal/lesson/domain"
	"github.com/sumelms/microservice-classroom/pkg/validator"
)

type createLessonRequest struct {
	Title       string `json:"title" validate:"required,max=100"`
	Subtitle    string `json:"subtitle" validate:"required,max=100"`
	Excerpt     string `json:"excerpt" validate:"required,max=140"`
	Description string `json:"description" validate:"required,max=255"`
	Module      string `json:"module" validate:"max=140"`
	SubjectID   string `json:"subject_id" validate:"required"`
}

type createLessonResponse struct {
	UUID        string    `json:"uuid"`
	Title       string    `json:"title"`
	Subtitle    string    `json:"subtitle"`
	Excerpt     string    `json:"excerpt"`
	Description string    `json:"description"`
	Module      string    `json:"module"`
	SubjectID   string    `json:"subject_id"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func NewCreateLessonHandler(s domain.ServiceInterface, opts ...kithttp.ServerOption) *kithttp.Server {
	return kithttp.NewServer(
		makeCreateLessonEndpoint(s),
		decodeCreateLessonRequest,
		encodeCreateLessonResponse,
		opts...,
	)
}

func makeCreateLessonEndpoint(s domain.ServiceInterface) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req, ok := request.(createLessonRequest)
		if !ok {
			return nil, fmt.Errorf("invalid argument")
		}

		v := validator.NewValidator()
		if err := v.Validate(req); err != nil {
			return nil, err
		}

		l := domain.Lesson{}
		data, _ := json.Marshal(req)
		err := json.Unmarshal(data, &l)
		if err != nil {
			return nil, err
		}

		created, err := s.CreateLesson(ctx, &l)
		if err != nil {
			return nil, err
		}

		return createLessonResponse{
			UUID:        created.UUID,
			Title:       created.Title,
			Subtitle:    created.Subtitle,
			Excerpt:     created.Excerpt,
			Description: created.Description,
			Module:      created.Module,
			SubjectID:   created.SubjectID,
			CreatedAt:   created.CreatedAt,
			UpdatedAt:   created.UpdatedAt,
		}, err
	}
}

func decodeCreateLessonRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req createLessonRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, err
	}
	return req, nil
}

func encodeCreateLessonResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	return kithttp.EncodeJSONResponse(ctx, w, response)
}
