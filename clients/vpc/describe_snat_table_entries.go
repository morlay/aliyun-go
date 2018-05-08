package vpc

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
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
	RequestId        common.String
	TotalCount       common.Integer
	PageNumber       common.Integer
	PageSize         common.Integer
	SnatTableEntries DescribeSnatTableEntriesSnatTableEntryList
}

type DescribeSnatTableEntriesSnatTableEntry struct {
	SnatTableId     common.String
	SnatEntryId     common.String
	SourceVSwitchId common.String
	SourceCIDR      common.String
	SnatIp          common.String
	Status          common.String
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
