package emr

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
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
	RequestId    string
	PageNumber   int
	PageSize     int
	Total        int
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
	Id          int64
	ConfigKey   string
	ConfigValue string
	ConfigType  string
	Category    string
	Status      string
	Note        string
}

type ListResourcePoolEcmResourceQueue struct {
	Id             int64
	Name           string
	QualifiedName  string
	QueueType      string
	ParentQueueId  int64
	Leaf           bool
	Status         string
	UserId         string
	ResourcePoolId int64
}

type ListResourcePoolEcmResourcePoolConfig2 struct {
	Id          int64
	ConfigKey   string
	ConfigValue string
	ConfigType  string
	Category    string
	Status      string
	Note        string
}

type ListResourcePoolEcmResourcePool struct {
	Id             int64
	Name           string
	PoolType       string
	Active         bool
	Note           string
	UserId         string
	YarnSiteConfig string
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
