package cbn

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeCenRegionDomainRouteEntriesRequest struct {
	requests.RpcRequest
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	CenId                string `position:"Query" name:"CenId"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	CenRegionId          string `position:"Query" name:"CenRegionId"`
	PageSize             int    `position:"Query" name:"PageSize"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	PageNumber           int    `position:"Query" name:"PageNumber"`
}

func (req *DescribeCenRegionDomainRouteEntriesRequest) Invoke(client *sdk.Client) (resp *DescribeCenRegionDomainRouteEntriesResponse, err error) {
	req.InitWithApiInfo("Cbn", "2017-09-12", "DescribeCenRegionDomainRouteEntries", "cbn", "")
	resp = &DescribeCenRegionDomainRouteEntriesResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeCenRegionDomainRouteEntriesResponse struct {
	responses.BaseResponse
	RequestId       string
	PageNumber      int
	TotalCount      int
	PageSize        int
	CenRouteEntries DescribeCenRegionDomainRouteEntriesCenRouteEntryList
}

type DescribeCenRegionDomainRouteEntriesCenRouteEntry struct {
	DestinationCidrBlock string
	Type                 string
	NextHopInstanceId    string
	NextHopType          string
	NextHopRegionId      string
}

type DescribeCenRegionDomainRouteEntriesCenRouteEntryList []DescribeCenRegionDomainRouteEntriesCenRouteEntry

func (list *DescribeCenRegionDomainRouteEntriesCenRouteEntryList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeCenRegionDomainRouteEntriesCenRouteEntry)
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
