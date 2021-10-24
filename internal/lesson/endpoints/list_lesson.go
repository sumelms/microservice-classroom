package endpoints

import (
	"context"
	"fmt"
	"net/http"

	"github.com/go-kit/kit/endpoint"
	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/sumelms/microservice-classroom/internal/lesson/domain"
)

type listLessonRequest struct {
	SubjectID string `json:"subject_id,omitempty"`
	Module    string `json:"module,omitempty"`
}

type listLessonResponse struct {
	Lessons []findLessonResponse `json:"lessons"`
}

func NewListLessonHandler(s domain.ServiceInterface, opts ...kithttp.ServerOption) *kithttp.Server {
	return kithttp.NewServer(
		makeListLessonEndpoint(s),
		decodeListLessonRequest,
		encodeListLessonResponse,
		opts...,
	)
}

func makeListLessonEndpoint(s domain.ServiceInterface) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req, ok := request.(listLessonRequest)
		if !ok {
			return nil, fmt.Errorf("invalid argument")
		}

		filters := make(map[string]interface{})
		if len(req.SubjectID) > 0 {
			filters["subject_id"] = req.SubjectID
		}
		if len(req.Module) > 0 {
			filters["module"] = req.Module
		}

		lessons, err := s.ListLesson(ctx, filters)
		if err != nil {
			return nil, err
		}

		var list []findLessonResponse
		for i := range lessons {
			l := lessons[i]
			list = append(list, findLessonResponse{
				UUID:        l.UUID,
				Title:       l.Title,
				Subtitle:    l.Subtitle,
				Excerpt:     l.Excerpt,
				Description: l.Description,
				Module:      l.Module,
				SubjectID:   l.SubjectID,
				CreatedAt:   l.CreatedAt,
				UpdatedAt:   l.UpdatedAt,
			})
		}

		return &listLessonResponse{Lessons: list}, nil
	}
}

func decodeListLessonRequest(_ context.Context, r *http.Request) (interface{}, error) {
	subjectID := r.FormValue("subject_id")
	module := r.FormValue("module")
	return listLessonRequest{SubjectID: subjectID, Module: module}, nil
}

func encodeListLessonResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	return kithttp.EncodeJSONResponse(ctx, w, response)
}
