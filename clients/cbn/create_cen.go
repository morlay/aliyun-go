package cbn

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type CreateCenRequest struct {
	requests.RpcRequest
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	ClientToken          string `position:"Query" name:"ClientToken"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	Name                 string `position:"Query" name:"Name"`
	Description          string `position:"Query" name:"Description"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (req *CreateCenRequest) Invoke(client *sdk.Client) (resp *CreateCenResponse, err error) {
	req.InitWithApiInfo("Cbn", "2017-09-12", "CreateCen", "cbn", "")
	resp = &CreateCenResponse{}
	err = client.DoAction(req, resp)
	return
}

type CreateCenResponse struct {
	responses.BaseResponse
	RequestId common.String
	CenId     common.String
}
