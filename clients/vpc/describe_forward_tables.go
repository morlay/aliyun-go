package vpc

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
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
	RequestId     string
	TotalCount    int
	PageNumber    int
	PageSize      int
	ForwardTables DescribeForwardTablesForwardTableList
}

type DescribeForwardTablesForwardTable struct {
	NatGatewayId   string
	ForwardTableId string
	CreationTime   string
	ForwardEntrys  DescribeForwardTablesForwardEntryList
}

type DescribeForwardTablesForwardEntry struct {
	ForwardTableId string
	ForwardEntryId string
	ExternalIp     string
	ExternalPort   string
	IpProtocol     string
	InternalIp     string
	InternalPort   string
	Status         string
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
