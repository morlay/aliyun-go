package rds

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeDatabaseLockDiagnosisRequest struct {
	requests.RpcRequest
	EndTime      string `position:"Query" name:"EndTime"`
	DBInstanceId string `position:"Query" name:"DBInstanceId"`
	StartTime    string `position:"Query" name:"StartTime"`
}

func (req *DescribeDatabaseLockDiagnosisRequest) Invoke(client *sdk.Client) (resp *DescribeDatabaseLockDiagnosisResponse, err error) {
	req.InitWithApiInfo("Rds", "2014-08-15", "DescribeDatabaseLockDiagnosis", "rds", "")
	resp = &DescribeDatabaseLockDiagnosisResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeDatabaseLockDiagnosisResponse struct {
	responses.BaseResponse
	RequestId    string
	DeadLockList DescribeDatabaseLockDiagnosisDeadLockListList
}

type DescribeDatabaseLockDiagnosisDeadLockListList []string

func (list *DescribeDatabaseLockDiagnosisDeadLockListList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]string)
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
