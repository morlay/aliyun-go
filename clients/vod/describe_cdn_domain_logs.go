package vod

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeCdnDomainLogsRequest struct {
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

func (r DescribeCdnDomainLogsRequest) Invoke(client *sdk.Client) (response *DescribeCdnDomainLogsResponse, err error) {
	req := struct {
		*requests.RpcRequest
		DescribeCdnDomainLogsRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("vod", "2017-03-21", "DescribeCdnDomainLogs", "", "")

	resp := struct {
		*responses.BaseResponse
		DescribeCdnDomainLogsResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.DescribeCdnDomainLogsResponse

	err = client.DoAction(&req, &resp)
	return
}

type DescribeCdnDomainLogsResponse struct {
	RequestId      string
	PageNo         int64
	PageSize       int64
	Total          int64
	DomainLogModel DescribeCdnDomainLogsDomainLogModel
}

type DescribeCdnDomainLogsDomainLogModel struct {
	DomainName       string
	DomainLogDetails DescribeCdnDomainLogsDomainLogDetailList
}

type DescribeCdnDomainLogsDomainLogDetail struct {
	LogPath   string
	StartTime string
	EndTime   string
	LogSize   int64
	LogName   string
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
