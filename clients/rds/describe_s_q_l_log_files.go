package rds

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type DescribeSQLLogFilesRequest struct {
	requests.RpcRequest
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	FileName             string `position:"Query" name:"FileName"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	PageSize             int    `position:"Query" name:"PageSize"`
	DBInstanceId         string `position:"Query" name:"DBInstanceId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	PageNumber           int    `position:"Query" name:"PageNumber"`
}

func (req *DescribeSQLLogFilesRequest) Invoke(client *sdk.Client) (resp *DescribeSQLLogFilesResponse, err error) {
	req.InitWithApiInfo("Rds", "2014-08-15", "DescribeSQLLogFiles", "rds", "")
	resp = &DescribeSQLLogFilesResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeSQLLogFilesResponse struct {
	responses.BaseResponse
	RequestId        common.String
	TotalRecordCount common.Integer
	PageNumber       common.Integer
	PageRecordCount  common.Integer
	Items            DescribeSQLLogFilesLogFileList
}

type DescribeSQLLogFilesLogFile struct {
	FileID         common.String
	LogStatus      common.String
	LogDownloadURL common.String
	LogSize        common.String
	LogStartTime   common.String
	LogEndTime     common.String
}

type DescribeSQLLogFilesLogFileList []DescribeSQLLogFilesLogFile

func (list *DescribeSQLLogFilesLogFileList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeSQLLogFilesLogFile)
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
