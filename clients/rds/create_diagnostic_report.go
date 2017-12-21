package rds

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type CreateDiagnosticReportRequest struct {
	EndTime      string `position:"Query" name:"EndTime"`
	DBInstanceId string `position:"Query" name:"DBInstanceId"`
	StartTime    string `position:"Query" name:"StartTime"`
}

func (r CreateDiagnosticReportRequest) Invoke(client *sdk.Client) (response *CreateDiagnosticReportResponse, err error) {
	req := struct {
		*requests.RpcRequest
		CreateDiagnosticReportRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Rds", "2014-08-15", "CreateDiagnosticReport", "rds", "")

	resp := struct {
		*responses.BaseResponse
		CreateDiagnosticReportResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.CreateDiagnosticReportResponse

	err = client.DoAction(&req, &resp)
	return
}

type CreateDiagnosticReportResponse struct {
	RequestId string
	ReportId  string
}
