package rds

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeDiagnosticReportListRequest struct {
	requests.RpcRequest
	DBInstanceId string `position:"Query" name:"DBInstanceId"`
}

func (req *DescribeDiagnosticReportListRequest) Invoke(client *sdk.Client) (resp *DescribeDiagnosticReportListResponse, err error) {
	req.InitWithApiInfo("Rds", "2014-08-15", "DescribeDiagnosticReportList", "rds", "")
	resp = &DescribeDiagnosticReportListResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeDiagnosticReportListResponse struct {
	responses.BaseResponse
	RequestId  string
	ReportList DescribeDiagnosticReportListReportList
}

type DescribeDiagnosticReportListReport struct {
	DiagnosticTime string
	Score          int
	StartTime      string
	EndTime        string
	DownloadURL    string
}

type DescribeDiagnosticReportListReportList []DescribeDiagnosticReportListReport

func (list *DescribeDiagnosticReportListReportList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeDiagnosticReportListReport)
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
