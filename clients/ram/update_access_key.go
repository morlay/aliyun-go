package ram

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type UpdateAccessKeyRequest struct {
	UserAccessKeyId string `position:"Query" name:"UserAccessKeyId"`
	UserName        string `position:"Query" name:"UserName"`
	Status          string `position:"Query" name:"Status"`
}

func (r UpdateAccessKeyRequest) Invoke(client *sdk.Client) (response *UpdateAccessKeyResponse, err error) {
	req := struct {
		*requests.RpcRequest
		UpdateAccessKeyRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Ram", "2015-05-01", "UpdateAccessKey", "", "")

	resp := struct {
		*responses.BaseResponse
		UpdateAccessKeyResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.UpdateAccessKeyResponse

	err = client.DoAction(&req, &resp)
	return
}

type UpdateAccessKeyResponse struct {
	RequestId string
}
