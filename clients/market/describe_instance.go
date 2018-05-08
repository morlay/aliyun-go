package market

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type DescribeInstanceRequest struct {
	requests.RpcRequest
	InstanceId string `position:"Query" name:"InstanceId"`
}

func (req *DescribeInstanceRequest) Invoke(client *sdk.Client) (resp *DescribeInstanceResponse, err error) {
	req.InitWithApiInfo("Market", "2015-11-01", "DescribeInstance", "yunmarket", "")
	resp = &DescribeInstanceResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeInstanceResponse struct {
	responses.BaseResponse
	InstanceId     common.Long
	OrderId        common.Long
	SupplierName   common.String
	ProductCode    common.String
	ProductSkuCode common.String
	ProductName    common.String
	ProductType    common.String
	Status         common.String
	BeganOn        common.Long
	EndOn          common.Long
	CreatedOn      common.Long
	ExtendJson     common.String
	HostJson       common.String
	AppJson        common.String
}
