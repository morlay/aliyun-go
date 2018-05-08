package rds

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type DescribeDBInstanceHAConfigRequest struct {
	requests.RpcRequest
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	DBInstanceId         string `position:"Query" name:"DBInstanceId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (req *DescribeDBInstanceHAConfigRequest) Invoke(client *sdk.Client) (resp *DescribeDBInstanceHAConfigResponse, err error) {
	req.InitWithApiInfo("Rds", "2014-08-15", "DescribeDBInstanceHAConfig", "rds", "")
	resp = &DescribeDBInstanceHAConfigResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeDBInstanceHAConfigResponse struct {
	responses.BaseResponse
	RequestId         common.String
	DBInstanceId      common.String
	SyncMode          common.String
	HAMode            common.String
	HostInstanceInfos DescribeDBInstanceHAConfigNodeInfoList
}

type DescribeDBInstanceHAConfigNodeInfo struct {
	NodeId       common.String
	RegionId     common.String
	LogSyncTime  common.String
	DataSyncTime common.String
	NodeType     common.String
	ZoneId       common.String
	SyncStatus   common.String
}

type DescribeDBInstanceHAConfigNodeInfoList []DescribeDBInstanceHAConfigNodeInfo

func (list *DescribeDBInstanceHAConfigNodeInfoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeDBInstanceHAConfigNodeInfo)
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
