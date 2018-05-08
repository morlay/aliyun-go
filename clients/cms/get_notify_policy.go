package cms

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type GetNotifyPolicyRequest struct {
	requests.RpcRequest
	PolicyType string `position:"Query" name:"PolicyType"`
	AlertName  string `position:"Query" name:"AlertName"`
	Id         string `position:"Query" name:"Id"`
	Dimensions string `position:"Query" name:"Dimensions"`
}

func (req *GetNotifyPolicyRequest) Invoke(client *sdk.Client) (resp *GetNotifyPolicyResponse, err error) {
	req.InitWithApiInfo("Cms", "2018-03-08", "GetNotifyPolicy", "cms", "")
	resp = &GetNotifyPolicyResponse{}
	err = client.DoAction(req, resp)
	return
}

type GetNotifyPolicyResponse struct {
	responses.BaseResponse
	Code    common.String
	Message common.String
	Success common.String
	TraceId common.String
	Result  GetNotifyPolicyResult
}

type GetNotifyPolicyResult struct {
	AlertName  common.String
	Dimensions common.String
	Type       common.String
	Id         common.String
	StartTime  common.Long
	EndTime    common.Long
}
