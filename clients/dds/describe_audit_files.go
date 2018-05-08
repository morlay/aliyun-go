package dds

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type DescribeAuditFilesRequest struct {
	requests.RpcRequest
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	SecurityToken        string `position:"Query" name:"SecurityToken"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	PageSize             int    `position:"Query" name:"PageSize"`
	DBInstanceId         string `position:"Query" name:"DBInstanceId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	NodeId               string `position:"Query" name:"NodeId"`
	PageNumber           int    `position:"Query" name:"PageNumber"`
}

func (req *DescribeAuditFilesRequest) Invoke(client *sdk.Client) (resp *DescribeAuditFilesResponse, err error) {
	req.InitWithApiInfo("Dds", "2015-12-01", "DescribeAuditFiles", "dds", "")
	resp = &DescribeAuditFilesResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeAuditFilesResponse struct {
	responses.BaseResponse
	RequestId        common.String
	TotalRecordCount common.Integer
	PageNumber       common.Integer
	PageRecordCount  common.Integer
	DBInstanceId     common.String
	Items            DescribeAuditFilesLogFileList
}

type DescribeAuditFilesLogFile struct {
	FileID         common.Integer
	LogStatus      common.String
	LogStartTime   common.String
	LogEndTime     common.String
	LogDownloadURL common.String
	LogSize        common.Long
}

type DescribeAuditFilesLogFileList []DescribeAuditFilesLogFile

func (list *DescribeAuditFilesLogFileList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeAuditFilesLogFile)
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
