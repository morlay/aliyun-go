package rds

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeOssDownloadsRequest struct {
	requests.RpcRequest
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	MigrateTaskId        string `position:"Query" name:"MigrateTaskId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	DBInstanceId         string `position:"Query" name:"DBInstanceId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (req *DescribeOssDownloadsRequest) Invoke(client *sdk.Client) (resp *DescribeOssDownloadsResponse, err error) {
	req.InitWithApiInfo("Rds", "2014-08-15", "DescribeOssDownloads", "rds", "")
	resp = &DescribeOssDownloadsResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeOssDownloadsResponse struct {
	responses.BaseResponse
	RequestId     string
	DBInstanceId  string
	MigrateTaskId string
	Items         DescribeOssDownloadsOssDownloadList
}

type DescribeOssDownloadsOssDownload struct {
	FileName    string
	CreateTime  string
	BackupMode  string
	FileSize    string
	Status      string
	IsAvailable string
	Description string
}

type DescribeOssDownloadsOssDownloadList []DescribeOssDownloadsOssDownload

func (list *DescribeOssDownloadsOssDownloadList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeOssDownloadsOssDownload)
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
