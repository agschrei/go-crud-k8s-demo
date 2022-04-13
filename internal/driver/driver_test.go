package driver_test

import (
	"strconv"
	"testing"
	"time"

	"github.com/agschrei/integration-test-sample/internal/config"
	"github.com/agschrei/integration-test-sample/internal/driver"
)

func TestDsnFromDbConfig(t *testing.T) {
	type testCase struct {
		dbConfig        *config.DbConfig
		expectedDsn     string
		shouldHaveError bool
	}

	for num, tc := range []testCase{
		{
			&config.DbConfig{},
			"",
			true,
		},
		{
			&config.DbConfig{
				Hostname: "test",
			},
			"postgresql://test",
			false,
		},
		{
			&config.DbConfig{
				Hostname: "test",
				Port:     "1234",
			},
			"postgresql://test:1234",
			false,
		},
		{
			&config.DbConfig{
				Hostname: "test",
				Port:     "1234",
				User:     "myUser",
			},
			"postgresql://myUser@test:1234",
			false,
		},
		{
			&config.DbConfig{
				Hostname: "test",
				Port:     "1234",
				User:     "myUser",
				Pass:     "mySecret",
			},
			"postgresql://myUser:mySecret@test:1234",
			false,
		},
		{
			&config.DbConfig{
				Hostname: "test",
				Port:     "1234",
				User:     "myUser",
				Pass:     "mySecret",
				DbName:   "myDb",
			},
			"postgresql://myUser:mySecret@test:1234/myDb",
			false,
		},
		{
			&config.DbConfig{
				Hostname: "test",
				User:     "myUser",
				DbName:   "myDb",
			},
			"postgresql://myUser@test/myDb",
			false,
		},
		{
			&config.DbConfig{
				Hostname:          "test",
				User:              "myUser",
				DbName:            "myDb",
				ConnectionTimeout: 5 * time.Second,
			},
			"postgresql://myUser@test/myDb?connect_timeout=5",
			false,
		},
		{
			&config.DbConfig{
				Hostname:    "test",
				User:        "myUser",
				DbName:      "myDb",
				SslDisabled: true,
			},
			"postgresql://myUser@test/myDb?sslmode=disable",
			false,
		},
		{
			&config.DbConfig{
				Hostname:          "test",
				User:              "myUser",
				DbName:            "myDb",
				ConnectionTimeout: 5 * time.Second,
				SslDisabled:       true,
			},
			"postgresql://myUser@test/myDb?connect_timeout=5&sslmode=disable",
			false,
		},
		{
			&config.DbConfig{
				Hostname:          "test",
				Port:              "1234",
				User:              "myUser",
				Pass:              "mySecret",
				DbName:            "myDb",
				ConnectionTimeout: 5 * time.Second,
				SslDisabled:       true,
			},
			"postgresql://myUser:mySecret@test:1234/myDb?connect_timeout=5&sslmode=disable",
			false,
		},
	} {
		t.Run(strconv.Itoa(num), func(t *testing.T) {
			dsn, err := driver.DsnFromDbConfig(tc.dbConfig)
			if tc.shouldHaveError && err == nil {
				t.Errorf("expected error, but got nil: %+v", tc.dbConfig)
			}
			if dsn != tc.expectedDsn {
				t.Errorf("expected: '%s' but got: '%s': %+v", tc.expectedDsn, dsn, tc.dbConfig)
			}
		})
	}
}
