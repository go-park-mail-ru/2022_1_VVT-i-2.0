package metrics

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
	"time"
)

func (m *Metric) CollectMetrics(next echo.HandlerFunc) echo.HandlerFunc {
	return func(context echo.Context) error {
		timeStart := time.Now()

		defer func() {
			m.TotalHits.Inc()
			if context.Response().Status != http.StatusOK && context.Response().Status != http.StatusNoContent {
				m.Errors.WithLabelValues(strconv.Itoa(context.Response().Status), context.Request().Method, context.Request().RequestURI).Inc()
			} else {
				m.Hits.WithLabelValues(strconv.Itoa(context.Response().Status), context.Request().Method, context.Request().RequestURI).Inc()
			}

			m.Durations.WithLabelValues(strconv.Itoa(context.Response().Status), context.Request().Method, context.Request().RequestURI).Observe(time.Since(timeStart).Seconds())
		}()

		return next(context)
	}
}
