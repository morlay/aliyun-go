package rds

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeCharacterSetNameRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	Engine               string `position:"Query" name:"Engine"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (r DescribeCharacterSetNameRequest) Invoke(client *sdk.Client) (response *DescribeCharacterSetNameResponse, err error) {
	req := struct {
		*requests.RpcRequest
		DescribeCharacterSetNameRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Rds", "2014-08-15", "DescribeCharacterSetName", "rds", "")

	resp := struct {
		*responses.BaseResponse
		DescribeCharacterSetNameResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.DescribeCharacterSetNameResponse

	err = client.DoAction(&req, &resp)
	return
}

type DescribeCharacterSetNameResponse struct {
	RequestId             string
	Engine                string
	CharacterSetNameItems DescribeCharacterSetNameCharacterSetNameItemList
}

type DescribeCharacterSetNameCharacterSetNameItemList []string

func (list *DescribeCharacterSetNameCharacterSetNameItemList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]string)
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
