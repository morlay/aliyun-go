package mts

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type ListAllMediaBucketRequest struct {
	requests.RpcRequest
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (req *ListAllMediaBucketRequest) Invoke(client *sdk.Client) (resp *ListAllMediaBucketResponse, err error) {
	req.InitWithApiInfo("Mts", "2014-06-18", "ListAllMediaBucket", "mts", "")
	resp = &ListAllMediaBucketResponse{}
	err = client.DoAction(req, resp)
	return
}

type ListAllMediaBucketResponse struct {
	responses.BaseResponse
	RequestId       common.String
	MediaBucketList ListAllMediaBucketMediaBucketList
}

type ListAllMediaBucketMediaBucket struct {
	Bucket common.String
	Type   common.String
}

type ListAllMediaBucketMediaBucketList []ListAllMediaBucketMediaBucket

func (list *ListAllMediaBucketMediaBucketList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListAllMediaBucketMediaBucket)
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
