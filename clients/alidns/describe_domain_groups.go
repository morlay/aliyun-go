package alidns

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type DescribeDomainGroupsRequest struct {
	requests.RpcRequest
	Lang         string `position:"Query" name:"Lang"`
	UserClientIp string `position:"Query" name:"UserClientIp"`
	KeyWord      string `position:"Query" name:"KeyWord"`
	PageNumber   int64  `position:"Query" name:"PageNumber"`
	PageSize     int64  `position:"Query" name:"PageSize"`
}

func (req *DescribeDomainGroupsRequest) Invoke(client *sdk.Client) (resp *DescribeDomainGroupsResponse, err error) {
	req.InitWithApiInfo("Alidns", "2015-01-09", "DescribeDomainGroups", "", "")
	resp = &DescribeDomainGroupsResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeDomainGroupsResponse struct {
	responses.BaseResponse
	RequestId    common.String
	TotalCount   common.Long
	PageNumber   common.Long
	PageSize     common.Long
	DomainGroups DescribeDomainGroupsDomainGroupList
}

type DescribeDomainGroupsDomainGroup struct {
	GroupId   common.String
	GroupName common.String
}

type DescribeDomainGroupsDomainGroupList []DescribeDomainGroupsDomainGroup

func (list *DescribeDomainGroupsDomainGroupList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeDomainGroupsDomainGroup)
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
