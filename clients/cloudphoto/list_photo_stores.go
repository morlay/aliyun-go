package cloudphoto

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type ListPhotoStoresRequest struct {
	requests.RpcRequest
}

func (req *ListPhotoStoresRequest) Invoke(client *sdk.Client) (resp *ListPhotoStoresResponse, err error) {
	req.InitWithApiInfo("CloudPhoto", "2017-07-11", "ListPhotoStores", "cloudphoto", "")
	resp = &ListPhotoStoresResponse{}
	err = client.DoAction(req, resp)
	return
}

type ListPhotoStoresResponse struct {
	responses.BaseResponse
	Code        common.String
	Message     common.String
	RequestId   common.String
	Action      common.String
	PhotoStores ListPhotoStoresPhotoStoreList
}

type ListPhotoStoresPhotoStore struct {
	Id               common.Long
	IdStr            common.String
	Name             common.String
	Remark           common.String
	AutoCleanEnabled bool
	AutoCleanDays    common.Integer
	DefaultQuota     common.Long
	Ctime            common.Long
	Mtime            common.Long
	Buckets          ListPhotoStoresBucketList
}

type ListPhotoStoresBucket struct {
	Name   common.String
	Region common.String
	State  common.String
}

type ListPhotoStoresPhotoStoreList []ListPhotoStoresPhotoStore

func (list *ListPhotoStoresPhotoStoreList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListPhotoStoresPhotoStore)
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

type ListPhotoStoresBucketList []ListPhotoStoresBucket

func (list *ListPhotoStoresBucketList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListPhotoStoresBucket)
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
