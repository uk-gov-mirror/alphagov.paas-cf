package main

import (
	"time"

	"code.cloudfoundry.org/lager"
	uaaclient "github.com/cloudfoundry-community/go-uaa"
)

type UAAClientConfig struct {
	Endpoint     string
	ClientID     string
	ClientSecret string
}

func UAAGauges(
	logger lager.Logger,
	cfg *UAAClientConfig,
	interval time.Duration,
) MetricReadCloser {
	return NewMetricPoller(interval, func(w MetricWriter) error {
		lsession := logger.Session("uaa-gauges")

		metrics, err := UAAMetrics(lsession, cfg)
		if err != nil {
			lsession.Error("Failed to get UAA metrics", err)
			return err
		}

		lsession.Info("Writing UAA metrics")
		return w.WriteMetrics(metrics)
	})
}

func UAAMetrics(logger lager.Logger, cfg *UAAClientConfig) ([]Metric, error) {
	lsession := logger.Session("uaa-metrics")
	lsession.Info("Started UAA metrics")

	uaa := uaaclient.New(cfg.Endpoint, "").WithClientCredentials(
		cfg.ClientID, cfg.ClientSecret, uaaclient.JSONWebToken,
	)

	err := uaa.Validate()
	if err != nil {
		lsession.Error("Failed to validate UAA client", err)
		return []Metric{}, err
	}
	lsession.Info("Validated UAA client")

	users, err := uaa.ListAllUsers("", "", "origin", "")
	if err != nil {
		lsession.Error("Failed to list all UAA users", err)
		return []Metric{}, err
	}
	lsession.Info("Listed all UAA users")

	metrics := make([]Metric, 0)

	lsession.Info("Computing UAA users by origin")
	metrics = append(metrics, UAAUsersByOriginGauges(users)...)

	lsession.Info("Finished UAA metrics")
	return metrics, nil
}

func UAAUsersByOriginGauges(users []uaaclient.User) []Metric {
	usersByOriginCount := make(map[string]int, 0)
	for _, user := range users {
		usersByOriginCount[user.Origin] += 1
	}

	metrics := make([]Metric, 0)
	for origin, usersCount := range usersByOriginCount {
		metrics = append(metrics, Metric{
			Kind:  Gauge,
			Time:  time.Now(),
			Name:  "uaa.users",
			Value: float64(usersCount),
			Tags:  MetricTags{MetricTag{Label: "origin", Value: origin}},
			Unit:  "count",
		})
	}

	return metrics
}
