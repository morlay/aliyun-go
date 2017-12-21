package cdn

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeCdnDomainLogsRequest struct {
	requests.RpcRequest
	SecurityToken string `position:"Query" name:"SecurityToken"`
	DomainName    string `position:"Query" name:"DomainName"`
	PageSize      int64  `position:"Query" name:"PageSize"`
	EndTime       string `position:"Query" name:"EndTime"`
	StartTime     string `position:"Query" name:"StartTime"`
	OwnerId       int64  `position:"Query" name:"OwnerId"`
	PageNumber    int64  `position:"Query" name:"PageNumber"`
	LogDay        string `position:"Query" name:"LogDay"`
}

func (req *DescribeCdnDomainLogsRequest) Invoke(client *sdk.Client) (resp *DescribeCdnDomainLogsResponse, err error) {
	req.InitWithApiInfo("Cdn", "2014-11-11", "DescribeCdnDomainLogs", "", "")
	resp = &DescribeCdnDomainLogsResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeCdnDomainLogsResponse struct {
	responses.BaseResponse
	RequestId      string
	PageNumber     int64
	PageSize       int64
	TotalCount     int64
	DomainLogModel DescribeCdnDomainLogsDomainLogModel
}

type DescribeCdnDomainLogsDomainLogModel struct {
	DomainName       string
	DomainLogDetails DescribeCdnDomainLogsDomainLogDetailList
}

type DescribeCdnDomainLogsDomainLogDetail struct {
	LogName   string
	LogPath   string
	LogSize   int64
	StartTime string
	EndTime   string
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
