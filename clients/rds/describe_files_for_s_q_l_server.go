package rds

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeFilesForSQLServerRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	PageSize             int    `position:"Query" name:"PageSize"`
	EndTime              string `position:"Query" name:"EndTime"`
	DBInstanceId         string `position:"Query" name:"DBInstanceId"`
	StartTime            string `position:"Query" name:"StartTime"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	PageNumber           int    `position:"Query" name:"PageNumber"`
}

func (r DescribeFilesForSQLServerRequest) Invoke(client *sdk.Client) (response *DescribeFilesForSQLServerResponse, err error) {
	req := struct {
		*requests.RpcRequest
		DescribeFilesForSQLServerRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Rds", "2014-08-15", "DescribeFilesForSQLServer", "rds", "")

	resp := struct {
		*responses.BaseResponse
		DescribeFilesForSQLServerResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.DescribeFilesForSQLServerResponse

	err = client.DoAction(&req, &resp)
	return
}

type DescribeFilesForSQLServerResponse struct {
	RequestId        string
	DBInstanceId     string
	TotalRecordCount int
	PageNumber       int
	PageRecordCount  int
	Items            DescribeFilesForSQLServerSQLServerUploadFileList
}

type DescribeFilesForSQLServerSQLServerUploadFile struct {
	DBName            string
	FileName          string
	FileSize          int64
	InternetFtpServer string
	InternetPort      int
	IntranetFtpserver string
	Intranetport      int
	UserName          string
	Password          string
	FileStatus        string
	Description       string
	CreationTime      string
}

type DescribeFilesForSQLServerSQLServerUploadFileList []DescribeFilesForSQLServerSQLServerUploadFile

func (list *DescribeFilesForSQLServerSQLServerUploadFileList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeFilesForSQLServerSQLServerUploadFile)
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
