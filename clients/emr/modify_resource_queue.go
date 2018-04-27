package emr

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ModifyResourceQueueRequest struct {
	requests.RpcRequest
	ResourceOwnerId int64                          `position:"Query" name:"ResourceOwnerId"`
	ParentQueueId   int64                          `position:"Query" name:"ParentQueueId"`
	Name            string                         `position:"Query" name:"Name"`
	QualifiedName   string                         `position:"Query" name:"QualifiedName"`
	ResourcePoolId  int64                          `position:"Query" name:"ResourcePoolId"`
	Id              string                         `position:"Query" name:"Id"`
	ClusterId       string                         `position:"Query" name:"ClusterId"`
	Leaf            string                         `position:"Query" name:"Leaf"`
	Configs         *ModifyResourceQueueConfigList `position:"Query" type:"Repeated" name:"Config"`
}

func (req *ModifyResourceQueueRequest) Invoke(client *sdk.Client) (resp *ModifyResourceQueueResponse, err error) {
	req.InitWithApiInfo("Emr", "2016-04-08", "ModifyResourceQueue", "", "")
	resp = &ModifyResourceQueueResponse{}
	err = client.DoAction(req, resp)
	return
}

type ModifyResourceQueueConfig struct {
	Id          int64  `name:"Id"`
	ConfigKey   string `name:"ConfigKey"`
	ConfigValue string `name:"ConfigValue"`
	Category    string `name:"Category"`
	Note        string `name:"Note"`
}

type ModifyResourceQueueResponse struct {
	responses.BaseResponse
	RequestId string
}

type ModifyResourceQueueConfigList []ModifyResourceQueueConfig

func (list *ModifyResourceQueueConfigList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ModifyResourceQueueConfig)
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
