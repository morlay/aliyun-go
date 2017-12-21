package rds

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeSQLDiagnosisRequest struct {
	requests.RpcRequest
	SQLDiagId    string `position:"Query" name:"SQLDiagId"`
	DBInstanceId string `position:"Query" name:"DBInstanceId"`
}

func (req *DescribeSQLDiagnosisRequest) Invoke(client *sdk.Client) (resp *DescribeSQLDiagnosisResponse, err error) {
	req.InitWithApiInfo("Rds", "2014-08-15", "DescribeSQLDiagnosis", "rds", "")
	resp = &DescribeSQLDiagnosisResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeSQLDiagnosisResponse struct {
	responses.BaseResponse
	RequestId string
	SQLList   DescribeSQLDiagnosisSQLListList
}

type DescribeSQLDiagnosisSQLListList []string

func (list *DescribeSQLDiagnosisSQLListList) UnmarshalJSON(data []byte) error {
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
