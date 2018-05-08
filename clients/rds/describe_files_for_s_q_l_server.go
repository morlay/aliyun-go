package rds

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type DescribeFilesForSQLServerRequest struct {
	requests.RpcRequest
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

func (req *DescribeFilesForSQLServerRequest) Invoke(client *sdk.Client) (resp *DescribeFilesForSQLServerResponse, err error) {
	req.InitWithApiInfo("Rds", "2014-08-15", "DescribeFilesForSQLServer", "rds", "")
	resp = &DescribeFilesForSQLServerResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeFilesForSQLServerResponse struct {
	responses.BaseResponse
	RequestId        common.String
	DBInstanceId     common.String
	TotalRecordCount common.Integer
	PageNumber       common.Integer
	PageRecordCount  common.Integer
	Items            DescribeFilesForSQLServerSQLServerUploadFileList
}

type DescribeFilesForSQLServerSQLServerUploadFile struct {
	DBName            common.String
	FileName          common.String
	FileSize          common.Long
	InternetFtpServer common.String
	InternetPort      common.Integer
	IntranetFtpserver common.String
	Intranetport      common.Integer
	UserName          common.String
	Password          common.String
	FileStatus        common.String
	Description       common.String
	CreationTime      common.String
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
