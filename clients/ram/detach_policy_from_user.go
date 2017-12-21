package ram

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DetachPolicyFromUserRequest struct {
	PolicyType string `position:"Query" name:"PolicyType"`
	PolicyName string `position:"Query" name:"PolicyName"`
	UserName   string `position:"Query" name:"UserName"`
}

func (r DetachPolicyFromUserRequest) Invoke(client *sdk.Client) (response *DetachPolicyFromUserResponse, err error) {
	req := struct {
		*requests.RpcRequest
		DetachPolicyFromUserRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Ram", "2015-05-01", "DetachPolicyFromUser", "", "")

	resp := struct {
		*responses.BaseResponse
		DetachPolicyFromUserResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.DetachPolicyFromUserResponse

	err = client.DoAction(&req, &resp)
	return
}

type DetachPolicyFromUserResponse struct {
	RequestId string
}
