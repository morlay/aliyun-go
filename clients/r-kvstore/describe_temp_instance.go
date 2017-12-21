package r_kvstore

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeTempInstanceRequest struct {
	requests.RpcRequest
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	InstanceId           string `position:"Domain" name:"InstanceId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (req *DescribeTempInstanceRequest) Invoke(client *sdk.Client) (resp *DescribeTempInstanceResponse, err error) {
	req.InitWithApiInfo("R-kvstore", "2015-01-01", "DescribeTempInstance", "redisa", "")
	resp = &DescribeTempInstanceResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeTempInstanceResponse struct {
	responses.BaseResponse
	RequestId     string
	TempInstances DescribeTempInstanceTempInstanceList
}

type DescribeTempInstanceTempInstance struct {
	InstanceId     string
	TempInstanceId string
	SnapshotId     string
	CreateTime     string
	Domain         string
	Status         string
	Memory         int64
	ExpireTime     string
}

type DescribeTempInstanceTempInstanceList []DescribeTempInstanceTempInstance

func (list *DescribeTempInstanceTempInstanceList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeTempInstanceTempInstance)
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
