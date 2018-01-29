package mts

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type CreateSessionRequest struct {
	requests.RpcRequest
	ResourceOwnerId      string `position:"Query" name:"ResourceOwnerId"`
	SessionTime          int    `position:"Query" name:"SessionTime"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	EndUserId            string `position:"Query" name:"EndUserId"`
	OwnerId              string `position:"Query" name:"OwnerId"`
}

func (req *CreateSessionRequest) Invoke(client *sdk.Client) (resp *CreateSessionResponse, err error) {
	req.InitWithApiInfo("Mts", "2014-06-18", "CreateSession", "mts", "")
	resp = &CreateSessionResponse{}
	err = client.DoAction(req, resp)
	return
}

type CreateSessionResponse struct {
	responses.BaseResponse
	RequestId string
	Session   string
}
