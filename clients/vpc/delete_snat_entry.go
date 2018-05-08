package vpc

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type DeleteSnatEntryRequest struct {
	requests.RpcRequest
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	SnatTableId          string `position:"Query" name:"SnatTableId"`
	SnatEntryId          string `position:"Query" name:"SnatEntryId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (req *DeleteSnatEntryRequest) Invoke(client *sdk.Client) (resp *DeleteSnatEntryResponse, err error) {
	req.InitWithApiInfo("Vpc", "2016-04-28", "DeleteSnatEntry", "vpc", "")
	resp = &DeleteSnatEntryResponse{}
	err = client.DoAction(req, resp)
	return
}

type DeleteSnatEntryResponse struct {
	responses.BaseResponse
	RequestId common.String
}
