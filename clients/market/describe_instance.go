package market

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
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
	InstanceId     int64
	OrderId        int64
	SupplierName   string
	ProductCode    string
	ProductSkuCode string
	ProductName    string
	ProductType    string
	Status         string
	BeganOn        int64
	EndOn          int64
	CreatedOn      int64
	ExtendJson     string
	HostJson       string
	AppJson        string
}
