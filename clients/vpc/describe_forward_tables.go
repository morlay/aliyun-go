package vpc

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type DescribeForwardTablesRequest struct {
	requests.RpcRequest
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	ForwardTableId       string `position:"Query" name:"ForwardTableId"`
	PageSize             int    `position:"Query" name:"PageSize"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	PageNumber           int    `position:"Query" name:"PageNumber"`
}

func (req *DescribeForwardTablesRequest) Invoke(client *sdk.Client) (resp *DescribeForwardTablesResponse, err error) {
	req.InitWithApiInfo("Vpc", "2016-04-28", "DescribeForwardTables", "vpc", "")
	resp = &DescribeForwardTablesResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeForwardTablesResponse struct {
	responses.BaseResponse
	RequestId     common.String
	TotalCount    common.Integer
	PageNumber    common.Integer
	PageSize      common.Integer
	ForwardTables DescribeForwardTablesForwardTableList
}

type DescribeForwardTablesForwardTable struct {
	NatGatewayId   common.String
	ForwardTableId common.String
	CreationTime   common.String
	ForwardEntrys  DescribeForwardTablesForwardEntryList
}

type DescribeForwardTablesForwardEntry struct {
	ForwardTableId common.String
	ForwardEntryId common.String
	ExternalIp     common.String
	ExternalPort   common.String
	IpProtocol     common.String
	InternalIp     common.String
	InternalPort   common.String
	Status         common.String
}

type DescribeForwardTablesForwardTableList []DescribeForwardTablesForwardTable

func (list *DescribeForwardTablesForwardTableList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeForwardTablesForwardTable)
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

type DescribeForwardTablesForwardEntryList []DescribeForwardTablesForwardEntry

func (list *DescribeForwardTablesForwardEntryList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeForwardTablesForwardEntry)
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
