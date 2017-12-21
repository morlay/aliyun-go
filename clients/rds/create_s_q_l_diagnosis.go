package rds

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type CreateSQLDiagnosisRequest struct {
	requests.RpcRequest
	EndTime      string `position:"Query" name:"EndTime"`
	DBInstanceId string `position:"Query" name:"DBInstanceId"`
	StartTime    string `position:"Query" name:"StartTime"`
}

func (req *CreateSQLDiagnosisRequest) Invoke(client *sdk.Client) (resp *CreateSQLDiagnosisResponse, err error) {
	req.InitWithApiInfo("Rds", "2014-08-15", "CreateSQLDiagnosis", "rds", "")
	resp = &CreateSQLDiagnosisResponse{}
	err = client.DoAction(req, resp)
	return
}

type CreateSQLDiagnosisResponse struct {
	responses.BaseResponse
	RequestId string
	SQLDiagId string
}
