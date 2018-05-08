package vpc

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type DeleteForwardEntryRequest struct {
	requests.RpcRequest
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	ForwardEntryId       string `position:"Query" name:"ForwardEntryId"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	ForwardTableId       string `position:"Query" name:"ForwardTableId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (req *DeleteForwardEntryRequest) Invoke(client *sdk.Client) (resp *DeleteForwardEntryResponse, err error) {
	req.InitWithApiInfo("Vpc", "2016-04-28", "DeleteForwardEntry", "vpc", "")
	resp = &DeleteForwardEntryResponse{}
	err = client.DoAction(req, resp)
	return
}

type DeleteForwardEntryResponse struct {
	responses.BaseResponse
	RequestId common.String
}
