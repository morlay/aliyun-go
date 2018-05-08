package cloudphoto

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type GetSimilarPhotosRequest struct {
	requests.RpcRequest
	LibraryId string `position:"Query" name:"LibraryId"`
	PhotoId   int64  `position:"Query" name:"PhotoId"`
	StoreName string `position:"Query" name:"StoreName"`
}

func (req *GetSimilarPhotosRequest) Invoke(client *sdk.Client) (resp *GetSimilarPhotosResponse, err error) {
	req.InitWithApiInfo("CloudPhoto", "2017-07-11", "GetSimilarPhotos", "cloudphoto", "")
	resp = &GetSimilarPhotosResponse{}
	err = client.DoAction(req, resp)
	return
}

type GetSimilarPhotosResponse struct {
	responses.BaseResponse
	Code      common.String
	Message   common.String
	RequestId common.String
	Action    common.String
	Photos    GetSimilarPhotosPhotoList
}

type GetSimilarPhotosPhoto struct {
	Id              common.Long
	IdStr           common.String
	Title           common.String
	FileId          common.String
	Location        common.String
	State           common.String
	Md5             common.String
	IsVideo         bool
	Remark          common.String
	Size            common.Long
	Width           common.Long
	Height          common.Long
	Ctime           common.Long
	Mtime           common.Long
	TakenAt         common.Long
	InactiveTime    common.Long
	ShareExpireTime common.Long
	Like            common.Long
}

type GetSimilarPhotosPhotoList []GetSimilarPhotosPhoto

func (list *GetSimilarPhotosPhotoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]GetSimilarPhotosPhoto)
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
