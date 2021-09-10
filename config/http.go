package config

import "time"

const (
	ListenHostEnv         = "LISTEN_HOST"
	ListenPortEnv         = "LISTEN_PORT"
	ServerReadTimeoutEnv  = "SERVER_READ_TIMEOUT_SECONDS"
	ServerWriteTimeoutEnv = "SERVER_WRITE_TIMEOUT_SECONDS"
)

type Http struct {
	ListenHost string `json:"listenHost"`
	ListenPort string `json:"listenPort"`

	ServerRTimeout time.Duration `json:"serverRTimeout"`
	ServerWTimeout time.Duration `json:"serverWTimeout"`
}

func (h *Http) Init() error {
	if ListenHost, err := getStringFromEnv(ListenHostEnv); h.ListenHost == "" && err != nil {
		return err
	} else {
		h.ListenHost = ListenHost
	}

	if ListenPort, err := getStringFromEnv(ListenPortEnv); h.ListenPort == "" && err != nil {
		return err
	} else {
		h.ListenPort = ListenPort
	}

	if ServerRTimeout, err := getSecondsDurationFromEnv(ServerReadTimeoutEnv); h.ServerRTimeout == 0 && err != nil {
		return err
	} else {
		h.ServerRTimeout = ServerRTimeout
	}

	if ServerWTimeout, err := getSecondsDurationFromEnv(ServerWriteTimeoutEnv); h.ServerWTimeout == 0 && err != nil {
		return err
	} else {
		h.ServerWTimeout = ServerWTimeout
	}
	return nil
}
