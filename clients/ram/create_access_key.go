package ram

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type CreateAccessKeyRequest struct {
	requests.RpcRequest
	UserName string `position:"Query" name:"UserName"`
}

func (req *CreateAccessKeyRequest) Invoke(client *sdk.Client) (resp *CreateAccessKeyResponse, err error) {
	req.InitWithApiInfo("Ram", "2015-05-01", "CreateAccessKey", "", "")
	resp = &CreateAccessKeyResponse{}
	err = client.DoAction(req, resp)
	return
}

type CreateAccessKeyResponse struct {
	responses.BaseResponse
	RequestId string
	AccessKey CreateAccessKeyAccessKey
}

type CreateAccessKeyAccessKey struct {
	AccessKeyId     string
	AccessKeySecret string
	Status          string
	CreateDate      string
}
