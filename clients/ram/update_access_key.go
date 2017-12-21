package ram

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type UpdateAccessKeyRequest struct {
	requests.RpcRequest
	UserAccessKeyId string `position:"Query" name:"UserAccessKeyId"`
	UserName        string `position:"Query" name:"UserName"`
	Status          string `position:"Query" name:"Status"`
}

func (req *UpdateAccessKeyRequest) Invoke(client *sdk.Client) (resp *UpdateAccessKeyResponse, err error) {
	req.InitWithApiInfo("Ram", "2015-05-01", "UpdateAccessKey", "", "")
	resp = &UpdateAccessKeyResponse{}
	err = client.DoAction(req, resp)
	return
}

type UpdateAccessKeyResponse struct {
	responses.BaseResponse
	RequestId string
}
