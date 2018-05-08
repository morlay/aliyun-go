package sas_api

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type RenewInstanceRequest struct {
	requests.RpcRequest
	Duration     int    `position:"Query" name:"Duration"`
	InstanceId   string `position:"Query" name:"InstanceId"`
	ClientToken  string `position:"Query" name:"ClientToken"`
	BuyNumber    int    `position:"Query" name:"BuyNumber"`
	OwnerId      int64  `position:"Query" name:"OwnerId"`
	PricingCycle string `position:"Query" name:"PricingCycle"`
}

func (req *RenewInstanceRequest) Invoke(client *sdk.Client) (resp *RenewInstanceResponse, err error) {
	req.InitWithApiInfo("Sas-api", "2017-07-05", "RenewInstance", "", "")
	resp = &RenewInstanceResponse{}
	err = client.DoAction(req, resp)
	return
}

type RenewInstanceResponse struct {
	responses.BaseResponse
	Code        common.String
	Message     common.String
	OrderId     common.String
	InstanceId  common.String
	RequestId   common.String
	InstanceIds RenewInstanceInstanceIdList
}

type RenewInstanceInstanceIdList []common.String

func (list *RenewInstanceInstanceIdList) UnmarshalJSON(data []byte) error {
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
