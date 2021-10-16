package endpoints

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/go-kit/kit/endpoint"
	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	"github.com/sumelms/microservice-classroom/internal/classroom/domain"
)

type findClassroomRequest struct {
	UUID string `json:"uuid"`
}

type findClassroomResponse struct {
	UUID        string    `json:"uuid"`
	Title       string    `json:"title"`
	Subtitle    string    `json:"subtitle"`
	Excerpt     string    `json:"excerpt"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
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

		c, err := s.FindClassroom(ctx, req.UUID)
		if err != nil {
			return nil, err
		}

		return &findClassroomResponse{
			UUID:        c.UUID,
			Title:       c.Title,
			Subtitle:    c.Subtitle,
			Excerpt:     c.Excerpt,
			Description: c.Description,
		}, nil
	}
}

func decodeFindClassroomRequest(_ context.Context, r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)
	id, ok := vars["uuid"]
	if !ok {
		return nil, fmt.Errorf("invalid argument")
	}

	return findClassroomRequest{UUID: id}, nil
}

func encodeFindClassroomResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	return kithttp.EncodeJSONResponse(ctx, w, response)
}
