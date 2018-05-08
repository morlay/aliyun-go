package emr

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type ModifyPrepayInstanceSpecRequest struct {
	requests.RpcRequest
	InstanceTypeConfigs *ModifyPrepayInstanceSpecInstanceTypeConfigList `position:"Query" type:"Repeated" name:"InstanceTypeConfig"`
	ResourceOwnerId     int64                                           `position:"Query" name:"ResourceOwnerId"`
	ClusterId           string                                          `position:"Query" name:"ClusterId"`
}

func (req *ModifyPrepayInstanceSpecRequest) Invoke(client *sdk.Client) (resp *ModifyPrepayInstanceSpecResponse, err error) {
	req.InitWithApiInfo("Emr", "2016-04-08", "ModifyPrepayInstanceSpec", "", "")
	resp = &ModifyPrepayInstanceSpecResponse{}
	err = client.DoAction(req, resp)
	return
}

type ModifyPrepayInstanceSpecInstanceTypeConfig struct {
	HostGroupId  string `name:"HostGroupId"`
	InstanceType string `name:"InstanceType"`
}

type ModifyPrepayInstanceSpecResponse struct {
	responses.BaseResponse
	RequestId common.String
	ClusterId common.String
}

type ModifyPrepayInstanceSpecInstanceTypeConfigList []ModifyPrepayInstanceSpecInstanceTypeConfig

func (list *ModifyPrepayInstanceSpecInstanceTypeConfigList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ModifyPrepayInstanceSpecInstanceTypeConfig)
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
