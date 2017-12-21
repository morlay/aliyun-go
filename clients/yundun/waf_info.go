package yundun

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type WafInfoRequest struct {
	InstanceId   string `position:"Query" name:"InstanceId"`
	InstanceType string `position:"Query" name:"InstanceType"`
}

func (r WafInfoRequest) Invoke(client *sdk.Client) (response *WafInfoResponse, err error) {
	req := struct {
		*requests.RpcRequest
		WafInfoRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Yundun", "2015-04-16", "WafInfo", "", "")

	resp := struct {
		*responses.BaseResponse
		WafInfoResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.WafInfoResponse

	err = client.DoAction(&req, &resp)
	return
}

type WafInfoResponse struct {
	RequestId    string
	WafDomainNum int
	WafInfos     WafInfoWafInfoList
}

type WafInfoWafInfo struct {
	Id     int
	Domain string
	Cname  string
	Status int
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
