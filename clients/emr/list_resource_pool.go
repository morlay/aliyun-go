package emr

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type ListResourcePoolRequest struct {
	requests.RpcRequest
	ResourceOwnerId int64  `position:"Query" name:"ResourceOwnerId"`
	ClusterId       string `position:"Query" name:"ClusterId"`
	PoolType        string `position:"Query" name:"PoolType"`
}

func (req *ListResourcePoolRequest) Invoke(client *sdk.Client) (resp *ListResourcePoolResponse, err error) {
	req.InitWithApiInfo("Emr", "2016-04-08", "ListResourcePool", "", "")
	resp = &ListResourcePoolResponse{}
	err = client.DoAction(req, resp)
	return
}

type ListResourcePoolResponse struct {
	responses.BaseResponse
	RequestId    common.String
	PageNumber   common.Integer
	PageSize     common.Integer
	Total        common.Integer
	PoolInfoList ListResourcePoolPoolInfoList
}

type ListResourcePoolPoolInfo struct {
	QueueList                 ListResourcePoolQueueList
	EcmResourcePoolConfigList ListResourcePoolEcmResourcePoolConfig2List
	EcmResourcePool           ListResourcePoolEcmResourcePool
}

type ListResourcePoolQueue struct {
	EcmResourcePoolConfigList1 ListResourcePoolEcmResourcePoolConfigList
	EcmResourceQueue           ListResourcePoolEcmResourceQueue
}

type ListResourcePoolEcmResourcePoolConfig struct {
	Id          common.Long
	ConfigKey   common.String
	ConfigValue common.String
	ConfigType  common.String
	Category    common.String
	Status      common.String
	Note        common.String
}

type ListResourcePoolEcmResourceQueue struct {
	Id             common.Long
	Name           common.String
	QualifiedName  common.String
	QueueType      common.String
	ParentQueueId  common.Long
	Leaf           bool
	Status         common.String
	UserId         common.String
	ResourcePoolId common.Long
}

type ListResourcePoolEcmResourcePoolConfig2 struct {
	Id          common.Long
	ConfigKey   common.String
	ConfigValue common.String
	ConfigType  common.String
	Category    common.String
	Status      common.String
	Note        common.String
}

type ListResourcePoolEcmResourcePool struct {
	Id             common.Long
	Name           common.String
	PoolType       common.String
	Active         bool
	Note           common.String
	UserId         common.String
	YarnSiteConfig common.String
}

type ListResourcePoolPoolInfoList []ListResourcePoolPoolInfo

func (list *ListResourcePoolPoolInfoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListResourcePoolPoolInfo)
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

type ListResourcePoolQueueList []ListResourcePoolQueue

func (list *ListResourcePoolQueueList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListResourcePoolQueue)
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

type ListResourcePoolEcmResourcePoolConfig2List []ListResourcePoolEcmResourcePoolConfig2

func (list *ListResourcePoolEcmResourcePoolConfig2List) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListResourcePoolEcmResourcePoolConfig2)
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

type ListResourcePoolEcmResourcePoolConfigList []ListResourcePoolEcmResourcePoolConfig

func (list *ListResourcePoolEcmResourcePoolConfigList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListResourcePoolEcmResourcePoolConfig)
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
