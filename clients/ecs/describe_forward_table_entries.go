package ecs

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
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
	req.InitWithApiInfo("Ecs", "2014-05-26", "DescribeForwardTableEntries", "ecs", "")
	resp = &DescribeForwardTableEntriesResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeForwardTableEntriesResponse struct {
	responses.BaseResponse
	RequestId           string
	TotalCount          int
	PageNumber          int
	PageSize            int
	ForwardTableEntries DescribeForwardTableEntriesForwardTableEntryList
}

type DescribeForwardTableEntriesForwardTableEntry struct {
	ForwardTableId string
	ForwardEntryId string
	ExternalIp     string
	ExternalPort   string
	IpProtocol     string
	InternalIp     string
	InternalPort   string
	Status         string
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
