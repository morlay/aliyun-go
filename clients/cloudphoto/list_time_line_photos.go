package cloudphoto

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type ListTimeLinePhotosRequest struct {
	requests.RpcRequest
	Size      int    `position:"Query" name:"Size"`
	LibraryId string `position:"Query" name:"LibraryId"`
	EndTime   int64  `position:"Query" name:"EndTime"`
	StoreName string `position:"Query" name:"StoreName"`
	Page      int    `position:"Query" name:"Page"`
	StartTime int64  `position:"Query" name:"StartTime"`
	FilterBy  string `position:"Query" name:"FilterBy"`
	Direction string `position:"Query" name:"Direction"`
	Order     string `position:"Query" name:"Order"`
}

func (req *ListTimeLinePhotosRequest) Invoke(client *sdk.Client) (resp *ListTimeLinePhotosResponse, err error) {
	req.InitWithApiInfo("CloudPhoto", "2017-07-11", "ListTimeLinePhotos", "cloudphoto", "")
	resp = &ListTimeLinePhotosResponse{}
	err = client.DoAction(req, resp)
	return
}

type ListTimeLinePhotosResponse struct {
	responses.BaseResponse
	Code       common.String
	Message    common.String
	TotalCount common.Integer
	RequestId  common.String
	Action     common.String
	Photos     ListTimeLinePhotosPhotoList
}

type ListTimeLinePhotosPhoto struct {
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
	ShareExpireTime common.Long
	Like            common.Long
}

type ListTimeLinePhotosPhotoList []ListTimeLinePhotosPhoto

func (list *ListTimeLinePhotosPhotoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListTimeLinePhotosPhoto)
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
