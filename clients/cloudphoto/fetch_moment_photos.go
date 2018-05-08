package cloudphoto

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type FetchMomentPhotosRequest struct {
	requests.RpcRequest
	Size      int    `position:"Query" name:"Size"`
	LibraryId string `position:"Query" name:"LibraryId"`
	OrderBy   string `position:"Query" name:"OrderBy"`
	StoreName string `position:"Query" name:"StoreName"`
	Page      int    `position:"Query" name:"Page"`
	MomentId  int64  `position:"Query" name:"MomentId"`
	Order     string `position:"Query" name:"Order"`
}

func (req *FetchMomentPhotosRequest) Invoke(client *sdk.Client) (resp *FetchMomentPhotosResponse, err error) {
	req.InitWithApiInfo("CloudPhoto", "2017-07-11", "FetchMomentPhotos", "cloudphoto", "")
	resp = &FetchMomentPhotosResponse{}
	err = client.DoAction(req, resp)
	return
}

type FetchMomentPhotosResponse struct {
	responses.BaseResponse
	Code       common.String
	Message    common.String
	TotalCount common.Integer
	RequestId  common.String
	Action     common.String
	Photos     FetchMomentPhotosPhotoList
}

type FetchMomentPhotosPhoto struct {
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
}

type FetchMomentPhotosPhotoList []FetchMomentPhotosPhoto

func (list *FetchMomentPhotosPhotoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]FetchMomentPhotosPhoto)
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
