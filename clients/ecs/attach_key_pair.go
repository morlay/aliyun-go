package ecs

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type AttachKeyPairRequest struct {
	requests.RpcRequest
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	InstanceIds          string `position:"Query" name:"InstanceIds"`
	KeyPairName          string `position:"Query" name:"KeyPairName"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (req *AttachKeyPairRequest) Invoke(client *sdk.Client) (resp *AttachKeyPairResponse, err error) {
	req.InitWithApiInfo("Ecs", "2014-05-26", "AttachKeyPair", "ecs", "")
	resp = &AttachKeyPairResponse{}
	err = client.DoAction(req, resp)
	return
}

type AttachKeyPairResponse struct {
	responses.BaseResponse
	RequestId   string
	TotalCount  string
	FailCount   string
	KeyPairName string
	Results     AttachKeyPairResultList
}

type AttachKeyPairResult struct {
	InstanceId string
	Success    string
	Code       string
	Message    string
}

type AttachKeyPairResultList []AttachKeyPairResult

func (list *AttachKeyPairResultList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]AttachKeyPairResult)
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
