package rds

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeSQLDiagnosisRequest struct {
	SQLDiagId    string `position:"Query" name:"SQLDiagId"`
	DBInstanceId string `position:"Query" name:"DBInstanceId"`
}

func (r DescribeSQLDiagnosisRequest) Invoke(client *sdk.Client) (response *DescribeSQLDiagnosisResponse, err error) {
	req := struct {
		*requests.RpcRequest
		DescribeSQLDiagnosisRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Rds", "2014-08-15", "DescribeSQLDiagnosis", "rds", "")

	resp := struct {
		*responses.BaseResponse
		DescribeSQLDiagnosisResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.DescribeSQLDiagnosisResponse

	err = client.DoAction(&req, &resp)
	return
}

type DescribeSQLDiagnosisResponse struct {
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
