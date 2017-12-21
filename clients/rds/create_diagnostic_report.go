package rds

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type CreateDiagnosticReportRequest struct {
	requests.RpcRequest
	EndTime      string `position:"Query" name:"EndTime"`
	DBInstanceId string `position:"Query" name:"DBInstanceId"`
	StartTime    string `position:"Query" name:"StartTime"`
}

func (req *CreateDiagnosticReportRequest) Invoke(client *sdk.Client) (resp *CreateDiagnosticReportResponse, err error) {
	req.InitWithApiInfo("Rds", "2014-08-15", "CreateDiagnosticReport", "rds", "")
	resp = &CreateDiagnosticReportResponse{}
	err = client.DoAction(req, resp)
	return
}

type CreateDiagnosticReportResponse struct {
	responses.BaseResponse
	RequestId string
	ReportId  string
}
