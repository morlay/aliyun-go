package alidns

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type DescribeDomainLogsRequest struct {
	requests.RpcRequest
	Lang         string `position:"Query" name:"Lang"`
	UserClientIp string `position:"Query" name:"UserClientIp"`
	KeyWord      string `position:"Query" name:"KeyWord"`
	GroupId      string `position:"Query" name:"GroupId"`
	PageNumber   int64  `position:"Query" name:"PageNumber"`
	PageSize     int64  `position:"Query" name:"PageSize"`
}

func (req *DescribeDomainLogsRequest) Invoke(client *sdk.Client) (resp *DescribeDomainLogsResponse, err error) {
	req.InitWithApiInfo("Alidns", "2015-01-09", "DescribeDomainLogs", "", "")
	resp = &DescribeDomainLogsResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeDomainLogsResponse struct {
	responses.BaseResponse
	RequestId  common.String
	TotalCount common.Long
	PageNumber common.Long
	PageSize   common.Long
	DomainLogs DescribeDomainLogsDomainLogList
}

type DescribeDomainLogsDomainLog struct {
	ActionTime common.String
	DomainName common.String
	Action     common.String
	Message    common.String
	ClientIp   common.String
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
