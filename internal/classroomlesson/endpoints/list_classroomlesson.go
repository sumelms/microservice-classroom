package endpoints

import (
	"context"
	"github.com/sumelms/microservice-classroom/internal/classroomlesson/domain"
	"net/http"

	"github.com/go-kit/kit/endpoint"
	kithttp "github.com/go-kit/kit/transport/http"
)

type listClassroomLessonResponse struct {
	ClassroomLessons []findClassroomLessonResponse `json:"classroom_lessons"`
}

func NewListClassroomLessonHandler(s domain.ServiceInterface, opts ...kithttp.ServerOption) *kithttp.Server {
	return kithttp.NewServer(
		makeListClassroomLessonEndpoint(s),
		decodeListClassroomLessonRequest,
		encodeListClassroomLessonResponse,
		opts...,
	)
}

func makeListClassroomLessonEndpoint(s domain.ServiceInterface) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		classrooms, err := s.ListClassroomLesson(ctx)
		if err != nil {
			return nil, err
		}

		var list []findClassroomLessonResponse
		for i := range classrooms {
			c := classrooms[i]
			list = append(list, findClassroomLessonResponse{
				UUID:        c.UUID,
				ClassroomID: c.ClassroomID,
				LessonID:    c.LessonID,
				StartsAt:    c.StartsAt,
				EndsAt:      c.EndsAt,
				CreatedAt:   c.CreatedAt,
				UpdatedAt:   c.UpdatedAt,
			})
		}

		return &listClassroomLessonResponse{ClassroomLessons: list}, nil
	}
}

func decodeListClassroomLessonRequest(_ context.Context, _ *http.Request) (interface{}, error) {
	return nil, nil
}

func encodeListClassroomLessonResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	return kithttp.EncodeJSONResponse(ctx, w, response)
}
