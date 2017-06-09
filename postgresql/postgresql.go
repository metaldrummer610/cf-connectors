package postgresql

import (
	"encoding/json"

	"github.com/pkg/errors"
	connectors "pivotal.io/cf-connectors"
)

const PostgresTag = "postgresql"

type PostgresCredentials struct {
	Uri      string `json:"uri"`
	MaxConns string `json:"max_conns"`
}

type PostgreSQLConfig struct {
	Credentials    PostgresCredentials `json:"credentials"`
	SyslogDrainUrl string `json:"syslog_drain_url"`
	VolumeMounts   []string `json:"volume_mounts"`
	Label          string `json:"label"`
	Provider       string `json:"provider"`
	Plan           string `json:"plan"`
	Name           string `json:"name"`
	Tags           []string `json:"tags"`
}

func Parse(vcap *connectors.VcapServices) ([]*PostgreSQLConfig, error) {
	configs := []*PostgreSQLConfig{}

	for _, v := range *vcap {
		for _, c := range v {
			minimal := connectors.MinimalVcapService{}
			if err := json.Unmarshal(*c, &minimal); err != nil {
				return nil, errors.WithStack(err)
			}

			for _, s := range minimal.Tags {
				if s == PostgresTag {
					config := &PostgreSQLConfig{}
					if err := json.Unmarshal(*c, config); err != nil {
						return nil, errors.WithStack(err)
					}

					configs = append(configs, config)
					break
				}
			}
		}
	}

	return configs, nil
}
