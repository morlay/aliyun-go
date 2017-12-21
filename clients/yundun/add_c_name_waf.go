package yundun

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
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
	RequestId   string
	WafInfoList AddCNameWafWafInfoList
}

type AddCNameWafWafInfo struct {
	Id     int
	Domain string
	Cname  string
	Status int
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
