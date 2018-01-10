package cms

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type CreateNotifyPolicyRequest struct {
	requests.RpcRequest
	PolicyType string `position:"Query" name:"PolicyType"`
	AlertName  string `position:"Query" name:"AlertName"`
	EndTime    int64  `position:"Query" name:"EndTime"`
	StartTime  int64  `position:"Query" name:"StartTime"`
	Dimensions string `position:"Query" name:"Dimensions"`
}

func (req *CreateNotifyPolicyRequest) Invoke(client *sdk.Client) (resp *CreateNotifyPolicyResponse, err error) {
	req.InitWithApiInfo("Cms", "2017-03-01", "CreateNotifyPolicy", "cms", "")
	resp = &CreateNotifyPolicyResponse{}
	err = client.DoAction(req, resp)
	return
}

type CreateNotifyPolicyResponse struct {
	responses.BaseResponse
	Code    string
	Message string
	Success string
	TraceId string
	Result  int
}
