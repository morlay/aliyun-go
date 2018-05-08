package sas_api

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type CreateInstanceRequest struct {
	requests.RpcRequest
	Duration          int    `position:"Query" name:"Duration"`
	IsAutoRenew       string `position:"Query" name:"IsAutoRenew"`
	ClientToken       string `position:"Query" name:"ClientToken"`
	BuyNumber         int    `position:"Query" name:"BuyNumber"`
	OwnerId           int64  `position:"Query" name:"OwnerId"`
	VersionCode       string `position:"Query" name:"VersionCode"`
	PricingCycle      string `position:"Query" name:"PricingCycle"`
	AutoRenewDuration int    `position:"Query" name:"AutoRenewDuration"`
}

func (req *CreateInstanceRequest) Invoke(client *sdk.Client) (resp *CreateInstanceResponse, err error) {
	req.InitWithApiInfo("Sas-api", "2017-07-05", "CreateInstance", "", "")
	resp = &CreateInstanceResponse{}
	err = client.DoAction(req, resp)
	return
}

type CreateInstanceResponse struct {
	responses.BaseResponse
	Code        common.String
	Message     common.String
	OrderId     common.String
	InstanceId  common.String
	RequestId   common.String
	InstanceIds CreateInstanceInstanceIdList
}

type CreateInstanceInstanceIdList []common.String

func (list *CreateInstanceInstanceIdList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]common.String)
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
