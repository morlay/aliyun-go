package alidns

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeRecordLogsRequest struct {
	Lang         string `position:"Query" name:"Lang"`
	UserClientIp string `position:"Query" name:"UserClientIp"`
	DomainName   string `position:"Query" name:"DomainName"`
	PageNumber   int64  `position:"Query" name:"PageNumber"`
	PageSize     int64  `position:"Query" name:"PageSize"`
	KeyWord      string `position:"Query" name:"KeyWord"`
}

func (r DescribeRecordLogsRequest) Invoke(client *sdk.Client) (response *DescribeRecordLogsResponse, err error) {
	req := struct {
		*requests.RpcRequest
		DescribeRecordLogsRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Alidns", "2015-01-09", "DescribeRecordLogs", "", "")

	resp := struct {
		*responses.BaseResponse
		DescribeRecordLogsResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.DescribeRecordLogsResponse

	err = client.DoAction(&req, &resp)
	return
}

type DescribeRecordLogsResponse struct {
	RequestId  string
	TotalCount int64
	PageNumber int64
	PageSize   int64
	RecordLogs DescribeRecordLogsRecordLogList
}

type DescribeRecordLogsRecordLog struct {
	ActionTime string
	Action     string
	Message    string
	ClientIp   string
}

type DescribeRecordLogsRecordLogList []DescribeRecordLogsRecordLog

func (list *DescribeRecordLogsRecordLogList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeRecordLogsRecordLog)
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
