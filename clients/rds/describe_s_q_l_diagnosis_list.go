package rds

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeSQLDiagnosisListRequest struct {
	DBInstanceId string `position:"Query" name:"DBInstanceId"`
}

func (r DescribeSQLDiagnosisListRequest) Invoke(client *sdk.Client) (response *DescribeSQLDiagnosisListResponse, err error) {
	req := struct {
		*requests.RpcRequest
		DescribeSQLDiagnosisListRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Rds", "2014-08-15", "DescribeSQLDiagnosisList", "rds", "")

	resp := struct {
		*responses.BaseResponse
		DescribeSQLDiagnosisListResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.DescribeSQLDiagnosisListResponse

	err = client.DoAction(&req, &resp)
	return
}

type DescribeSQLDiagnosisListResponse struct {
	RequestId   string
	SQLDiagList DescribeSQLDiagnosisListSQLDiagList
}

type DescribeSQLDiagnosisListSQLDiag struct {
	SQLDiagId string
	StartTime string
	EndTime   string
	Status    int
	Progress  int
}

type DescribeSQLDiagnosisListSQLDiagList []DescribeSQLDiagnosisListSQLDiag

func (list *DescribeSQLDiagnosisListSQLDiagList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeSQLDiagnosisListSQLDiag)
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
