package emr

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type CreateResourcePoolRequest struct {
	requests.RpcRequest
	Note            string                        `position:"Query" name:"Note"`
	ResourceOwnerId int64                         `position:"Query" name:"ResourceOwnerId"`
	Name            string                        `position:"Query" name:"Name"`
	Active          string                        `position:"Query" name:"Active"`
	ClusterId       string                        `position:"Query" name:"ClusterId"`
	YarnSiteConfig  string                        `position:"Query" name:"YarnSiteConfig"`
	Configs         *CreateResourcePoolConfigList `position:"Query" type:"Repeated" name:"Config"`
	PoolType        string                        `position:"Query" name:"PoolType"`
}

func (req *CreateResourcePoolRequest) Invoke(client *sdk.Client) (resp *CreateResourcePoolResponse, err error) {
	req.InitWithApiInfo("Emr", "2016-04-08", "CreateResourcePool", "", "")
	resp = &CreateResourcePoolResponse{}
	err = client.DoAction(req, resp)
	return
}

type CreateResourcePoolConfig struct {
	ConfigKey   string `name:"ConfigKey"`
	ConfigValue string `name:"ConfigValue"`
	ConfigType  string `name:"ConfigType"`
	TargetId    string `name:"TargetId"`
	Category    string `name:"Category"`
	Note        string `name:"Note"`
}

type CreateResourcePoolResponse struct {
	responses.BaseResponse
	RequestId string
}

type CreateResourcePoolConfigList []CreateResourcePoolConfig

func (list *CreateResourcePoolConfigList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]CreateResourcePoolConfig)
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
