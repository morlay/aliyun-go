package cbn

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type DeleteCenRequest struct {
	requests.RpcRequest
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	CenId                string `position:"Query" name:"CenId"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (req *DeleteCenRequest) Invoke(client *sdk.Client) (resp *DeleteCenResponse, err error) {
	req.InitWithApiInfo("Cbn", "2017-09-12", "DeleteCen", "cbn", "")
	resp = &DeleteCenResponse{}
	err = client.DoAction(req, resp)
	return
}

type DeleteCenResponse struct {
	responses.BaseResponse
	RequestId common.String
}
