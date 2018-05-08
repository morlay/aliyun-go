package cloudphoto

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type ListTimeLinesRequest struct {
	requests.RpcRequest
	Cursor        int64  `position:"Query" name:"Cursor"`
	PhotoSize     int    `position:"Query" name:"PhotoSize"`
	TimeLineCount int    `position:"Query" name:"TimeLineCount"`
	LibraryId     string `position:"Query" name:"LibraryId"`
	StoreName     string `position:"Query" name:"StoreName"`
	TimeLineUnit  string `position:"Query" name:"TimeLineUnit"`
	FilterBy      string `position:"Query" name:"FilterBy"`
	Direction     string `position:"Query" name:"Direction"`
	Order         string `position:"Query" name:"Order"`
}

func (req *ListTimeLinesRequest) Invoke(client *sdk.Client) (resp *ListTimeLinesResponse, err error) {
	req.InitWithApiInfo("CloudPhoto", "2017-07-11", "ListTimeLines", "cloudphoto", "")
	resp = &ListTimeLinesResponse{}
	err = client.DoAction(req, resp)
	return
}

type ListTimeLinesResponse struct {
	responses.BaseResponse
	Code       common.String
	Message    common.String
	NextCursor common.Integer
	RequestId  common.String
	Action     common.String
	TimeLines  ListTimeLinesTimeLineList
}

type ListTimeLinesTimeLine struct {
	StartTime   common.Long
	EndTime     common.Long
	TotalCount  common.Integer
	PhotosCount common.Integer
	Photos      ListTimeLinesPhotoList
}

type ListTimeLinesPhoto struct {
	Id              common.Long
	IdStr           common.String
	Title           common.String
	Location        common.String
	FileId          common.String
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

type ListTimeLinesTimeLineList []ListTimeLinesTimeLine

func (list *ListTimeLinesTimeLineList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListTimeLinesTimeLine)
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

type ListTimeLinesPhotoList []ListTimeLinesPhoto

func (list *ListTimeLinesPhotoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListTimeLinesPhoto)
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
