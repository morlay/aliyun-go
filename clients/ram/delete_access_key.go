package ram

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DeleteAccessKeyRequest struct {
	requests.RpcRequest
	UserAccessKeyId string `position:"Query" name:"UserAccessKeyId"`
	UserName        string `position:"Query" name:"UserName"`
}

func (req *DeleteAccessKeyRequest) Invoke(client *sdk.Client) (resp *DeleteAccessKeyResponse, err error) {
	req.InitWithApiInfo("Ram", "2015-05-01", "DeleteAccessKey", "", "")
	resp = &DeleteAccessKeyResponse{}
	err = client.DoAction(req, resp)
	return
}

type DeleteAccessKeyResponse struct {
	responses.BaseResponse
	RequestId string
}
