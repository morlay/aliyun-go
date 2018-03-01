package imm

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type GetFaceSetDetailRequest struct {
	requests.RpcRequest
	Marker          string `position:"Query" name:"Marker"`
	Project         string `position:"Query" name:"Project"`
	SetId           string `position:"Query" name:"SetId"`
	ReturnAttribute string `position:"Query" name:"ReturnAttribute"`
}

func (req *GetFaceSetDetailRequest) Invoke(client *sdk.Client) (resp *GetFaceSetDetailResponse, err error) {
	req.InitWithApiInfo("imm", "2017-09-06", "GetFaceSetDetail", "imm", "")
	resp = &GetFaceSetDetailResponse{}
	err = client.DoAction(req, resp)
	return
}

type GetFaceSetDetailResponse struct {
	responses.BaseResponse
	RequestId   string
	SetId       string
	NextMarker  string
	FaceDetails GetFaceSetDetailFaceDetailsItemList
}

type GetFaceSetDetailFaceDetailsItem struct {
	FaceId        string
	SrcUri        string
	PhotoId       string
	GroupId       string
	UnGroupReason string
	FaceRectangle GetFaceSetDetailFaceRectangle
	FaceAttribute GetFaceSetDetailFaceAttribute
}

type GetFaceSetDetailFaceRectangle struct {
	Top    int
	Left   int
	Width  int
	Height int
}

type GetFaceSetDetailFaceAttribute struct {
	Gender      GetFaceSetDetailGender
	Age         GetFaceSetDetailAge
	Headpose    GetFaceSetDetailHeadpose
	Eyestatus   GetFaceSetDetailEyestatus
	Blur        GetFaceSetDetailBlur
	Facequality GetFaceSetDetailFacequality
}

type GetFaceSetDetailGender struct {
	Value string
}

type GetFaceSetDetailAge struct {
	Value string
}

type GetFaceSetDetailHeadpose struct {
	Pitch_angle float32
	Roll_angle  float32
	Yaw_angle   float32
}

type GetFaceSetDetailEyestatus struct {
	Left_eye_status  GetFaceSetDetailLeft_eye_status
	Right_eye_status GetFaceSetDetailRight_eye_status
}

type GetFaceSetDetailLeft_eye_status struct {
	Normal_glass_eye_open  float32
	No_glass_eye_close     float32
	Occlusion              float32
	No_glass_eye_open      float32
	Normal_glass_eye_close float32
	Dark_glasses           float32
}

type GetFaceSetDetailRight_eye_status struct {
	Normal_glass_eye_open  float32
	No_glass_eye_close     float32
	Occlusion              float32
	No_glass_eye_open      float32
	Normal_glass_eye_close float32
	Dark_glasses           float32
}

type GetFaceSetDetailBlur struct {
	Blurness GetFaceSetDetailBlurness
}

type GetFaceSetDetailBlurness struct {
	Value     float32
	Threshold float32
}

type GetFaceSetDetailFacequality struct {
	Value     float32
	Threshold float32
}

type GetFaceSetDetailFaceDetailsItemList []GetFaceSetDetailFaceDetailsItem

func (list *GetFaceSetDetailFaceDetailsItemList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]GetFaceSetDetailFaceDetailsItem)
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
