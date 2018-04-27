package emr

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type CreateResourceQueueRequest struct {
	requests.RpcRequest
	ResourceOwnerId int64                          `position:"Query" name:"ResourceOwnerId"`
	ParentQueueId   int64                          `position:"Query" name:"ParentQueueId"`
	Name            string                         `position:"Query" name:"Name"`
	QualifiedName   string                         `position:"Query" name:"QualifiedName"`
	ResourcePoolId  int64                          `position:"Query" name:"ResourcePoolId"`
	ClusterId       string                         `position:"Query" name:"ClusterId"`
	Leaf            string                         `position:"Query" name:"Leaf"`
	Configs         *CreateResourceQueueConfigList `position:"Query" type:"Repeated" name:"Config"`
}

func (req *CreateResourceQueueRequest) Invoke(client *sdk.Client) (resp *CreateResourceQueueResponse, err error) {
	req.InitWithApiInfo("Emr", "2016-04-08", "CreateResourceQueue", "", "")
	resp = &CreateResourceQueueResponse{}
	err = client.DoAction(req, resp)
	return
}

type CreateResourceQueueConfig struct {
	ConfigKey   string `name:"ConfigKey"`
	ConfigValue string `name:"ConfigValue"`
	Category    string `name:"Category"`
	Note        string `name:"Note"`
}

type CreateResourceQueueResponse struct {
	responses.BaseResponse
	RequestId string
}

type CreateResourceQueueConfigList []CreateResourceQueueConfig

func (list *CreateResourceQueueConfigList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]CreateResourceQueueConfig)
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
