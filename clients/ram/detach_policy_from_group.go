package ram

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DetachPolicyFromGroupRequest struct {
	PolicyType string `position:"Query" name:"PolicyType"`
	PolicyName string `position:"Query" name:"PolicyName"`
	GroupName  string `position:"Query" name:"GroupName"`
}

func (r DetachPolicyFromGroupRequest) Invoke(client *sdk.Client) (response *DetachPolicyFromGroupResponse, err error) {
	req := struct {
		*requests.RpcRequest
		DetachPolicyFromGroupRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Ram", "2015-05-01", "DetachPolicyFromGroup", "", "")

	resp := struct {
		*responses.BaseResponse
		DetachPolicyFromGroupResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.DetachPolicyFromGroupResponse

	err = client.DoAction(&req, &resp)
	return
}

type DetachPolicyFromGroupResponse struct {
	RequestId string
}
