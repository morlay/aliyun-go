package vpc

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type ModifyRouteTableAttributesRequest struct {
	requests.RpcRequest
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	Bandwidth            string `position:"Query" name:"Bandwidth"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	Description          string `position:"Query" name:"Description"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	KbpsBandwidth        string `position:"Query" name:"KbpsBandwidth"`
	RouteTableName       string `position:"Query" name:"RouteTableName"`
	ResourceUid          int64  `position:"Query" name:"ResourceUid"`
	ResourceBid          string `position:"Query" name:"ResourceBid"`
	RouteTableId         string `position:"Query" name:"RouteTableId"`
}

func (req *ModifyRouteTableAttributesRequest) Invoke(client *sdk.Client) (resp *ModifyRouteTableAttributesResponse, err error) {
	req.InitWithApiInfo("Vpc", "2016-04-28", "ModifyRouteTableAttributes", "vpc", "")
	resp = &ModifyRouteTableAttributesResponse{}
	err = client.DoAction(req, resp)
	return
}

type ModifyRouteTableAttributesResponse struct {
	responses.BaseResponse
	RequestId common.String
	Code      common.String
	Message   common.String
	Success   bool
}
