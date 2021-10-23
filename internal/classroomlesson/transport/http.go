package transport

import (
	"github.com/sumelms/microservice-classroom/internal/classroomlesson/domain"
	"net/http"

	"github.com/sumelms/microservice-classroom/internal/classroomlesson/endpoints"
	"github.com/sumelms/microservice-classroom/pkg/errors"

	"github.com/go-kit/kit/log"
	kittransport "github.com/go-kit/kit/transport"
	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
)

func NewHTTPHandler(r *mux.Router, s domain.ServiceInterface, logger log.Logger) {
	opts := []kithttp.ServerOption{
		kithttp.ServerErrorHandler(kittransport.NewLogErrorHandler(logger)),
		kithttp.ServerErrorEncoder(errors.EncodeError),
	}

	listClassroomHandler := endpoints.NewListClassroomLessonHandler(s, opts...)
	createClassroomHandler := endpoints.NewCreateClassroomLessonHandler(s, opts...)
	findClassroomHandler := endpoints.NewFindClassroomLessonHandler(s, opts...)
	updateClassroomHandler := endpoints.NewUpdateClassroomLessonHandler(s, opts...)
	deleteClassroomHandler := endpoints.NewDeleteClassroomLessonHandler(s, opts...)

	r.Handle("/classroomslesson", createClassroomHandler).Methods(http.MethodPost)
	r.Handle("/classroomslesson", listClassroomHandler).Methods(http.MethodGet)
	r.Handle("/classroomslesson/{uuid}", findClassroomHandler).Methods(http.MethodGet)
	r.Handle("/classroomslesson/{uuid}", updateClassroomHandler).Methods(http.MethodPut)
	r.Handle("/classroomslesson/{uuid}", deleteClassroomHandler).Methods(http.MethodDelete)
}