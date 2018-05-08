package polardb

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type DescribeBackupsRequest struct {
	requests.RpcRequest
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	DBClusterId          string `position:"Query" name:"DBClusterId"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	BackupId             string `position:"Query" name:"BackupId"`
	EndTime              string `position:"Query" name:"EndTime"`
	StartTime            string `position:"Query" name:"StartTime"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	PageNumber           int    `position:"Query" name:"PageNumber"`
	BackupStatus         string `position:"Query" name:"BackupStatus"`
	PageSize             int    `position:"Query" name:"PageSize"`
	BackupMode           string `position:"Query" name:"BackupMode"`
}

func (req *DescribeBackupsRequest) Invoke(client *sdk.Client) (resp *DescribeBackupsResponse, err error) {
	req.InitWithApiInfo("polardb", "2017-08-01", "DescribeBackups", "polardb", "")
	resp = &DescribeBackupsResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeBackupsResponse struct {
	responses.BaseResponse
	RequestId        common.String
	TotalRecordCount common.String
	PageNumber       common.String
	PageRecordCount  common.String
	TotalBackupSize  common.Long
	Items            DescribeBackupsBackupList
}

type DescribeBackupsBackup struct {
	BackupId                  common.String
	DBInstanceId              common.String
	BackupStatus              common.String
	BackupStartTime           common.String
	BackupEndTime             common.String
	BackupType                common.String
	BackupMode                common.String
	BackupMethod              common.String
	BackupDownloadURL         common.String
	BackupIntranetDownloadURL common.String
	BackupLocation            common.String
	BackupExtractionStatus    common.String
	BackupScale               common.String
	BackupDBNames             common.String
	TotalBackupSize           common.Long
	BackupSize                common.Long
	HostInstanceID            common.String
	StoreStatus               common.String
}

type DescribeBackupsBackupList []DescribeBackupsBackup

func (list *DescribeBackupsBackupList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeBackupsBackup)
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
