package endpoints

import (
	"context"
	"net/http"

	"github.com/sumelms/microservice-classroom/internal/classroom/domain"

	"github.com/go-kit/kit/endpoint"
	kithttp "github.com/go-kit/kit/transport/http"
)

type listClassroomResponse struct {
	Classrooms []findClassroomResponse `json:"classrooms"`
}

func NewListClassroomHandler(s domain.ServiceInterface, opts ...kithttp.ServerOption) *kithttp.Server {
	return kithttp.NewServer(
		makeListClassroomEndpoint(s),
		decodeListClassroomRequest,
		encodeListClassroomResponse,
		opts...,
	)
}

func makeListClassroomEndpoint(s domain.ServiceInterface) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		classrooms, err := s.Classrooms(ctx)
		if err != nil {
			return nil, err
		}

		var list []findClassroomResponse
		for i := range classrooms {
			c := classrooms[i]
			list = append(list, findClassroomResponse{
				UUID:        c.UUID,
				Title:       c.Title,
				Description: c.Description,
				SubjectID:   c.SubjectID,
				CourseID:    c.CourseID,
				CreatedAt:   c.CreatedAt,
				UpdatedAt:   c.UpdatedAt,
			})
		}

		return &listClassroomResponse{Classrooms: list}, nil
	}
}

func decodeListClassroomRequest(_ context.Context, _ *http.Request) (interface{}, error) {
	return nil, nil
}

func encodeListClassroomResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	return kithttp.EncodeJSONResponse(ctx, w, response)
}
