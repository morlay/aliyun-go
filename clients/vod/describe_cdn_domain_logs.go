package vod

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type DescribeCdnDomainLogsRequest struct {
	requests.RpcRequest
	ResourceOwnerId      string `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	PageNo               int64  `position:"Query" name:"PageNo"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	DomainName           string `position:"Query" name:"DomainName"`
	PageSize             int64  `position:"Query" name:"PageSize"`
	EndTime              string `position:"Query" name:"EndTime"`
	StartTime            string `position:"Query" name:"StartTime"`
	OwnerId              string `position:"Query" name:"OwnerId"`
	LogDay               string `position:"Query" name:"LogDay"`
}

func (req *DescribeCdnDomainLogsRequest) Invoke(client *sdk.Client) (resp *DescribeCdnDomainLogsResponse, err error) {
	req.InitWithApiInfo("vod", "2017-03-21", "DescribeCdnDomainLogs", "vod", "")
	resp = &DescribeCdnDomainLogsResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeCdnDomainLogsResponse struct {
	responses.BaseResponse
	RequestId      common.String
	PageNo         common.Long
	PageSize       common.Long
	Total          common.Long
	DomainLogModel DescribeCdnDomainLogsDomainLogModel
}

type DescribeCdnDomainLogsDomainLogModel struct {
	DomainName       common.String
	DomainLogDetails DescribeCdnDomainLogsDomainLogDetailList
}

type DescribeCdnDomainLogsDomainLogDetail struct {
	LogPath   common.String
	StartTime common.String
	EndTime   common.String
	LogSize   common.Long
	LogName   common.String
}

type DescribeCdnDomainLogsDomainLogDetailList []DescribeCdnDomainLogsDomainLogDetail

func (list *DescribeCdnDomainLogsDomainLogDetailList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeCdnDomainLogsDomainLogDetail)
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
