package ram

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type DeletePublicKeyRequest struct {
	requests.RpcRequest
	UserPublicKeyId string `position:"Query" name:"UserPublicKeyId"`
	UserName        string `position:"Query" name:"UserName"`
}

func (req *DeletePublicKeyRequest) Invoke(client *sdk.Client) (resp *DeletePublicKeyResponse, err error) {
	req.InitWithApiInfo("Ram", "2015-05-01", "DeletePublicKey", "", "")
	resp = &DeletePublicKeyResponse{}
	err = client.DoAction(req, resp)
	return
}

type DeletePublicKeyResponse struct {
	responses.BaseResponse
	RequestId common.String
}
