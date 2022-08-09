package classroom

import (
	"github.com/go-kit/kit/log"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"

	"github.com/sumelms/microservice-classroom/internal/classroom/domain"

	"github.com/sumelms/microservice-classroom/internal/classroom/database"
	"github.com/sumelms/microservice-classroom/internal/classroom/transport"
)

func NewService(db *sqlx.DB, logger log.Logger) (*domain.Service, error) {
	classroom, err := database.NewClassroomRepository(db)
	if err != nil {
		return nil, err
	}

	service, err := domain.NewService(
		domain.WithLogger(logger),
		domain.WithClassroomRepository(classroom))
	if err != nil {
		return nil, err
	}
	return service, nil
}

func NewHTTPService(router *mux.Router, service domain.ServiceInterface, logger log.Logger) error {
	transport.NewHTTPHandler(router, service, logger)
	return nil
}
