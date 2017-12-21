package vpc

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeSnatTableEntriesRequest struct {
	requests.RpcRequest
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	PageSize             int    `position:"Query" name:"PageSize"`
	SnatTableId          string `position:"Query" name:"SnatTableId"`
	SnatEntryId          string `position:"Query" name:"SnatEntryId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	PageNumber           int    `position:"Query" name:"PageNumber"`
}

func (req *DescribeSnatTableEntriesRequest) Invoke(client *sdk.Client) (resp *DescribeSnatTableEntriesResponse, err error) {
	req.InitWithApiInfo("Vpc", "2016-04-28", "DescribeSnatTableEntries", "vpc", "")
	resp = &DescribeSnatTableEntriesResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeSnatTableEntriesResponse struct {
	responses.BaseResponse
	RequestId        string
	TotalCount       int
	PageNumber       int
	PageSize         int
	SnatTableEntries DescribeSnatTableEntriesSnatTableEntryList
}

type DescribeSnatTableEntriesSnatTableEntry struct {
	SnatTableId     string
	SnatEntryId     string
	SourceVSwitchId string
	SourceCIDR      string
	SnatIp          string
	Status          string
}

type DescribeSnatTableEntriesSnatTableEntryList []DescribeSnatTableEntriesSnatTableEntry

func (list *DescribeSnatTableEntriesSnatTableEntryList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeSnatTableEntriesSnatTableEntry)
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
