package yundun

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type AddCNameWafRequest struct {
	requests.RpcRequest
	InstanceId   string `position:"Query" name:"InstanceId"`
	InstanceType string `position:"Query" name:"InstanceType"`
	Domain       string `position:"Query" name:"Domain"`
}

func (req *AddCNameWafRequest) Invoke(client *sdk.Client) (resp *AddCNameWafResponse, err error) {
	req.InitWithApiInfo("Yundun", "2015-04-16", "AddCNameWaf", "", "")
	resp = &AddCNameWafResponse{}
	err = client.DoAction(req, resp)
	return
}

type AddCNameWafResponse struct {
	responses.BaseResponse
	RequestId   common.String
	WafInfoList AddCNameWafWafInfoList
}

type AddCNameWafWafInfo struct {
	Id     common.Integer
	Domain common.String
	Cname  common.String
	Status common.Integer
}

type AddCNameWafWafInfoList []AddCNameWafWafInfo

func (list *AddCNameWafWafInfoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]AddCNameWafWafInfo)
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
