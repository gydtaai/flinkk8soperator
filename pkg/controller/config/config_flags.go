// Code generated by go generate; DO NOT EDIT.
// This file was generated by robots.

package config

import (
	"encoding/json"
	"reflect"

	"fmt"

	"github.com/spf13/pflag"
)

// If v is a pointer, it will get its element value or the zero value of the element type.
// If v is not a pointer, it will return it as is.
func (Config) elemValueOrNil(v interface{}) interface{} {
	if t := reflect.TypeOf(v); t.Kind() == reflect.Ptr {
		if reflect.ValueOf(v).IsNil() {
			return reflect.Zero(t.Elem()).Interface()
		} else {
			return reflect.ValueOf(v).Interface()
		}
	} else if v == nil {
		return reflect.Zero(t).Interface()
	}

	return v
}

func (Config) mustMarshalJSON(v json.Marshaler) string {
	raw, err := v.MarshalJSON()
	if err != nil {
		panic(err)
	}

	return string(raw)
}

// GetPFlagSet will return strongly types pflags for all fields in Config and its nested types. The format of the
// flags is json-name.json-sub-name... etc.
func (cfg Config) GetPFlagSet(prefix string) *pflag.FlagSet {
	cmdFlags := pflag.NewFlagSet("Config", pflag.ExitOnError)
	cmdFlags.String(fmt.Sprintf("%v%v", prefix, "resyncPeriod"), "30s", "Determines the resync period for all watchers.")
	cmdFlags.String(fmt.Sprintf("%v%v", prefix, "limitNamespace"), "", "Namespaces to watch for by flink operator")
	cmdFlags.String(fmt.Sprintf("%v%v", prefix, "metricsPrefix"), "flinkk8soperator", "Prefix for metrics propagated to prometheus")
	cmdFlags.String(fmt.Sprintf("%v%v", prefix, "profilerPort"), "10254", "Profiler port")
	cmdFlags.String(fmt.Sprintf("%v%v", prefix, "ingressUrlFormat"), *new(string), "")
	cmdFlags.Bool(fmt.Sprintf("%v%v", prefix, "useKubectlProxy"), *new(bool), "")
	cmdFlags.String(fmt.Sprintf("%v%v", prefix, "proxyPort"), "8001", "The port at which flink cluster runs locally")
	cmdFlags.String(fmt.Sprintf("%v%v", prefix, "containerNameFormat"), *new(string), "")
	cmdFlags.Int(fmt.Sprintf("%v%v", prefix, "workers"), 4, "Number of routines to process custom resource")
	cmdFlags.String(fmt.Sprintf("%v%v", prefix, "baseBackoffDuration"), "100ms", "Determines the base backoff for exponential retries.")
	cmdFlags.String(fmt.Sprintf("%v%v", prefix, "maxBackoffDuration"), "30s", "Determines the max backoff for exponential retries.")
	cmdFlags.String(fmt.Sprintf("%v%v", prefix, "maxErrDuration"), "5m", "Determines the max time to wait on errors.")
	return cmdFlags
}
