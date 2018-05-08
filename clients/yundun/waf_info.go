package yundun

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type WafInfoRequest struct {
	requests.RpcRequest
	InstanceId   string `position:"Query" name:"InstanceId"`
	InstanceType string `position:"Query" name:"InstanceType"`
}

func (req *WafInfoRequest) Invoke(client *sdk.Client) (resp *WafInfoResponse, err error) {
	req.InitWithApiInfo("Yundun", "2015-04-16", "WafInfo", "", "")
	resp = &WafInfoResponse{}
	err = client.DoAction(req, resp)
	return
}

type WafInfoResponse struct {
	responses.BaseResponse
	RequestId    common.String
	WafDomainNum common.Integer
	WafInfos     WafInfoWafInfoList
}

type WafInfoWafInfo struct {
	Id     common.Integer
	Domain common.String
	Cname  common.String
	Status common.Integer
}

type WafInfoWafInfoList []WafInfoWafInfo

func (list *WafInfoWafInfoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]WafInfoWafInfo)
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
