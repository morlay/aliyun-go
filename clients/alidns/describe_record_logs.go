package alidns

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type DescribeRecordLogsRequest struct {
	requests.RpcRequest
	Lang         string `position:"Query" name:"Lang"`
	UserClientIp string `position:"Query" name:"UserClientIp"`
	DomainName   string `position:"Query" name:"DomainName"`
	PageNumber   int64  `position:"Query" name:"PageNumber"`
	PageSize     int64  `position:"Query" name:"PageSize"`
	KeyWord      string `position:"Query" name:"KeyWord"`
}

func (req *DescribeRecordLogsRequest) Invoke(client *sdk.Client) (resp *DescribeRecordLogsResponse, err error) {
	req.InitWithApiInfo("Alidns", "2015-01-09", "DescribeRecordLogs", "", "")
	resp = &DescribeRecordLogsResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeRecordLogsResponse struct {
	responses.BaseResponse
	RequestId  common.String
	TotalCount common.Long
	PageNumber common.Long
	PageSize   common.Long
	RecordLogs DescribeRecordLogsRecordLogList
}

type DescribeRecordLogsRecordLog struct {
	ActionTime common.String
	Action     common.String
	Message    common.String
	ClientIp   common.String
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
