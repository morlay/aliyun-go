package alidns

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeDomainGroupsRequest struct {
	Lang         string `position:"Query" name:"Lang"`
	UserClientIp string `position:"Query" name:"UserClientIp"`
	KeyWord      string `position:"Query" name:"KeyWord"`
	PageNumber   int64  `position:"Query" name:"PageNumber"`
	PageSize     int64  `position:"Query" name:"PageSize"`
}

func (r DescribeDomainGroupsRequest) Invoke(client *sdk.Client) (response *DescribeDomainGroupsResponse, err error) {
	req := struct {
		*requests.RpcRequest
		DescribeDomainGroupsRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Alidns", "2015-01-09", "DescribeDomainGroups", "", "")

	resp := struct {
		*responses.BaseResponse
		DescribeDomainGroupsResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.DescribeDomainGroupsResponse

	err = client.DoAction(&req, &resp)
	return
}

type DescribeDomainGroupsResponse struct {
	RequestId    string
	TotalCount   int64
	PageNumber   int64
	PageSize     int64
	DomainGroups DescribeDomainGroupsDomainGroupList
}

type DescribeDomainGroupsDomainGroup struct {
	GroupId   string
	GroupName string
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
