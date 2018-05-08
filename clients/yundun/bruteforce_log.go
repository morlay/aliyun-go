package yundun

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
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
	RequestId  common.String
	StartTime  common.String
	EndTime    common.String
	PageNumber common.Integer
	PageSize   common.Integer
	TotalCount common.Integer
	LogList    BruteforceLogBruteforceLogList
}

type BruteforceLogBruteforceLog struct {
	BlockTimes common.Integer
	SourceIp   common.String
	Status     common.Integer
	Time       common.String
	Location   common.String
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
