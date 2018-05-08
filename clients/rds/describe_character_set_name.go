package rds

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type DescribeCharacterSetNameRequest struct {
	requests.RpcRequest
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	Engine               string `position:"Query" name:"Engine"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (req *DescribeCharacterSetNameRequest) Invoke(client *sdk.Client) (resp *DescribeCharacterSetNameResponse, err error) {
	req.InitWithApiInfo("Rds", "2014-08-15", "DescribeCharacterSetName", "rds", "")
	resp = &DescribeCharacterSetNameResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeCharacterSetNameResponse struct {
	responses.BaseResponse
	RequestId             common.String
	Engine                common.String
	CharacterSetNameItems DescribeCharacterSetNameCharacterSetNameItemList
}

type DescribeCharacterSetNameCharacterSetNameItemList []common.String

func (list *DescribeCharacterSetNameCharacterSetNameItemList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]common.String)
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
