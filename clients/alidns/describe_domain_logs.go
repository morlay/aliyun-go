package alidns

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeDomainLogsRequest struct {
	Lang         string `position:"Query" name:"Lang"`
	UserClientIp string `position:"Query" name:"UserClientIp"`
	KeyWord      string `position:"Query" name:"KeyWord"`
	GroupId      string `position:"Query" name:"GroupId"`
	PageNumber   int64  `position:"Query" name:"PageNumber"`
	PageSize     int64  `position:"Query" name:"PageSize"`
}

func (r DescribeDomainLogsRequest) Invoke(client *sdk.Client) (response *DescribeDomainLogsResponse, err error) {
	req := struct {
		*requests.RpcRequest
		DescribeDomainLogsRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Alidns", "2015-01-09", "DescribeDomainLogs", "", "")

	resp := struct {
		*responses.BaseResponse
		DescribeDomainLogsResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.DescribeDomainLogsResponse

	err = client.DoAction(&req, &resp)
	return
}

type DescribeDomainLogsResponse struct {
	RequestId  string
	TotalCount int64
	PageNumber int64
	PageSize   int64
	DomainLogs DescribeDomainLogsDomainLogList
}

type DescribeDomainLogsDomainLog struct {
	ActionTime string
	DomainName string
	Action     string
	Message    string
	ClientIp   string
}

type DescribeDomainLogsDomainLogList []DescribeDomainLogsDomainLog

func (list *DescribeDomainLogsDomainLogList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeDomainLogsDomainLog)
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
