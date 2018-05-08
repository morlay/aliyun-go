package rds

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type DescribeBackupSetsForSecurityRequest struct {
	requests.RpcRequest
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	TargetAliBid         string `position:"Query" name:"TargetAliBid"`
	BackupId             string `position:"Query" name:"BackupId"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	EndTime              string `position:"Query" name:"EndTime"`
	StartTime            string `position:"Query" name:"StartTime"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	PageNumber           int    `position:"Query" name:"PageNumber"`
	BackupStatus         string `position:"Query" name:"BackupStatus"`
	BackupLocation       string `position:"Query" name:"BackupLocation"`
	TargetAliUid         string `position:"Query" name:"TargetAliUid"`
	PageSize             int    `position:"Query" name:"PageSize"`
	DBInstanceId         string `position:"Query" name:"DBInstanceId"`
	BackupMode           string `position:"Query" name:"BackupMode"`
}

func (req *DescribeBackupSetsForSecurityRequest) Invoke(client *sdk.Client) (resp *DescribeBackupSetsForSecurityResponse, err error) {
	req.InitWithApiInfo("Rds", "2014-08-15", "DescribeBackupSetsForSecurity", "rds", "")
	resp = &DescribeBackupSetsForSecurityResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeBackupSetsForSecurityResponse struct {
	responses.BaseResponse
	RequestId        common.String
	TotalRecordCount common.String
	PageNumber       common.String
	PageRecordCount  common.String
	TotalBackupSize  common.Long
	Items            DescribeBackupSetsForSecurityBackupList
}

type DescribeBackupSetsForSecurityBackup struct {
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

type DescribeBackupSetsForSecurityBackupList []DescribeBackupSetsForSecurityBackup

func (list *DescribeBackupSetsForSecurityBackupList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeBackupSetsForSecurityBackup)
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
