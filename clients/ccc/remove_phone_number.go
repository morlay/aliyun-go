package ccc

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type RemovePhoneNumberRequest struct {
	InstanceId    string `position:"Query" name:"InstanceId"`
	PhoneNumberId string `position:"Query" name:"PhoneNumberId"`
}

func (r RemovePhoneNumberRequest) Invoke(client *sdk.Client) (response *RemovePhoneNumberResponse, err error) {
	req := struct {
		*requests.RpcRequest
		RemovePhoneNumberRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("CCC", "2017-07-05", "RemovePhoneNumber", "ccc", "")

	resp := struct {
		*responses.BaseResponse
		RemovePhoneNumberResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.RemovePhoneNumberResponse

	err = client.DoAction(&req, &resp)
	return
}

type RemovePhoneNumberResponse struct {
	RequestId      string
	Success        bool
	Code           string
	Message        string
	HttpStatusCode int
}
