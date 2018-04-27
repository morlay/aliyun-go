package emr

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ModifyResourcePoolRequest struct {
	requests.RpcRequest
	ResourceOwnerId int64                         `position:"Query" name:"ResourceOwnerId"`
	Name            string                        `position:"Query" name:"Name"`
	Active          string                        `position:"Query" name:"Active"`
	Id              string                        `position:"Query" name:"Id"`
	ClusterId       string                        `position:"Query" name:"ClusterId"`
	Yarnsiteconfig  string                        `position:"Query" name:"Yarnsiteconfig"`
	Configs         *ModifyResourcePoolConfigList `position:"Query" type:"Repeated" name:"Config"`
}

func (req *ModifyResourcePoolRequest) Invoke(client *sdk.Client) (resp *ModifyResourcePoolResponse, err error) {
	req.InitWithApiInfo("Emr", "2016-04-08", "ModifyResourcePool", "", "")
	resp = &ModifyResourcePoolResponse{}
	err = client.DoAction(req, resp)
	return
}

type ModifyResourcePoolConfig struct {
	Id          string `name:"Id"`
	ConfigKey   string `name:"ConfigKey"`
	ConfigValue string `name:"ConfigValue"`
	Category    string `name:"Category"`
	Note        string `name:"Note"`
}

type ModifyResourcePoolResponse struct {
	responses.BaseResponse
	RequestId string
}

type ModifyResourcePoolConfigList []ModifyResourcePoolConfig

func (list *ModifyResourcePoolConfigList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ModifyResourcePoolConfig)
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
