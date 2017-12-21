package ram

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DeleteAccessKeyRequest struct {
	UserAccessKeyId string `position:"Query" name:"UserAccessKeyId"`
	UserName        string `position:"Query" name:"UserName"`
}

func (r DeleteAccessKeyRequest) Invoke(client *sdk.Client) (response *DeleteAccessKeyResponse, err error) {
	req := struct {
		*requests.RpcRequest
		DeleteAccessKeyRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Ram", "2015-05-01", "DeleteAccessKey", "", "")

	resp := struct {
		*responses.BaseResponse
		DeleteAccessKeyResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.DeleteAccessKeyResponse

	err = client.DoAction(&req, &resp)
	return
}

type DeleteAccessKeyResponse struct {
	RequestId string
}
