package subscription

import (
	"github.com/go-kit/kit/log"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"github.com/sumelms/microservice-classroom/internal/subscription/database"
	"github.com/sumelms/microservice-classroom/internal/subscription/domain"
	"github.com/sumelms/microservice-classroom/internal/subscription/transport"
)

func NewHTTPService(router *mux.Router, db *gorm.DB, logger log.Logger) {
	repository := database.NewRepository(db, logger)
	service := domain.NewService(repository, logger)

	transport.NewHTTPHandler(router, service, logger)
}
