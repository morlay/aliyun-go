package sas_api

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type UpgradeInstanceRequest struct {
	requests.RpcRequest
	InstanceId  string `position:"Query" name:"InstanceId"`
	ClientToken string `position:"Query" name:"ClientToken"`
	BuyNumber   int    `position:"Query" name:"BuyNumber"`
	OwnerId     int64  `position:"Query" name:"OwnerId"`
	VersionCode string `position:"Query" name:"VersionCode"`
}

func (req *UpgradeInstanceRequest) Invoke(client *sdk.Client) (resp *UpgradeInstanceResponse, err error) {
	req.InitWithApiInfo("Sas-api", "2017-07-05", "UpgradeInstance", "", "")
	resp = &UpgradeInstanceResponse{}
	err = client.DoAction(req, resp)
	return
}

type UpgradeInstanceResponse struct {
	responses.BaseResponse
	Code        string
	Message     string
	OrderId     string
	InstanceId  string
	RequestId   string
	InstanceIds UpgradeInstanceInstanceIdList
}

type UpgradeInstanceInstanceIdList []string

func (list *UpgradeInstanceInstanceIdList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]string)
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
