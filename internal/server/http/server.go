package internalhttp

import (
	"context"
	"fmt"
	"net/http"

	"github.com/pkg/errors"

	"catalog/internal/app"
	"catalog/internal/config"
)

type Server interface {
	Start(ctx context.Context) error
	Stop(ctx context.Context) error
}

type server struct {
	app *app.App
	cfg *config.Config
	srv *http.Server
}

type Application interface {
	Handle(phrase string) error
}

func NewServer(cfg *config.Config, app *app.App) Server {
	return &server{
		cfg: cfg,
		app: app,
	}
}

func (m *server) Start(ctx context.Context) error {
	router := NewGinRouter(m.app)

	m.srv = &http.Server{}
	m.srv.Addr = m.cfg.Server.HTTP.Host + ":" + m.cfg.Server.HTTP.Port
	m.srv.Handler = router

	fmt.Println("listen: ", m.srv.Addr)

	err := m.srv.ListenAndServe()
	if err != nil {
		return errors.Wrap(err, "cannot listen and serve")
	}

	return nil
}

func (m *server) Stop(ctx context.Context) error {
	err := m.srv.Shutdown(ctx)
	if err != nil {
		return errors.Wrap(err, "cannot shutdown server")
	}

	return nil
}
