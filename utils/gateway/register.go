package gateway_utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/mbilarusdev/fool_base/v2/log"
	"go.uber.org/zap"
)

type RegisterForm struct {
	ServiceName       string `json:"service_name"`
	GatewayAddr       string
	GatewayProtocol   string
	ServiceRemoteAddr string `json:"service_remote_addr"`
	ServiceProtocol   string `json:"service_protocol"`
}

func RegisterInGateway(form RegisterForm) {
	retryTimeout := time.Second * 10
	for {
		err := tryPostToGateway(form)
		if err != nil {
			log.Err(
				err,
				"Failed to register in gateway: ",
				zap.String(form.ServiceName, joinUrl(form.ServiceProtocol, form.ServiceRemoteAddr)),
				zap.String("Gateway", joinUrl(form.GatewayProtocol, form.GatewayAddr)),
			)
			time.Sleep(retryTimeout)
			continue
		}
		break
	}
}

func tryPostToGateway(form RegisterForm) error {
	jsonForm, err := json.Marshal(form)
	if err != nil {
		return err
	}

	req, err := http.NewRequest(
		http.MethodPost,
		fmt.Sprintf("%v://%v", form.GatewayProtocol, form.GatewayAddr),
		bytes.NewBuffer(jsonForm),
	)
	if err != nil {
		return err
	}

	client := new(http.Client)

	res, err := client.Do(req)
	if err != nil || res.StatusCode != http.StatusOK {
		return err
	}

	return nil
}

func joinUrl(addr string, protocol string) string {
	return fmt.Sprintf("%v://%v", protocol, addr)
}
