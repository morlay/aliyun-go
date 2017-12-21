package market

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeInstanceRequest struct {
	InstanceId string `position:"Query" name:"InstanceId"`
}

func (r DescribeInstanceRequest) Invoke(client *sdk.Client) (response *DescribeInstanceResponse, err error) {
	req := struct {
		*requests.RpcRequest
		DescribeInstanceRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Market", "2015-11-01", "DescribeInstance", "", "")

	resp := struct {
		*responses.BaseResponse
		DescribeInstanceResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.DescribeInstanceResponse

	err = client.DoAction(&req, &resp)
	return
}

type DescribeInstanceResponse struct {
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
