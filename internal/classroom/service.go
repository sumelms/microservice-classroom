package classroom

import (
	"github.com/go-kit/kit/log"
	"github.com/gorilla/mux"
	"github.com/sumelms/microservice-classroom/internal/classroom/domain"

	"github.com/jinzhu/gorm"
	"github.com/sumelms/microservice-classroom/internal/classroom/database"
	"github.com/sumelms/microservice-classroom/internal/classroom/transport"
)

func NewHTTPService(router *mux.Router, db *gorm.DB, logger log.Logger) {
	repository := database.NewRepository(db, logger)
	service := domain.NewService(repository, logger)

	transport.NewHTTPHandler(router, service, logger)
}
