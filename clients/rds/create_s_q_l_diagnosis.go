package rds

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type CreateSQLDiagnosisRequest struct {
	EndTime      string `position:"Query" name:"EndTime"`
	DBInstanceId string `position:"Query" name:"DBInstanceId"`
	StartTime    string `position:"Query" name:"StartTime"`
}

func (r CreateSQLDiagnosisRequest) Invoke(client *sdk.Client) (response *CreateSQLDiagnosisResponse, err error) {
	req := struct {
		*requests.RpcRequest
		CreateSQLDiagnosisRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Rds", "2014-08-15", "CreateSQLDiagnosis", "rds", "")

	resp := struct {
		*responses.BaseResponse
		CreateSQLDiagnosisResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.CreateSQLDiagnosisResponse

	err = client.DoAction(&req, &resp)
	return
}

type CreateSQLDiagnosisResponse struct {
	RequestId string
	SQLDiagId string
}
