package yundun

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type DeleteCNameWafRequest struct {
	requests.RpcRequest
	InstanceId   string `position:"Query" name:"InstanceId"`
	Domain       string `position:"Query" name:"Domain"`
	CnameId      int    `position:"Query" name:"CnameId"`
	InstanceType string `position:"Query" name:"InstanceType"`
}

func (req *DeleteCNameWafRequest) Invoke(client *sdk.Client) (resp *DeleteCNameWafResponse, err error) {
	req.InitWithApiInfo("Yundun", "2015-04-16", "DeleteCNameWaf", "", "")
	resp = &DeleteCNameWafResponse{}
	err = client.DoAction(req, resp)
	return
}

type DeleteCNameWafResponse struct {
	responses.BaseResponse
	RequestId   common.String
	WafInfoList DeleteCNameWafWafInfoList
}

type DeleteCNameWafWafInfo struct {
	Id     common.Integer
	Domain common.String
	Cname  common.String
	Status common.Integer
}

type DeleteCNameWafWafInfoList []DeleteCNameWafWafInfo

func (list *DeleteCNameWafWafInfoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DeleteCNameWafWafInfo)
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
