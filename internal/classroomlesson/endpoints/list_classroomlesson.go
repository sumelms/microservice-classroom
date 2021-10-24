package endpoints

import (
	"context"
	"fmt"
	"github.com/sumelms/microservice-classroom/internal/classroomlesson/domain"
	"net/http"

	"github.com/go-kit/kit/endpoint"
	kithttp "github.com/go-kit/kit/transport/http"
)

type listClassroomLessonRequest struct {
	ClassroomID string `json:"classroom_id,omitempty"`
	LessonID    string `json:"lesson_id,omitempty"`
}

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
		req, ok := request.(listClassroomLessonRequest)
		if !ok {
			return nil, fmt.Errorf("invalid argument")
		}

		filters := make(map[string]interface{})
		if len(req.ClassroomID) > 0 {
			filters["classroom_id"] = req.ClassroomID
		}
		if len(req.LessonID) > 0 {
			filters["lesson_id"] = req.LessonID
		}

		classrooms, err := s.ListClassroomLesson(ctx, filters)
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

func decodeListClassroomLessonRequest(_ context.Context, r *http.Request) (interface{}, error) {
	classroomID := r.FormValue("classroom_id")
	lessonID := r.FormValue("lesson_id")
	return listClassroomLessonRequest{ClassroomID: classroomID, LessonID: lessonID}, nil
}

func encodeListClassroomLessonResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	return kithttp.EncodeJSONResponse(ctx, w, response)
}
