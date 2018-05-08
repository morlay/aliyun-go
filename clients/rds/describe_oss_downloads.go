package rds

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
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
	RequestId     common.String
	DBInstanceId  common.String
	MigrateTaskId common.String
	Items         DescribeOssDownloadsOssDownloadList
}

type DescribeOssDownloadsOssDownload struct {
	FileName    common.String
	CreateTime  common.String
	BackupMode  common.String
	FileSize    common.String
	Status      common.String
	IsAvailable common.String
	Description common.String
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
