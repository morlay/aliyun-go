package mts

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ListAllMediaBucketRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (r ListAllMediaBucketRequest) Invoke(client *sdk.Client) (response *ListAllMediaBucketResponse, err error) {
	req := struct {
		*requests.RpcRequest
		ListAllMediaBucketRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Mts", "2014-06-18", "ListAllMediaBucket", "", "")

	resp := struct {
		*responses.BaseResponse
		ListAllMediaBucketResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.ListAllMediaBucketResponse

	err = client.DoAction(&req, &resp)
	return
}

type ListAllMediaBucketResponse struct {
	RequestId       string
	MediaBucketList ListAllMediaBucketMediaBucketList
}

type ListAllMediaBucketMediaBucket struct {
	Bucket string
	Type   string
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
