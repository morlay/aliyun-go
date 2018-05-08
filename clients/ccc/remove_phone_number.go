package ccc

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type RemovePhoneNumberRequest struct {
	requests.RpcRequest
	InstanceId    string `position:"Query" name:"InstanceId"`
	PhoneNumberId string `position:"Query" name:"PhoneNumberId"`
}

func (req *RemovePhoneNumberRequest) Invoke(client *sdk.Client) (resp *RemovePhoneNumberResponse, err error) {
	req.InitWithApiInfo("CCC", "2017-07-05", "RemovePhoneNumber", "ccc", "")
	resp = &RemovePhoneNumberResponse{}
	err = client.DoAction(req, resp)
	return
}

type RemovePhoneNumberResponse struct {
	responses.BaseResponse
	RequestId      common.String
	Success        bool
	Code           common.String
	Message        common.String
	HttpStatusCode common.Integer
}
