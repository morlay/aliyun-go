package rds

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
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
	RequestId      common.String
	DBInstanceName common.String
	MigrateIaskId  common.String
	Items          DescribeOssDownloadsForSQLServerOssDownloadList
}

type DescribeOssDownloadsForSQLServerOssDownload struct {
	FileName    common.String
	CreateTime  common.String
	CreateTime1 common.String
	BakType     common.String
	FileSize    common.String
	Status      common.String
	IsAvail     common.String
	Desc        common.String
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
