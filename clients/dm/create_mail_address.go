package dm

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type CreateMailAddressRequest struct {
	requests.RpcRequest
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	AccountName          string `position:"Query" name:"AccountName"`
	ReplyAddress         string `position:"Query" name:"ReplyAddress"`
	Sendtype             string `position:"Query" name:"Sendtype"`
}

func (req *CreateMailAddressRequest) Invoke(client *sdk.Client) (resp *CreateMailAddressResponse, err error) {
	req.InitWithApiInfo("Dm", "2015-11-23", "CreateMailAddress", "", "")
	resp = &CreateMailAddressResponse{}
	err = client.DoAction(req, resp)
	return
}

type CreateMailAddressResponse struct {
	responses.BaseResponse
	RequestId string
}
