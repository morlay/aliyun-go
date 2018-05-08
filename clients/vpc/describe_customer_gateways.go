package vpc

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type DescribeCustomerGatewaysRequest struct {
	requests.RpcRequest
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	PageSize             int    `position:"Query" name:"PageSize"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	CustomerGatewayId    string `position:"Query" name:"CustomerGatewayId"`
	PageNumber           int    `position:"Query" name:"PageNumber"`
}

func (req *DescribeCustomerGatewaysRequest) Invoke(client *sdk.Client) (resp *DescribeCustomerGatewaysResponse, err error) {
	req.InitWithApiInfo("Vpc", "2016-04-28", "DescribeCustomerGateways", "vpc", "")
	resp = &DescribeCustomerGatewaysResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeCustomerGatewaysResponse struct {
	responses.BaseResponse
	RequestId        common.String
	TotalCount       common.Integer
	PageNumber       common.Integer
	PageSize         common.Integer
	CustomerGateways DescribeCustomerGatewaysCustomerGatewayList
}

type DescribeCustomerGatewaysCustomerGateway struct {
	CustomerGatewayId common.String
	Name              common.String
	IpAddress         common.String
	Description       common.String
	CreateTime        common.Long
}

type DescribeCustomerGatewaysCustomerGatewayList []DescribeCustomerGatewaysCustomerGateway

func (list *DescribeCustomerGatewaysCustomerGatewayList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeCustomerGatewaysCustomerGateway)
	err := json.Unmarshal(data, &m)
	if err != nil {
		return err
	}
	for _, v := range m {
		*list = v
		break
	}
	return nil
}
