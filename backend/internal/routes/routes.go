package routes

import (
	"encoding/json"
	"fmt"
	"github.com/go-kit/kit/metrics"
	"github.com/go-kit/kit/metrics/graphite"
	dkmetrics "github.com/horahoradev/DataKhan/backend/internal/metrics"
	echo "github.com/labstack/echo/v4"
	"io/ioutil"
	"net/url"
)

type RouteHandler struct {
	CreateCounter func(name string) metrics.Counter
}

func NewRouteHandler(createCounter func(name string) metrics.Counter) RouteHandler {

	return RouteHandler{CreateCounter: createCounter}
}

type ViewObj struct {
	IP        string `json:"ip"`
	Useragent string `json:"useragent"`
	URI       string `json:"uri"`
}

func (r *RouteHandler) handleView(c echo.Context) error {
	b, err := ioutil.ReadAll(c.Request().Body)

	var parsed ViewObj
	err = json.Unmarshal(b, &parsed)
	if err != nil {
		return err
	}

	uri, err := url.Parse(parsed.URI)
	if err != nil {
		return err
	}

	// Increment the number of requests for path
	r.CreateCounter(fmt.Sprintf("%s.NumberRequests;path=%s;useragent=%s;ip=%s", uri.Hostname(), uri.Path, parsed.Useragent, parsed.IP))
	return nil
	// TODO
}

func (r *RouteHandler) handleEvent(c echo.Context) error {
	return nil
}

func SetupRoutes(e *echo.Echo) {
	r := NewRouteHandler(dkmetrics.ConcreteCounter)

	e.POST("/view/", r.handleView)
	e.POST("/event/", r.handleEvent)
}
