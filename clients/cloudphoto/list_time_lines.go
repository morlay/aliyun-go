package cloudphoto

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
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
	Code       string
	Message    string
	NextCursor int
	RequestId  string
	Action     string
	TimeLines  ListTimeLinesTimeLineList
}

type ListTimeLinesTimeLine struct {
	StartTime   int64
	EndTime     int64
	TotalCount  int
	PhotosCount int
	Photos      ListTimeLinesPhotoList
}

type ListTimeLinesPhoto struct {
	Id              int64
	Title           string
	FileId          string
	State           string
	Md5             string
	IsVideo         bool
	Remark          string
	Width           int64
	Height          int64
	Ctime           int64
	Mtime           int64
	TakenAt         int64
	ShareExpireTime int64
	Like            int64
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
