package rds

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeOssDownloadsForSQLServerRequest struct {
	requests.RpcRequest
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	MigrateTaskId        string `position:"Query" name:"MigrateTaskId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	DBInstanceId         string `position:"Query" name:"DBInstanceId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (req *DescribeOssDownloadsForSQLServerRequest) Invoke(client *sdk.Client) (resp *DescribeOssDownloadsForSQLServerResponse, err error) {
	req.InitWithApiInfo("Rds", "2014-08-15", "DescribeOssDownloadsForSQLServer", "rds", "")
	resp = &DescribeOssDownloadsForSQLServerResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeOssDownloadsForSQLServerResponse struct {
	responses.BaseResponse
	RequestId      string
	DBInstanceName string
	MigrateIaskId  string
	Items          DescribeOssDownloadsForSQLServerOssDownloadList
}

type DescribeOssDownloadsForSQLServerOssDownload struct {
	FileName    string
	CreateTime  string
	CreateTime1 string
	BakType     string
	FileSize    string
	Status      string
	IsAvail     string
	Desc        string
}

type DescribeOssDownloadsForSQLServerOssDownloadList []DescribeOssDownloadsForSQLServerOssDownload

func (list *DescribeOssDownloadsForSQLServerOssDownloadList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeOssDownloadsForSQLServerOssDownload)
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
