package endpoints

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/go-kit/kit/endpoint"
	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	"github.com/sumelms/microservice-classroom/internal/lesson/domain"
)

type findLessonRequest struct {
	UUID string `json:"uuid"`
}

type findLessonResponse struct {
	UUID        string    `json:"uuid"`
	Title       string    `json:"title"`
	Subtitle    string    `json:"subtitle"`
	Excerpt     string    `json:"excerpt"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func NewFindLessonHandler(s domain.ServiceInterface, opts ...kithttp.ServerOption) *kithttp.Server {
	return kithttp.NewServer(
		makeFindLessonEndpoint(s),
		decodeFindLessonRequest,
		encodeFindLessonResponse,
		opts...,
	)
}

func makeFindLessonEndpoint(s domain.ServiceInterface) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req, ok := request.(findLessonRequest)
		if !ok {
			return nil, fmt.Errorf("invalid argument")
		}

		l, err := s.FindLesson(ctx, req.UUID)
		if err != nil {
			return nil, err
		}

		return &findLessonResponse{
			UUID:        l.UUID,
			Title:       l.Title,
			Subtitle:    l.Subtitle,
			Excerpt:     l.Excerpt,
			Description: l.Description,
		}, nil
	}
}

func decodeFindLessonRequest(_ context.Context, r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)
	id, ok := vars["uuid"]
	if !ok {
		return nil, fmt.Errorf("invalid argument")
	}

	return findLessonRequest{UUID: id}, nil
}

func encodeFindLessonResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	return kithttp.EncodeJSONResponse(ctx, w, response)
}
