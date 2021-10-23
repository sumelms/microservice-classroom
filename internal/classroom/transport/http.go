package transport

import (
	"github.com/sumelms/microservice-classroom/internal/classroom/domain"
	"net/http"

	"github.com/sumelms/microservice-classroom/internal/classroom/endpoints"
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

	listClassroomHandler := endpoints.NewListClassroomHandler(s, opts...)
	createClassroomHandler := endpoints.NewCreateClassroomHandler(s, opts...)
	findClassroomHandler := endpoints.NewFindClassroomHandler(s, opts...)
	updateClassroomHandler := endpoints.NewUpdateClassroomHandler(s, opts...)
	deleteClassroomHandler := endpoints.NewDeleteClassroomHandler(s, opts...)

	r.Handle("/classrooms", createClassroomHandler).Methods(http.MethodPost)
	r.Handle("/classrooms", listClassroomHandler).Methods(http.MethodGet)
	r.Handle("/classrooms/{uuid}", findClassroomHandler).Methods(http.MethodGet)
	r.Handle("/classrooms/{uuid}", updateClassroomHandler).Methods(http.MethodPut)
	r.Handle("/classrooms/{uuid}", deleteClassroomHandler).Methods(http.MethodDelete)
}
