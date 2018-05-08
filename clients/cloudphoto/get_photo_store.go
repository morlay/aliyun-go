package cloudphoto

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type GetPhotoStoreRequest struct {
	requests.RpcRequest
	StoreName string `position:"Query" name:"StoreName"`
}

func (req *GetPhotoStoreRequest) Invoke(client *sdk.Client) (resp *GetPhotoStoreResponse, err error) {
	req.InitWithApiInfo("CloudPhoto", "2017-07-11", "GetPhotoStore", "cloudphoto", "")
	resp = &GetPhotoStoreResponse{}
	err = client.DoAction(req, resp)
	return
}

type GetPhotoStoreResponse struct {
	responses.BaseResponse
	Code       common.String
	Message    common.String
	RequestId  common.String
	Action     common.String
	PhotoStore GetPhotoStorePhotoStore
}

type GetPhotoStorePhotoStore struct {
	Id                common.Long
	IdStr             common.String
	Name              common.String
	Remark            common.String
	AutoCleanEnabled  bool
	AutoCleanDays     common.Integer
	DefaultQuota      common.Long
	DefaultTrashQuota common.Long
	Ctime             common.Long
	Mtime             common.Long
	Buckets           GetPhotoStoreBucketList
}

type GetPhotoStoreBucket struct {
	Name   common.String
	Region common.String
	State  common.String
	Acl    common.String
}

type GetPhotoStoreBucketList []GetPhotoStoreBucket

func (list *GetPhotoStoreBucketList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]GetPhotoStoreBucket)
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
