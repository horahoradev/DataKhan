package routes

import (
	"encoding/json"
	echo "github.com/labstack/echo/v4"
	prometheus "github.com/prometheus/client_golang/prometheus"
	"io/ioutil"
	"net/url"
)

type RouteHandler struct {
	RequestCounter *prometheus.CounterVec
}

func NewRouteHandler() RouteHandler {
	c := prometheus.NewCounterVec(prometheus.CounterOpts{
		Namespace: "ViewStats",
		//Subsystem:   subsystem,
		Name: "Number_of_Requests",
		Help: "The total number of requests received",
		//ConstLabels:
	}, []string{"URI", "Path", "IP", "UserAgent"})
	prometheus.MustRegister(c)

	return RouteHandler{c}
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

	// Increment the number of requests
	requestLabels := []string{uri.Hostname(), uri.Path, parsed.IP, parsed.Useragent}
	r.RequestCounter.WithLabelValues(requestLabels...).Add(1)

	return nil
	// TODO
}

func (r *RouteHandler) handleEvent(c echo.Context) error {
	return nil
}

func SetupRoutes(e *echo.Echo) {
	r := NewRouteHandler()

	e.POST("/view/", r.handleView)
	e.POST("/event/", r.handleEvent)
}
