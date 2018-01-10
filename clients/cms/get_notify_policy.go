package cms

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type GetNotifyPolicyRequest struct {
	requests.RpcRequest
	PolicyType string `position:"Query" name:"PolicyType"`
	AlertName  string `position:"Query" name:"AlertName"`
	Id         string `position:"Query" name:"Id"`
	Dimensions string `position:"Query" name:"Dimensions"`
}

func (req *GetNotifyPolicyRequest) Invoke(client *sdk.Client) (resp *GetNotifyPolicyResponse, err error) {
	req.InitWithApiInfo("Cms", "2017-03-01", "GetNotifyPolicy", "cms", "")
	resp = &GetNotifyPolicyResponse{}
	err = client.DoAction(req, resp)
	return
}

type GetNotifyPolicyResponse struct {
	responses.BaseResponse
	Code    string
	Message string
	Success string
	TraceId string
	Result  GetNotifyPolicyResult
}

type GetNotifyPolicyResult struct {
	AlertName  string
	Dimensions string
	Type       string
	Id         string
	StartTime  int64
	EndTime    int64
}
