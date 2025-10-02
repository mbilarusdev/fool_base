package gateway_utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/mbilarusdev/fool_base/v2/log"
	gateway_models "github.com/mbilarusdev/fool_base/v2/utils/gateway/models"
	"go.uber.org/zap"
)

func RegisterInGateway(form gateway_models.RegisterForm) {
	err := tryPostToGateway(form)

	if err != nil {
		log.Err(
			err,
			"Failed to register in gateway: ",
			zap.String(form.ServiceName, joinUrl(form.ServiceProtocol, form.ServiceRemoteAddr)),
			zap.String("Gateway", joinUrl(form.GatewayProtocol, form.GatewayAddr)),
		)
		return
	}

	log.Info("Gateway connected with success")
}

func tryPostToGateway(form gateway_models.RegisterForm) error {
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
