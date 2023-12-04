package app

import (
	"database/sql"
	"log"
	"net/http"
)

type App struct {
	log        log.Logger
	httpServer http.Server
	database   sql.DB
}
