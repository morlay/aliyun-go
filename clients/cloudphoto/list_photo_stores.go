package cloudphoto

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ListPhotoStoresRequest struct {
}

func (r ListPhotoStoresRequest) Invoke(client *sdk.Client) (response *ListPhotoStoresResponse, err error) {
	req := struct {
		*requests.RpcRequest
		ListPhotoStoresRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("CloudPhoto", "2017-07-11", "ListPhotoStores", "cloudphoto", "")

	resp := struct {
		*responses.BaseResponse
		ListPhotoStoresResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.ListPhotoStoresResponse

	err = client.DoAction(&req, &resp)
	return
}

type ListPhotoStoresResponse struct {
	Code        string
	Message     string
	RequestId   string
	Action      string
	PhotoStores ListPhotoStoresPhotoStoreList
}

type ListPhotoStoresPhotoStore struct {
	Id               int64
	Name             string
	Remark           string
	AutoCleanEnabled bool
	AutoCleanDays    int
	DefaultQuota     int64
	Ctime            int64
	Mtime            int64
	Buckets          ListPhotoStoresBucketList
}

type ListPhotoStoresBucket struct {
	Name   string
	Region string
	State  string
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
