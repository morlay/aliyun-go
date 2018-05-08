package ons

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type OnsConsumerGetConnectionRequest struct {
	requests.RpcRequest
	PreventCache int64  `position:"Query" name:"PreventCache"`
	OnsRegionId  string `position:"Query" name:"OnsRegionId"`
	OnsPlatform  string `position:"Query" name:"OnsPlatform"`
	ConsumerId   string `position:"Query" name:"ConsumerId"`
}

func (req *OnsConsumerGetConnectionRequest) Invoke(client *sdk.Client) (resp *OnsConsumerGetConnectionResponse, err error) {
	req.InitWithApiInfo("Ons", "2017-09-18", "OnsConsumerGetConnection", "", "")
	resp = &OnsConsumerGetConnectionResponse{}
	err = client.DoAction(req, resp)
	return
}

type OnsConsumerGetConnectionResponse struct {
	responses.BaseResponse
	RequestId common.String
	HelpUrl   common.String
	Data      OnsConsumerGetConnectionData
}

type OnsConsumerGetConnectionData struct {
	ConnectionList OnsConsumerGetConnectionConnectionDoList
}

type OnsConsumerGetConnectionConnectionDo struct {
	ClientId   common.String
	ClientAddr common.String
	Language   common.String
	Version    common.String
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
