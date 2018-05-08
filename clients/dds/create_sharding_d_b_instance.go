package dds

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type CreateShardingDBInstanceRequest struct {
	requests.RpcRequest
	ResourceOwnerId       int64                                     `position:"Query" name:"ResourceOwnerId"`
	ClientToken           string                                    `position:"Query" name:"ClientToken"`
	EngineVersion         string                                    `position:"Query" name:"EngineVersion"`
	NetworkType           string                                    `position:"Query" name:"NetworkType"`
	ReplicaSets           *CreateShardingDBInstanceReplicaSetList   `position:"Query" type:"Repeated" name:"ReplicaSet"`
	StorageEngine         string                                    `position:"Query" name:"StorageEngine"`
	SecurityToken         string                                    `position:"Query" name:"SecurityToken"`
	Engine                string                                    `position:"Query" name:"Engine"`
	DBInstanceDescription string                                    `position:"Query" name:"DBInstanceDescription"`
	Period                int                                       `position:"Query" name:"Period"`
	RestoreTime           string                                    `position:"Query" name:"RestoreTime"`
	ResourceOwnerAccount  string                                    `position:"Query" name:"ResourceOwnerAccount"`
	SrcDBInstanceId       string                                    `position:"Query" name:"SrcDBInstanceId"`
	OwnerAccount          string                                    `position:"Query" name:"OwnerAccount"`
	ConfigServers         *CreateShardingDBInstanceConfigServerList `position:"Query" type:"Repeated" name:"ConfigServer"`
	OwnerId               int64                                     `position:"Query" name:"OwnerId"`
	Mongoss               *CreateShardingDBInstanceMongosList       `position:"Query" type:"Repeated" name:"Mongos"`
	SecurityIPList        string                                    `position:"Query" name:"SecurityIPList"`
	VSwitchId             string                                    `position:"Query" name:"VSwitchId"`
	AccountPassword       string                                    `position:"Query" name:"AccountPassword"`
	VpcId                 string                                    `position:"Query" name:"VpcId"`
	ZoneId                string                                    `position:"Query" name:"ZoneId"`
	ChargeType            string                                    `position:"Query" name:"ChargeType"`
}

func (req *CreateShardingDBInstanceRequest) Invoke(client *sdk.Client) (resp *CreateShardingDBInstanceResponse, err error) {
	req.InitWithApiInfo("Dds", "2015-12-01", "CreateShardingDBInstance", "dds", "")
	resp = &CreateShardingDBInstanceResponse{}
	err = client.DoAction(req, resp)
	return
}

type CreateShardingDBInstanceReplicaSet struct {
	_class  string `name:"_class"`
	Storage int    `name:"Storage"`
}

type CreateShardingDBInstanceConfigServer struct {
	_class  string `name:"_class"`
	Storage int    `name:"Storage"`
}

type CreateShardingDBInstanceMongos struct {
	_class string `name:"_class"`
}

type CreateShardingDBInstanceResponse struct {
	responses.BaseResponse
	RequestId    common.String
	OrderId      common.String
	DBInstanceId common.String
}

type CreateShardingDBInstanceReplicaSetList []CreateShardingDBInstanceReplicaSet

func (list *CreateShardingDBInstanceReplicaSetList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]CreateShardingDBInstanceReplicaSet)
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

type CreateShardingDBInstanceConfigServerList []CreateShardingDBInstanceConfigServer

func (list *CreateShardingDBInstanceConfigServerList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]CreateShardingDBInstanceConfigServer)
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

type CreateShardingDBInstanceMongosList []CreateShardingDBInstanceMongos

func (list *CreateShardingDBInstanceMongosList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]CreateShardingDBInstanceMongos)
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
