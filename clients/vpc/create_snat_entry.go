package vpc

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type CreateSnatEntryRequest struct {
	requests.RpcRequest
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	SourceVSwitchId      string `position:"Query" name:"SourceVSwitchId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	SourceCIDR           string `position:"Query" name:"SourceCIDR"`
	SnatTableId          string `position:"Query" name:"SnatTableId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	SnatIp               string `position:"Query" name:"SnatIp"`
}

func (req *CreateSnatEntryRequest) Invoke(client *sdk.Client) (resp *CreateSnatEntryResponse, err error) {
	req.InitWithApiInfo("Vpc", "2016-04-28", "CreateSnatEntry", "vpc", "")
	resp = &CreateSnatEntryResponse{}
	err = client.DoAction(req, resp)
	return
}

type CreateSnatEntryResponse struct {
	responses.BaseResponse
	RequestId   string
	SnatEntryId string
}
