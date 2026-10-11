package weather_test

import (
	"testing"

	"example.com/weather"
)

func TestForecast(t *testing.T) {
	if got := weather.Forecast(20); got != "68°F" {
		t.Errorf("Forecast(20) = %q, want %q", got, "68°F")
	}
}
