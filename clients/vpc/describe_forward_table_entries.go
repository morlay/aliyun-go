package vpc

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type DescribeForwardTableEntriesRequest struct {
	requests.RpcRequest
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	ForwardEntryId       string `position:"Query" name:"ForwardEntryId"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	ForwardTableId       string `position:"Query" name:"ForwardTableId"`
	PageSize             int    `position:"Query" name:"PageSize"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	PageNumber           int    `position:"Query" name:"PageNumber"`
}

func (req *DescribeForwardTableEntriesRequest) Invoke(client *sdk.Client) (resp *DescribeForwardTableEntriesResponse, err error) {
	req.InitWithApiInfo("Vpc", "2016-04-28", "DescribeForwardTableEntries", "vpc", "")
	resp = &DescribeForwardTableEntriesResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeForwardTableEntriesResponse struct {
	responses.BaseResponse
	RequestId           common.String
	TotalCount          common.Integer
	PageNumber          common.Integer
	PageSize            common.Integer
	ForwardTableEntries DescribeForwardTableEntriesForwardTableEntryList
}

type DescribeForwardTableEntriesForwardTableEntry struct {
	ForwardTableId common.String
	ForwardEntryId common.String
	ExternalIp     common.String
	ExternalPort   common.String
	IpProtocol     common.String
	InternalIp     common.String
	InternalPort   common.String
	Status         common.String
}

type DescribeForwardTableEntriesForwardTableEntryList []DescribeForwardTableEntriesForwardTableEntry

func (list *DescribeForwardTableEntriesForwardTableEntryList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeForwardTableEntriesForwardTableEntry)
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
