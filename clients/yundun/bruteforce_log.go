package yundun

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type BruteforceLogRequest struct {
	requests.RpcRequest
	JstOwnerId int64  `position:"Query" name:"JstOwnerId"`
	PageNumber int    `position:"Query" name:"PageNumber"`
	PageSize   int    `position:"Query" name:"PageSize"`
	InstanceId string `position:"Query" name:"InstanceId"`
	RecordType int    `position:"Query" name:"RecordType"`
}

func (req *BruteforceLogRequest) Invoke(client *sdk.Client) (resp *BruteforceLogResponse, err error) {
	req.InitWithApiInfo("Yundun", "2015-04-16", "BruteforceLog", "", "")
	resp = &BruteforceLogResponse{}
	err = client.DoAction(req, resp)
	return
}

type BruteforceLogResponse struct {
	responses.BaseResponse
	RequestId  string
	StartTime  string
	EndTime    string
	PageNumber int
	PageSize   int
	TotalCount int
	LogList    BruteforceLogBruteforceLogList
}

type BruteforceLogBruteforceLog struct {
	BlockTimes int
	SourceIp   string
	Status     int
	Time       string
	Location   string
}

type BruteforceLogBruteforceLogList []BruteforceLogBruteforceLog

func (list *BruteforceLogBruteforceLogList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]BruteforceLogBruteforceLog)
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
