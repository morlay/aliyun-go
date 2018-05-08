package ccc

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type ListConfigRequest struct {
	requests.RpcRequest
	InstanceId  string                    `position:"Query" name:"InstanceId"`
	ConfigItems *ListConfigConfigItemList `position:"Query" type:"Repeated" name:"ConfigItem"`
}

func (req *ListConfigRequest) Invoke(client *sdk.Client) (resp *ListConfigResponse, err error) {
	req.InitWithApiInfo("CCC", "2017-07-05", "ListConfig", "ccc", "")
	resp = &ListConfigResponse{}
	err = client.DoAction(req, resp)
	return
}

type ListConfigResponse struct {
	responses.BaseResponse
	RequestId      common.String
	Success        bool
	Code           common.String
	Message        common.String
	HttpStatusCode common.Integer
	ConfigItems    ListConfigConfigItemList
}

type ListConfigConfigItem struct {
	Name  common.String
	Value common.String
}

type ListConfigConfigItemList []ListConfigConfigItem

func (list *ListConfigConfigItemList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListConfigConfigItem)
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
