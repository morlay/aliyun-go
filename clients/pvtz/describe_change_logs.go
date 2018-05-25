package pvtz

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type DescribeChangeLogsRequest struct {
	requests.RpcRequest
	EntityType     string `position:"Query" name:"EntityType"`
	PageSize       int    `position:"Query" name:"PageSize"`
	UserClientIp   string `position:"Query" name:"UserClientIp"`
	ZoneId         string `position:"Query" name:"ZoneId"`
	Keyword        string `position:"Query" name:"Keyword"`
	Lang           string `position:"Query" name:"Lang"`
	StartTimestamp int64  `position:"Query" name:"StartTimestamp"`
	PageNumber     int    `position:"Query" name:"PageNumber"`
	EndTimestamp   int64  `position:"Query" name:"EndTimestamp"`
}

func (req *DescribeChangeLogsRequest) Invoke(client *sdk.Client) (resp *DescribeChangeLogsResponse, err error) {
	req.InitWithApiInfo("pvtz", "2018-01-01", "DescribeChangeLogs", "pvtz", "")
	resp = &DescribeChangeLogsResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeChangeLogsResponse struct {
	responses.BaseResponse
	RequestId  common.String
	TotalItems common.Integer
	TotalPages common.Integer
	PageSize   common.Integer
	PageNumber common.Integer
	ChangeLogs DescribeChangeLogsChangeLogList
}

type DescribeChangeLogsChangeLog struct {
	OperTime      common.String
	OperAction    common.String
	OperObject    common.String
	EntityId      common.String
	EntityName    common.String
	OperIp        common.String
	OperTimestamp common.Long
	Id            common.Long
	Content       common.String
}

type DescribeChangeLogsChangeLogList []DescribeChangeLogsChangeLog

func (list *DescribeChangeLogsChangeLogList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeChangeLogsChangeLog)
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
