package lesson

import (
	"github.com/gorilla/mux"

	"github.com/go-kit/kit/log"

	"github.com/jinzhu/gorm"
	"github.com/sumelms/microservice-classroom/internal/lesson/database"
	"github.com/sumelms/microservice-classroom/internal/lesson/domain"
	"github.com/sumelms/microservice-classroom/internal/lesson/transport"
)

func NewHTTPService(router *mux.Router, db *gorm.DB, logger log.Logger) {
	repository := database.NewRepository(db, logger)
	service := domain.NewService(repository, logger)

	transport.NewHTTPHandler(router, service, logger)
}
