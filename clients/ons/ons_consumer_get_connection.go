package ons

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type OnsConsumerGetConnectionRequest struct {
	PreventCache int64  `position:"Query" name:"PreventCache"`
	OnsRegionId  string `position:"Query" name:"OnsRegionId"`
	OnsPlatform  string `position:"Query" name:"OnsPlatform"`
	ConsumerId   string `position:"Query" name:"ConsumerId"`
}

func (r OnsConsumerGetConnectionRequest) Invoke(client *sdk.Client) (response *OnsConsumerGetConnectionResponse, err error) {
	req := struct {
		*requests.RpcRequest
		OnsConsumerGetConnectionRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Ons", "2017-09-18", "OnsConsumerGetConnection", "", "")

	resp := struct {
		*responses.BaseResponse
		OnsConsumerGetConnectionResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.OnsConsumerGetConnectionResponse

	err = client.DoAction(&req, &resp)
	return
}

type OnsConsumerGetConnectionResponse struct {
	RequestId string
	HelpUrl   string
	Data      OnsConsumerGetConnectionData
}

type OnsConsumerGetConnectionData struct {
	ConnectionList OnsConsumerGetConnectionConnectionDoList
}

type OnsConsumerGetConnectionConnectionDo struct {
	ClientId   string
	ClientAddr string
	Language   string
	Version    string
}

type OnsConsumerGetConnectionConnectionDoList []OnsConsumerGetConnectionConnectionDo

func (list *OnsConsumerGetConnectionConnectionDoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]OnsConsumerGetConnectionConnectionDo)
	err := json.Unmarshal(data, &m)
	if err != nil {
		return err
	}
	for _, v := range m {
		*list = v
		break
	}
	return nil
}
