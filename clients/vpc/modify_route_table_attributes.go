package vpc

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ModifyRouteTableAttributesRequest struct {
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

func (r ModifyRouteTableAttributesRequest) Invoke(client *sdk.Client) (response *ModifyRouteTableAttributesResponse, err error) {
	req := struct {
		*requests.RpcRequest
		ModifyRouteTableAttributesRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Vpc", "2016-04-28", "ModifyRouteTableAttributes", "vpc", "")

	resp := struct {
		*responses.BaseResponse
		ModifyRouteTableAttributesResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.ModifyRouteTableAttributesResponse

	err = client.DoAction(&req, &resp)
	return
}

type ModifyRouteTableAttributesResponse struct {
	RequestId string
	Code      string
	Message   string
	Success   bool
}
