package routes

import (
	"encoding/json"
	"fmt"
	"github.com/go-kit/kit/metrics"
	dkmetrics "github.com/horahoradev/DataKhan/backend/internal/metrics"
	echo "github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
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
	r.CreateCounter(fmt.Sprintf(".%s.NumberRequests;path=%s;useragent=%s;ip=%s;hostname=%s", uri.Hostname(), uri.Path, parsed.Useragent, parsed.IP, strings.Replace(uri.Hostname(), ".", "", -1))).Add(1)
	r.CreateCounter("."+uri.Hostname()+".NumberRequests").With("path", uri.Path).Add(1)

	log.Infof("Created counter for website %s", uri.Hostname())
	return nil
	// TODO
}

func (r *RouteHandler) handleRequestsCount(c echo.Context) error {
	resp, err := http.Get("http://graphite/render?target=sumSeries(seriesByTag(%27hostname=myaweseomwebsitecom%27))&format=json")
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	return c.String(http.StatusOK, string(b))
}

func (r *RouteHandler) handleEvent(c echo.Context) error {
	return nil
}

func SetupRoutes(e *echo.Echo) {
	r := NewRouteHandler(dkmetrics.ConcreteCounter)

	e.POST("/view/", r.handleView)
	e.POST("/event/", r.handleEvent)
	e.GET("/requestcounts/", r.handleRequestsCount)
}
