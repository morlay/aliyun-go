package emr

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ModifyChargeTypeRequest struct {
	requests.RpcRequest
	ResourceOwnerId   int64                                 `position:"Query" name:"ResourceOwnerId"`
	ChargeTypeConfigs *ModifyChargeTypeChargeTypeConfigList `position:"Query" type:"Repeated" name:"ChargeTypeConfig"`
	ClusterId         string                                `position:"Query" name:"ClusterId"`
}

func (req *ModifyChargeTypeRequest) Invoke(client *sdk.Client) (resp *ModifyChargeTypeResponse, err error) {
	req.InitWithApiInfo("Emr", "2016-04-08", "ModifyChargeType", "", "")
	resp = &ModifyChargeTypeResponse{}
	err = client.DoAction(req, resp)
	return
}

type ModifyChargeTypeChargeTypeConfig struct {
	HostGroupId string `name:"HostGroupId"`
	ChargeType  string `name:"ChargeType"`
	Period      int    `name:"Period"`
}

type ModifyChargeTypeResponse struct {
	responses.BaseResponse
	RequestId string
	ClusterId string
}

type ModifyChargeTypeChargeTypeConfigList []ModifyChargeTypeChargeTypeConfig

func (list *ModifyChargeTypeChargeTypeConfigList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ModifyChargeTypeChargeTypeConfig)
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
