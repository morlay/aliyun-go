package cms

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type DeleteNotifyPolicyRequest struct {
	requests.RpcRequest
	PolicyType string `position:"Query" name:"PolicyType"`
	AlertName  string `position:"Query" name:"AlertName"`
	Id         string `position:"Query" name:"Id"`
	Dimensions string `position:"Query" name:"Dimensions"`
}

func (req *DeleteNotifyPolicyRequest) Invoke(client *sdk.Client) (resp *DeleteNotifyPolicyResponse, err error) {
	req.InitWithApiInfo("Cms", "2018-03-08", "DeleteNotifyPolicy", "cms", "")
	resp = &DeleteNotifyPolicyResponse{}
	err = client.DoAction(req, resp)
	return
}

type DeleteNotifyPolicyResponse struct {
	responses.BaseResponse
	Code    common.String
	Message common.String
	Success common.String
	TraceId common.String
	Result  common.Integer
}
