package cloudphoto

import "encoding/json"

func (c *CloudphotoClient) EditPhotos(req *EditPhotosArgs) (resp *EditPhotosResponse, err error) {
	resp = &EditPhotosResponse{}
	err = c.Request("EditPhotos", req, resp)
	return
}

type EditPhotosResult struct {
	Id      int64
	Code    string
	Message string
}
type EditPhotosArgs struct {
	LibraryId       string
	ShareExpireTime int64
	PhotoIds        EditPhotosPhotoIdList
	StoreName       string
	Remark          string
	Title           string
}
type EditPhotosResponse struct {
	Code      string
	Message   string
	RequestId string
	Action    string
	Results   EditPhotosResultList
}

type EditPhotosPhotoIdList []int64

func (list *EditPhotosPhotoIdList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]int64)
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

type EditPhotosResultList []EditPhotosResult

func (list *EditPhotosResultList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]EditPhotosResult)
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

func (c *CloudphotoClient) ListTimeLines(req *ListTimeLinesArgs) (resp *ListTimeLinesResponse, err error) {
	resp = &ListTimeLinesResponse{}
	err = c.Request("ListTimeLines", req, resp)
	return
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
type ListTimeLinesArgs struct {
	Cursor        int64
	PhotoSize     int
	TimeLineCount int
	LibraryId     string
	StoreName     string
	TimeLineUnit  string
	FilterBy      string
	Direction     string
	Order         string
}
type ListTimeLinesResponse struct {
	Code       string
	Message    string
	NextCursor int
	RequestId  string
	Action     string
	TimeLines  ListTimeLinesTimeLineList
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

func (c *CloudphotoClient) GetThumbnails(req *GetThumbnailsArgs) (resp *GetThumbnailsResponse, err error) {
	resp = &GetThumbnailsResponse{}
	err = c.Request("GetThumbnails", req, resp)
	return
}

type GetThumbnailsResult struct {
	Code         string
	Message      string
	PhotoId      int64
	ThumbnailUrl string
}
type GetThumbnailsArgs struct {
	LibraryId string
	PhotoIds  GetThumbnailsPhotoIdList
	StoreName string
	ZoomType  string
}
type GetThumbnailsResponse struct {
	Code      string
	Message   string
	RequestId string
	Action    string
	Results   GetThumbnailsResultList
}

type GetThumbnailsPhotoIdList []int64

func (list *GetThumbnailsPhotoIdList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]int64)
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

type GetThumbnailsResultList []GetThumbnailsResult

func (list *GetThumbnailsResultList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]GetThumbnailsResult)
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

func (c *CloudphotoClient) SetAlbumCover(req *SetAlbumCoverArgs) (resp *SetAlbumCoverResponse, err error) {
	resp = &SetAlbumCoverResponse{}
	err = c.Request("SetAlbumCover", req, resp)
	return
}

type SetAlbumCoverArgs struct {
	LibraryId string
	AlbumId   int64
	PhotoId   int64
	StoreName string
}
type SetAlbumCoverResponse struct {
	Code      string
	Message   string
	RequestId string
	Action    string
}

func (c *CloudphotoClient) ListAlbums(req *ListAlbumsArgs) (resp *ListAlbumsResponse, err error) {
	resp = &ListAlbumsResponse{}
	err = c.Request("ListAlbums", req, resp)
	return
}

type ListAlbumsAlbum struct {
	Id          int64
	Name        string
	State       string
	PhotosCount int64
	Ctime       int64
	Mtime       int64
	Cover       ListAlbumsCover
}

type ListAlbumsCover struct {
	Id      int64
	Title   string
	FileId  string
	State   string
	Md5     string
	IsVideo bool
	Remark  string
	Width   int64
	Height  int64
	Ctime   int64
	Mtime   int64
}
type ListAlbumsArgs struct {
	Cursor    string
	Size      int
	LibraryId string
	StoreName string
	State     string
	Direction string
}
type ListAlbumsResponse struct {
	Code       string
	Message    string
	NextCursor string
	TotalCount int
	RequestId  string
	Action     string
	Albums     ListAlbumsAlbumList
}

type ListAlbumsAlbumList []ListAlbumsAlbum

func (list *ListAlbumsAlbumList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListAlbumsAlbum)
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

func (c *CloudphotoClient) DeleteFaces(req *DeleteFacesArgs) (resp *DeleteFacesResponse, err error) {
	resp = &DeleteFacesResponse{}
	err = c.Request("DeleteFaces", req, resp)
	return
}

type DeleteFacesResult struct {
	Id      int64
	Code    string
	Message string
}
type DeleteFacesArgs struct {
	LibraryId string
	StoreName string
	FaceIds   DeleteFacesFaceIdList
}
type DeleteFacesResponse struct {
	Code      string
	Message   string
	RequestId string
	Action    string
	Results   DeleteFacesResultList
}

type DeleteFacesFaceIdList []int64

func (list *DeleteFacesFaceIdList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]int64)
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

type DeleteFacesResultList []DeleteFacesResult

func (list *DeleteFacesResultList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DeleteFacesResult)
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

func (c *CloudphotoClient) SearchPhotos(req *SearchPhotosArgs) (resp *SearchPhotosResponse, err error) {
	resp = &SearchPhotosResponse{}
	err = c.Request("SearchPhotos", req, resp)
	return
}

type SearchPhotosPhoto struct {
	Id              int64
	Title           string
	FileId          string
	State           string
	Md5             string
	IsVideo         bool
	Width           int64
	Height          int64
	Ctime           int64
	Mtime           int64
	TakenAt         int64
	ShareExpireTime int64
}
type SearchPhotosArgs struct {
	Size      int
	LibraryId string
	StoreName string
	Page      int
	Keyword   string
}
type SearchPhotosResponse struct {
	Code       string
	Message    string
	TotalCount int
	RequestId  string
	Action     string
	Photos     SearchPhotosPhotoList
}

type SearchPhotosPhotoList []SearchPhotosPhoto

func (list *SearchPhotosPhotoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]SearchPhotosPhoto)
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

func (c *CloudphotoClient) RemoveAlbumPhotos(req *RemoveAlbumPhotosArgs) (resp *RemoveAlbumPhotosResponse, err error) {
	resp = &RemoveAlbumPhotosResponse{}
	err = c.Request("RemoveAlbumPhotos", req, resp)
	return
}

type RemoveAlbumPhotosResult struct {
	Id      int64
	Code    string
	Message string
}
type RemoveAlbumPhotosArgs struct {
	LibraryId string
	AlbumId   int64
	PhotoIds  RemoveAlbumPhotosPhotoIdList
	StoreName string
}
type RemoveAlbumPhotosResponse struct {
	Code      string
	Message   string
	RequestId string
	Action    string
	Results   RemoveAlbumPhotosResultList
}

type RemoveAlbumPhotosPhotoIdList []int64

func (list *RemoveAlbumPhotosPhotoIdList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]int64)
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

type RemoveAlbumPhotosResultList []RemoveAlbumPhotosResult

func (list *RemoveAlbumPhotosResultList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]RemoveAlbumPhotosResult)
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

func (c *CloudphotoClient) GetVideoCover(req *GetVideoCoverArgs) (resp *GetVideoCoverResponse, err error) {
	resp = &GetVideoCoverResponse{}
	err = c.Request("GetVideoCover", req, resp)
	return
}

type GetVideoCoverArgs struct {
	LibraryId string
	PhotoId   int64
	StoreName string
	ZoomType  string
}
type GetVideoCoverResponse struct {
	Code          string
	Message       string
	VideoCoverUrl string
	RequestId     string
	Action        string
}

func (c *CloudphotoClient) EditPhotoStore(req *EditPhotoStoreArgs) (resp *EditPhotoStoreResponse, err error) {
	resp = &EditPhotoStoreResponse{}
	err = c.Request("EditPhotoStore", req, resp)
	return
}

type EditPhotoStoreArgs struct {
	AutoCleanEnabled string
	StoreName        string
	Remark           string
	DefaultQuota     int64
	AutoCleanDays    int
}
type EditPhotoStoreResponse struct {
	Code      string
	Message   string
	RequestId string
	Action    string
}

func (c *CloudphotoClient) CreatePhoto(req *CreatePhotoArgs) (resp *CreatePhotoResponse, err error) {
	resp = &CreatePhotoResponse{}
	err = c.Request("CreatePhoto", req, resp)
	return
}

type CreatePhotoPhoto struct {
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
}
type CreatePhotoArgs struct {
	PhotoTitle      string
	LibraryId       string
	ShareExpireTime int64
	StoreName       string
	UploadType      string
	Remark          string
	SessionId       string
	Staging         string
	FileId          string
}
type CreatePhotoResponse struct {
	Code      string
	Message   string
	RequestId string
	Action    string
	Photo     CreatePhotoPhoto
}

func (c *CloudphotoClient) MergeFaces(req *MergeFacesArgs) (resp *MergeFacesResponse, err error) {
	resp = &MergeFacesResponse{}
	err = c.Request("MergeFaces", req, resp)
	return
}

type MergeFacesResult struct {
	Id      int64
	Code    string
	Message string
}
type MergeFacesArgs struct {
	LibraryId    string
	TargetFaceId int64
	StoreName    string
	FaceIds      MergeFacesFaceIdList
}
type MergeFacesResponse struct {
	Code      string
	Message   string
	RequestId string
	Action    string
	Results   MergeFacesResultList
}

type MergeFacesFaceIdList []int64

func (list *MergeFacesFaceIdList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]int64)
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

type MergeFacesResultList []MergeFacesResult

func (list *MergeFacesResultList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]MergeFacesResult)
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

func (c *CloudphotoClient) CreateTransaction(req *CreateTransactionArgs) (resp *CreateTransactionResponse, err error) {
	resp = &CreateTransactionResponse{}
	err = c.Request("CreateTransaction", req, resp)
	return
}

type CreateTransactionTransaction struct {
	Upload CreateTransactionUpload
}

type CreateTransactionUpload struct {
	Bucket          string
	FileId          string
	OssEndpoint     string
	ObjectKey       string
	SessionId       string
	AccessKeyId     string
	AccessKeySecret string
	StsToken        string
}
type CreateTransactionArgs struct {
	Ext       string
	Size      int64
	LibraryId string
	StoreName string
	Force     string
	Md5       string
}
type CreateTransactionResponse struct {
	Code        string
	Message     string
	RequestId   string
	Action      string
	Transaction CreateTransactionTransaction
}

func (c *CloudphotoClient) SetQuota(req *SetQuotaArgs) (resp *SetQuotaResponse, err error) {
	resp = &SetQuotaResponse{}
	err = c.Request("SetQuota", req, resp)
	return
}

type SetQuotaArgs struct {
	TotalQuota int64
	LibraryId  string
	StoreName  string
}
type SetQuotaResponse struct {
	Code      string
	Message   string
	RequestId string
	Action    string
}

func (c *CloudphotoClient) TagPhoto(req *TagPhotoArgs) (resp *TagPhotoResponse, err error) {
	resp = &TagPhotoResponse{}
	err = c.Request("TagPhoto", req, resp)
	return
}

type TagPhotoArgs struct {
	LibraryId   string
	Confidences TagPhotoConfidenceList
	StoreName   string
	PhotoId     int64
	TagKeys     TagPhotoTagKeyList
}
type TagPhotoResponse struct {
	Code      string
	Message   string
	RequestId string
	Action    string
}

type TagPhotoConfidenceList []float32

func (list *TagPhotoConfidenceList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]float32)
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

type TagPhotoTagKeyList []string

func (list *TagPhotoTagKeyList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]string)
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

func (c *CloudphotoClient) CreatePhotoStore(req *CreatePhotoStoreArgs) (resp *CreatePhotoStoreResponse, err error) {
	resp = &CreatePhotoStoreResponse{}
	err = c.Request("CreatePhotoStore", req, resp)
	return
}

type CreatePhotoStoreArgs struct {
	BucketName   string
	StoreName    string
	Remark       string
	DefaultQuota int64
}
type CreatePhotoStoreResponse struct {
	Code      string
	Message   string
	RequestId string
	Action    string
}

func (c *CloudphotoClient) DeleteAlbums(req *DeleteAlbumsArgs) (resp *DeleteAlbumsResponse, err error) {
	resp = &DeleteAlbumsResponse{}
	err = c.Request("DeleteAlbums", req, resp)
	return
}

type DeleteAlbumsResult struct {
	Id      int64
	Code    string
	Message string
}
type DeleteAlbumsArgs struct {
	LibraryId string
	AlbumIds  DeleteAlbumsAlbumIdList
	StoreName string
}
type DeleteAlbumsResponse struct {
	Code      string
	Message   string
	RequestId string
	Action    string
	Results   DeleteAlbumsResultList
}

type DeleteAlbumsAlbumIdList []int64

func (list *DeleteAlbumsAlbumIdList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]int64)
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

type DeleteAlbumsResultList []DeleteAlbumsResult

func (list *DeleteAlbumsResultList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DeleteAlbumsResult)
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

func (c *CloudphotoClient) LikePhoto(req *LikePhotoArgs) (resp *LikePhotoResponse, err error) {
	resp = &LikePhotoResponse{}
	err = c.Request("LikePhoto", req, resp)
	return
}

type LikePhotoArgs struct {
	LibraryId string
	PhotoId   int64
	StoreName string
}
type LikePhotoResponse struct {
	Code      string
	Message   string
	RequestId string
	Action    string
}

func (c *CloudphotoClient) CreateAlbum(req *CreateAlbumArgs) (resp *CreateAlbumResponse, err error) {
	resp = &CreateAlbumResponse{}
	err = c.Request("CreateAlbum", req, resp)
	return
}

type CreateAlbumAlbum struct {
	Id          int64
	Name        string
	State       string
	Remark      string
	PhotosCount int64
	Ctime       int64
	Mtime       int64
	Cover       CreateAlbumCover
}

type CreateAlbumCover struct {
	Id      int64
	Title   string
	FileId  string
	State   string
	Md5     string
	IsVideo bool
	Width   int64
	Height  int64
	Ctime   int64
	Mtime   int64
	Remark  string
}
type CreateAlbumArgs struct {
	AlbumName string
	LibraryId string
	StoreName string
	Remark    string
}
type CreateAlbumResponse struct {
	Code      string
	Message   string
	RequestId string
	Action    string
	Album     CreateAlbumAlbum
}

func (c *CloudphotoClient) GetLibrary(req *GetLibraryArgs) (resp *GetLibraryResponse, err error) {
	resp = &GetLibraryResponse{}
	err = c.Request("GetLibrary", req, resp)
	return
}

type GetLibraryLibrary struct {
	Quota           GetLibraryQuota
	AutoCleanConfig GetLibraryAutoCleanConfig
}

type GetLibraryQuota struct {
	TotalQuota  int64
	FacesCount  int
	PhotosCount int
	UsedQuota   int64
	VideosCount int
}

type GetLibraryAutoCleanConfig struct {
	AutoCleanEnabled bool
	AutoCleanDays    int
}
type GetLibraryArgs struct {
	LibraryId string
	StoreName string
}
type GetLibraryResponse struct {
	Code      string
	Message   string
	RequestId string
	Action    string
	Library   GetLibraryLibrary
}

func (c *CloudphotoClient) ListRegisteredTags(req *ListRegisteredTagsArgs) (resp *ListRegisteredTagsResponse, err error) {
	resp = &ListRegisteredTagsResponse{}
	err = c.Request("ListRegisteredTags", req, resp)
	return
}

type ListRegisteredTagsRegisteredTag struct {
	TagKey    string
	TagValues ListRegisteredTagsTagValueList
}

type ListRegisteredTagsTagValue struct {
	Lang string
	Text string
}
type ListRegisteredTagsArgs struct {
	StoreName string
	Langs     ListRegisteredTagsLangList
}
type ListRegisteredTagsResponse struct {
	Code           string
	Message        string
	RequestId      string
	Action         string
	RegisteredTags ListRegisteredTagsRegisteredTagList
}

type ListRegisteredTagsTagValueList []ListRegisteredTagsTagValue

func (list *ListRegisteredTagsTagValueList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListRegisteredTagsTagValue)
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

type ListRegisteredTagsLangList []string

func (list *ListRegisteredTagsLangList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]string)
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

type ListRegisteredTagsRegisteredTagList []ListRegisteredTagsRegisteredTag

func (list *ListRegisteredTagsRegisteredTagList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListRegisteredTagsRegisteredTag)
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

func (c *CloudphotoClient) ListPhotoStores(req *ListPhotoStoresArgs) (resp *ListPhotoStoresResponse, err error) {
	resp = &ListPhotoStoresResponse{}
	err = c.Request("ListPhotoStores", req, resp)
	return
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
type ListPhotoStoresArgs struct {
}
type ListPhotoStoresResponse struct {
	Code        string
	Message     string
	RequestId   string
	Action      string
	PhotoStores ListPhotoStoresPhotoStoreList
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

func (c *CloudphotoClient) MoveFacePhotos(req *MoveFacePhotosArgs) (resp *MoveFacePhotosResponse, err error) {
	resp = &MoveFacePhotosResponse{}
	err = c.Request("MoveFacePhotos", req, resp)
	return
}

type MoveFacePhotosResult struct {
	Id      int64
	Code    string
	Message string
}
type MoveFacePhotosArgs struct {
	LibraryId    string
	TargetFaceId int64
	PhotoIds     MoveFacePhotosPhotoIdList
	StoreName    string
	SourceFaceId int64
}
type MoveFacePhotosResponse struct {
	Code      string
	Message   string
	RequestId string
	Action    string
	Results   MoveFacePhotosResultList
}

type MoveFacePhotosPhotoIdList []int64

func (list *MoveFacePhotosPhotoIdList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]int64)
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

type MoveFacePhotosResultList []MoveFacePhotosResult

func (list *MoveFacePhotosResultList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]MoveFacePhotosResult)
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

func (c *CloudphotoClient) ListTags(req *ListTagsArgs) (resp *ListTagsResponse, err error) {
	resp = &ListTagsResponse{}
	err = c.Request("ListTags", req, resp)
	return
}

type ListTagsTag struct {
	Id        int64
	Name      string
	IsSubTag  bool
	ParentTag string
	Cover     ListTagsCover
}

type ListTagsCover struct {
	Id      int64
	Title   string
	FileId  string
	State   string
	Md5     string
	IsVideo bool
	Remark  string
	Width   int64
	Height  int64
	Ctime   int64
	Mtime   int64
}
type ListTagsArgs struct {
	LibraryId string
	StoreName string
	Lang      string
}
type ListTagsResponse struct {
	Code      string
	Message   string
	RequestId string
	Action    string
	Tags      ListTagsTagList
}

type ListTagsTagList []ListTagsTag

func (list *ListTagsTagList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListTagsTag)
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

func (c *CloudphotoClient) ActivatePhotos(req *ActivatePhotosArgs) (resp *ActivatePhotosResponse, err error) {
	resp = &ActivatePhotosResponse{}
	err = c.Request("ActivatePhotos", req, resp)
	return
}

type ActivatePhotosResult struct {
	Id      int64
	Code    string
	Message string
}
type ActivatePhotosArgs struct {
	LibraryId string
	PhotoIds  ActivatePhotosPhotoIdList
	StoreName string
}
type ActivatePhotosResponse struct {
	Code      string
	Message   string
	RequestId string
	Action    string
	Results   ActivatePhotosResultList
}

type ActivatePhotosPhotoIdList []int64

func (list *ActivatePhotosPhotoIdList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]int64)
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

type ActivatePhotosResultList []ActivatePhotosResult

func (list *ActivatePhotosResultList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ActivatePhotosResult)
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

func (c *CloudphotoClient) InactivatePhotos(req *InactivatePhotosArgs) (resp *InactivatePhotosResponse, err error) {
	resp = &InactivatePhotosResponse{}
	err = c.Request("InactivatePhotos", req, resp)
	return
}

type InactivatePhotosResult struct {
	Id      int64
	Code    string
	Message string
}
type InactivatePhotosArgs struct {
	LibraryId    string
	PhotoIds     InactivatePhotosPhotoIdList
	StoreName    string
	InactiveTime int64
}
type InactivatePhotosResponse struct {
	Code      string
	Message   string
	RequestId string
	Action    string
	Results   InactivatePhotosResultList
}

type InactivatePhotosPhotoIdList []int64

func (list *InactivatePhotosPhotoIdList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]int64)
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

type InactivatePhotosResultList []InactivatePhotosResult

func (list *InactivatePhotosResultList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]InactivatePhotosResult)
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

func (c *CloudphotoClient) GetQuota(req *GetQuotaArgs) (resp *GetQuotaResponse, err error) {
	resp = &GetQuotaResponse{}
	err = c.Request("GetQuota", req, resp)
	return
}

type GetQuotaQuota struct {
	TotalQuota  int64
	FacesCount  int
	PhotosCount int
	UsedQuota   int64
	VideosCount int
}
type GetQuotaArgs struct {
	LibraryId string
	StoreName string
}
type GetQuotaResponse struct {
	Code      string
	Message   string
	RequestId string
	Action    string
	Quota     GetQuotaQuota
}

func (c *CloudphotoClient) MoveAlbumPhotos(req *MoveAlbumPhotosArgs) (resp *MoveAlbumPhotosResponse, err error) {
	resp = &MoveAlbumPhotosResponse{}
	err = c.Request("MoveAlbumPhotos", req, resp)
	return
}

type MoveAlbumPhotosResult struct {
	Id      int64
	Code    string
	Message string
}
type MoveAlbumPhotosArgs struct {
	SourceAlbumId int64
	TargetAlbumId int64
	LibraryId     string
	PhotoIds      MoveAlbumPhotosPhotoIdList
	StoreName     string
}
type MoveAlbumPhotosResponse struct {
	Code      string
	Message   string
	RequestId string
	Action    string
	Results   MoveAlbumPhotosResultList
}

type MoveAlbumPhotosPhotoIdList []int64

func (list *MoveAlbumPhotosPhotoIdList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]int64)
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

type MoveAlbumPhotosResultList []MoveAlbumPhotosResult

func (list *MoveAlbumPhotosResultList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]MoveAlbumPhotosResult)
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

func (c *CloudphotoClient) GetFramedPhotoUrls(req *GetFramedPhotoUrlsArgs) (resp *GetFramedPhotoUrlsResponse, err error) {
	resp = &GetFramedPhotoUrlsResponse{}
	err = c.Request("GetFramedPhotoUrls", req, resp)
	return
}

type GetFramedPhotoUrlsResult struct {
	Code           string
	Message        string
	PhotoId        int64
	FramedPhotoUrl string
}
type GetFramedPhotoUrlsArgs struct {
	FrameId   string
	LibraryId string
	PhotoIds  GetFramedPhotoUrlsPhotoIdList
	StoreName string
}
type GetFramedPhotoUrlsResponse struct {
	Code      string
	Message   string
	RequestId string
	Action    string
	Results   GetFramedPhotoUrlsResultList
}

type GetFramedPhotoUrlsPhotoIdList []int64

func (list *GetFramedPhotoUrlsPhotoIdList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]int64)
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

type GetFramedPhotoUrlsResultList []GetFramedPhotoUrlsResult

func (list *GetFramedPhotoUrlsResultList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]GetFramedPhotoUrlsResult)
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

func (c *CloudphotoClient) ListTimeLinePhotos(req *ListTimeLinePhotosArgs) (resp *ListTimeLinePhotosResponse, err error) {
	resp = &ListTimeLinePhotosResponse{}
	err = c.Request("ListTimeLinePhotos", req, resp)
	return
}

type ListTimeLinePhotosPhoto struct {
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
type ListTimeLinePhotosArgs struct {
	Size      int
	LibraryId string
	EndTime   int64
	StoreName string
	Page      int
	StartTime int64
	FilterBy  string
	Direction string
	Order     string
}
type ListTimeLinePhotosResponse struct {
	Code       string
	Message    string
	TotalCount int
	RequestId  string
	Action     string
	Photos     ListTimeLinePhotosPhotoList
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

func (c *CloudphotoClient) ListPhotos(req *ListPhotosArgs) (resp *ListPhotosResponse, err error) {
	resp = &ListPhotosResponse{}
	err = c.Request("ListPhotos", req, resp)
	return
}

type ListPhotosPhoto struct {
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
	InactiveTime    int64
	ShareExpireTime int64
}
type ListPhotosArgs struct {
	Cursor    string
	Size      int
	LibraryId string
	StoreName string
	State     string
	Direction string
}
type ListPhotosResponse struct {
	Code       string
	Message    string
	NextCursor string
	TotalCount int
	RequestId  string
	Action     string
	Photos     ListPhotosPhotoList
}

type ListPhotosPhotoList []ListPhotosPhoto

func (list *ListPhotosPhotoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListPhotosPhoto)
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

func (c *CloudphotoClient) GetPrivateAccessUrls(req *GetPrivateAccessUrlsArgs) (resp *GetPrivateAccessUrlsResponse, err error) {
	resp = &GetPrivateAccessUrlsResponse{}
	err = c.Request("GetPrivateAccessUrls", req, resp)
	return
}

type GetPrivateAccessUrlsResult struct {
	Code      string
	Message   string
	PhotoId   int64
	AccessUrl string
}
type GetPrivateAccessUrlsArgs struct {
	LibraryId string
	PhotoIds  GetPrivateAccessUrlsPhotoIdList
	StoreName string
	ZoomType  string
}
type GetPrivateAccessUrlsResponse struct {
	Code      string
	Message   string
	RequestId string
	Action    string
	Results   GetPrivateAccessUrlsResultList
}

type GetPrivateAccessUrlsPhotoIdList []int64

func (list *GetPrivateAccessUrlsPhotoIdList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]int64)
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

type GetPrivateAccessUrlsResultList []GetPrivateAccessUrlsResult

func (list *GetPrivateAccessUrlsResultList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]GetPrivateAccessUrlsResult)
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

func (c *CloudphotoClient) RenameAlbum(req *RenameAlbumArgs) (resp *RenameAlbumResponse, err error) {
	resp = &RenameAlbumResponse{}
	err = c.Request("RenameAlbum", req, resp)
	return
}

type RenameAlbumArgs struct {
	AlbumName string
	LibraryId string
	AlbumId   int64
	StoreName string
}
type RenameAlbumResponse struct {
	Code      string
	Message   string
	RequestId string
	Action    string
}

func (c *CloudphotoClient) GetThumbnail(req *GetThumbnailArgs) (resp *GetThumbnailResponse, err error) {
	resp = &GetThumbnailResponse{}
	err = c.Request("GetThumbnail", req, resp)
	return
}

type GetThumbnailArgs struct {
	LibraryId string
	PhotoId   int64
	StoreName string
	ZoomType  string
}
type GetThumbnailResponse struct {
	Code         string
	Message      string
	ThumbnailUrl string
	RequestId    string
	Action       string
}

func (c *CloudphotoClient) GetPhotos(req *GetPhotosArgs) (resp *GetPhotosResponse, err error) {
	resp = &GetPhotosResponse{}
	err = c.Request("GetPhotos", req, resp)
	return
}

type GetPhotosPhoto struct {
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
	InactiveTime    int64
	ShareExpireTime int64
	Like            int64
}
type GetPhotosArgs struct {
	LibraryId string
	PhotoIds  GetPhotosPhotoIdList
	StoreName string
}
type GetPhotosResponse struct {
	Code      string
	Message   string
	RequestId string
	Action    string
	Photos    GetPhotosPhotoList
}

type GetPhotosPhotoIdList []int64

func (list *GetPhotosPhotoIdList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]int64)
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

type GetPhotosPhotoList []GetPhotosPhoto

func (list *GetPhotosPhotoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]GetPhotosPhoto)
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

func (c *CloudphotoClient) GetPhotoStore(req *GetPhotoStoreArgs) (resp *GetPhotoStoreResponse, err error) {
	resp = &GetPhotoStoreResponse{}
	err = c.Request("GetPhotoStore", req, resp)
	return
}

type GetPhotoStorePhotoStore struct {
	Id               int64
	Name             string
	Remark           string
	AutoCleanEnabled bool
	AutoCleanDays    int
	DefaultQuota     int64
	Ctime            int64
	Mtime            int64
	Buckets          GetPhotoStoreBucketList
}

type GetPhotoStoreBucket struct {
	Name   string
	Region string
	State  string
	Acl    string
}
type GetPhotoStoreArgs struct {
	StoreName string
}
type GetPhotoStoreResponse struct {
	Code       string
	Message    string
	RequestId  string
	Action     string
	PhotoStore GetPhotoStorePhotoStore
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

func (c *CloudphotoClient) ListFacePhotos(req *ListFacePhotosArgs) (resp *ListFacePhotosResponse, err error) {
	resp = &ListFacePhotosResponse{}
	err = c.Request("ListFacePhotos", req, resp)
	return
}

type ListFacePhotosResult struct {
	PhotoId int64
	State   string
}
type ListFacePhotosArgs struct {
	Cursor    string
	Size      int
	LibraryId string
	StoreName string
	FaceId    int64
	State     string
	Direction string
}
type ListFacePhotosResponse struct {
	Code       string
	Message    string
	NextCursor string
	TotalCount int
	RequestId  string
	Action     string
	Results    ListFacePhotosResultList
}

type ListFacePhotosResultList []ListFacePhotosResult

func (list *ListFacePhotosResultList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListFacePhotosResult)
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

func (c *CloudphotoClient) GetPublicAccessUrls(req *GetPublicAccessUrlsArgs) (resp *GetPublicAccessUrlsResponse, err error) {
	resp = &GetPublicAccessUrlsResponse{}
	err = c.Request("GetPublicAccessUrls", req, resp)
	return
}

type GetPublicAccessUrlsResult struct {
	Code      string
	Message   string
	PhotoId   int64
	AccessUrl string
}
type GetPublicAccessUrlsArgs struct {
	DomainType string
	LibraryId  string
	PhotoIds   GetPublicAccessUrlsPhotoIdList
	StoreName  string
	ZoomType   string
}
type GetPublicAccessUrlsResponse struct {
	Code      string
	Message   string
	RequestId string
	Action    string
	Results   GetPublicAccessUrlsResultList
}

type GetPublicAccessUrlsPhotoIdList []int64

func (list *GetPublicAccessUrlsPhotoIdList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]int64)
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

type GetPublicAccessUrlsResultList []GetPublicAccessUrlsResult

func (list *GetPublicAccessUrlsResultList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]GetPublicAccessUrlsResult)
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

func (c *CloudphotoClient) ListFaces(req *ListFacesArgs) (resp *ListFacesResponse, err error) {
	resp = &ListFacesResponse{}
	err = c.Request("ListFaces", req, resp)
	return
}

type ListFacesFace struct {
	Id          int64
	Name        string
	PhotosCount int
	State       string
	IsMe        bool
	Ctime       int64
	Mtime       int64
	Axis        ListFacesAxiList
	Cover       ListFacesCover
}

type ListFacesCover struct {
	Id      int64
	Title   string
	FileId  string
	State   string
	Md5     string
	IsVideo bool
	Width   int64
	Height  int64
	Ctime   int64
	Mtime   int64
	Remark  string
}
type ListFacesArgs struct {
	Cursor      string
	HasFaceName string
	Size        int
	LibraryId   string
	StoreName   string
	State       string
	Direction   string
}
type ListFacesResponse struct {
	Code       string
	Message    string
	NextCursor string
	TotalCount int
	RequestId  string
	Action     string
	Faces      ListFacesFaceList
}

type ListFacesAxiList []string

func (list *ListFacesAxiList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]string)
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

type ListFacesFaceList []ListFacesFace

func (list *ListFacesFaceList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListFacesFace)
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

func (c *CloudphotoClient) DeletePhotoStore(req *DeletePhotoStoreArgs) (resp *DeletePhotoStoreResponse, err error) {
	resp = &DeletePhotoStoreResponse{}
	err = c.Request("DeletePhotoStore", req, resp)
	return
}

type DeletePhotoStoreArgs struct {
	StoreName string
}
type DeletePhotoStoreResponse struct {
	Code      string
	Message   string
	RequestId string
	Action    string
}

func (c *CloudphotoClient) RemoveFacePhotos(req *RemoveFacePhotosArgs) (resp *RemoveFacePhotosResponse, err error) {
	resp = &RemoveFacePhotosResponse{}
	err = c.Request("RemoveFacePhotos", req, resp)
	return
}

type RemoveFacePhotosResult struct {
	Id      int64
	Code    string
	Message string
}
type RemoveFacePhotosArgs struct {
	LibraryId string
	PhotoIds  RemoveFacePhotosPhotoIdList
	StoreName string
	FaceId    int64
}
type RemoveFacePhotosResponse struct {
	Code      string
	Message   string
	RequestId string
	Action    string
	Results   RemoveFacePhotosResultList
}

type RemoveFacePhotosPhotoIdList []int64

func (list *RemoveFacePhotosPhotoIdList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]int64)
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

type RemoveFacePhotosResultList []RemoveFacePhotosResult

func (list *RemoveFacePhotosResultList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]RemoveFacePhotosResult)
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

func (c *CloudphotoClient) RenameFace(req *RenameFaceArgs) (resp *RenameFaceResponse, err error) {
	resp = &RenameFaceResponse{}
	err = c.Request("RenameFace", req, resp)
	return
}

type RenameFaceArgs struct {
	LibraryId string
	StoreName string
	FaceId    int64
	FaceName  string
}
type RenameFaceResponse struct {
	Code      string
	Message   string
	RequestId string
	Action    string
}

func (c *CloudphotoClient) GetDownloadUrl(req *GetDownloadUrlArgs) (resp *GetDownloadUrlResponse, err error) {
	resp = &GetDownloadUrlResponse{}
	err = c.Request("GetDownloadUrl", req, resp)
	return
}

type GetDownloadUrlArgs struct {
	LibraryId string
	PhotoId   int64
	StoreName string
}
type GetDownloadUrlResponse struct {
	Code        string
	Message     string
	DownloadUrl string
	RequestId   string
	Action      string
}

func (c *CloudphotoClient) ToggleFeatures(req *ToggleFeaturesArgs) (resp *ToggleFeaturesResponse, err error) {
	resp = &ToggleFeaturesResponse{}
	err = c.Request("ToggleFeatures", req, resp)
	return
}

type ToggleFeaturesArgs struct {
	DisabledFeaturess ToggleFeaturesDisabledFeaturesList
	StoreName         string
	EnabledFeaturess  ToggleFeaturesEnabledFeaturesList
}
type ToggleFeaturesResponse struct {
	Code      string
	Message   string
	RequestId string
	Action    string
}

type ToggleFeaturesDisabledFeaturesList []string

func (list *ToggleFeaturesDisabledFeaturesList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]string)
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

type ToggleFeaturesEnabledFeaturesList []string

func (list *ToggleFeaturesEnabledFeaturesList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]string)
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

func (c *CloudphotoClient) ListMoments(req *ListMomentsArgs) (resp *ListMomentsResponse, err error) {
	resp = &ListMomentsResponse{}
	err = c.Request("ListMoments", req, resp)
	return
}

type ListMomentsMoment struct {
	Id           int64
	LocationName string
	PhotosCount  int
	State        string
	TakenAt      int64
	Ctime        int64
	Mtime        int64
}
type ListMomentsArgs struct {
	Cursor    string
	Size      int
	LibraryId string
	StoreName string
	State     string
	Direction string
}
type ListMomentsResponse struct {
	Code       string
	Message    string
	NextCursor string
	TotalCount int
	RequestId  string
	Action     string
	Moments    ListMomentsMomentList
}

type ListMomentsMomentList []ListMomentsMoment

func (list *ListMomentsMomentList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListMomentsMoment)
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

func (c *CloudphotoClient) DeletePhotos(req *DeletePhotosArgs) (resp *DeletePhotosResponse, err error) {
	resp = &DeletePhotosResponse{}
	err = c.Request("DeletePhotos", req, resp)
	return
}

type DeletePhotosResult struct {
	Id      int64
	Code    string
	Message string
}
type DeletePhotosArgs struct {
	LibraryId string
	PhotoIds  DeletePhotosPhotoIdList
	StoreName string
}
type DeletePhotosResponse struct {
	Code      string
	Message   string
	RequestId string
	Action    string
	Results   DeletePhotosResultList
}

type DeletePhotosPhotoIdList []int64

func (list *DeletePhotosPhotoIdList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]int64)
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

type DeletePhotosResultList []DeletePhotosResult

func (list *DeletePhotosResultList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DeletePhotosResult)
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

func (c *CloudphotoClient) ListTagPhotos(req *ListTagPhotosArgs) (resp *ListTagPhotosResponse, err error) {
	resp = &ListTagPhotosResponse{}
	err = c.Request("ListTagPhotos", req, resp)
	return
}

type ListTagPhotosResult struct {
	PhotoId int64
	State   string
}
type ListTagPhotosArgs struct {
	Cursor    string
	Size      int
	TagId     int64
	LibraryId string
	StoreName string
	State     string
	Direction string
}
type ListTagPhotosResponse struct {
	Code       string
	Message    string
	NextCursor string
	TotalCount int
	RequestId  string
	Action     string
	Results    ListTagPhotosResultList
}

type ListTagPhotosResultList []ListTagPhotosResult

func (list *ListTagPhotosResultList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListTagPhotosResult)
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

func (c *CloudphotoClient) ListAlbumPhotos(req *ListAlbumPhotosArgs) (resp *ListAlbumPhotosResponse, err error) {
	resp = &ListAlbumPhotosResponse{}
	err = c.Request("ListAlbumPhotos", req, resp)
	return
}

type ListAlbumPhotosResult struct {
	PhotoId int64
	State   string
}
type ListAlbumPhotosArgs struct {
	Cursor    string
	Size      int
	LibraryId string
	AlbumId   int64
	StoreName string
	State     string
	Direction string
}
type ListAlbumPhotosResponse struct {
	Code       string
	Message    string
	NextCursor string
	TotalCount int
	RequestId  string
	Action     string
	Results    ListAlbumPhotosResultList
}

type ListAlbumPhotosResultList []ListAlbumPhotosResult

func (list *ListAlbumPhotosResultList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListAlbumPhotosResult)
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

func (c *CloudphotoClient) ReactivatePhotos(req *ReactivatePhotosArgs) (resp *ReactivatePhotosResponse, err error) {
	resp = &ReactivatePhotosResponse{}
	err = c.Request("ReactivatePhotos", req, resp)
	return
}

type ReactivatePhotosResult struct {
	Id      int64
	Code    string
	Message string
}
type ReactivatePhotosArgs struct {
	LibraryId string
	PhotoIds  ReactivatePhotosPhotoIdList
	StoreName string
}
type ReactivatePhotosResponse struct {
	Code      string
	Message   string
	RequestId string
	Action    string
	Results   ReactivatePhotosResultList
}

type ReactivatePhotosPhotoIdList []int64

func (list *ReactivatePhotosPhotoIdList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]int64)
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

type ReactivatePhotosResultList []ReactivatePhotosResult

func (list *ReactivatePhotosResultList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ReactivatePhotosResult)
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

func (c *CloudphotoClient) AddAlbumPhotos(req *AddAlbumPhotosArgs) (resp *AddAlbumPhotosResponse, err error) {
	resp = &AddAlbumPhotosResponse{}
	err = c.Request("AddAlbumPhotos", req, resp)
	return
}

type AddAlbumPhotosResult struct {
	Id      int64
	Code    string
	Message string
}
type AddAlbumPhotosArgs struct {
	LibraryId string
	AlbumId   int64
	PhotoIds  AddAlbumPhotosPhotoIdList
	StoreName string
}
type AddAlbumPhotosResponse struct {
	Code      string
	Message   string
	RequestId string
	Action    string
	Results   AddAlbumPhotosResultList
}

type AddAlbumPhotosPhotoIdList []int64

func (list *AddAlbumPhotosPhotoIdList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]int64)
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

type AddAlbumPhotosResultList []AddAlbumPhotosResult

func (list *AddAlbumPhotosResultList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]AddAlbumPhotosResult)
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

func (c *CloudphotoClient) RegisterTag(req *RegisterTagArgs) (resp *RegisterTagResponse, err error) {
	resp = &RegisterTagResponse{}
	err = c.Request("RegisterTag", req, resp)
	return
}

type RegisterTagArgs struct {
	StoreName string
	Text      string
	TagKey    string
	Lang      string
}
type RegisterTagResponse struct {
	Code      string
	Message   string
	RequestId string
	Action    string
}

func (c *CloudphotoClient) ListPhotoTags(req *ListPhotoTagsArgs) (resp *ListPhotoTagsResponse, err error) {
	resp = &ListPhotoTagsResponse{}
	err = c.Request("ListPhotoTags", req, resp)
	return
}

type ListPhotoTagsTag struct {
	Id        int64
	IsSubTag  bool
	Name      string
	ParentTag string
}
type ListPhotoTagsArgs struct {
	LibraryId string
	PhotoId   int64
	StoreName string
	Lang      string
}
type ListPhotoTagsResponse struct {
	Code      string
	Message   string
	RequestId string
	Action    string
	Tags      ListPhotoTagsTagList
}

type ListPhotoTagsTagList []ListPhotoTagsTag

func (list *ListPhotoTagsTagList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListPhotoTagsTag)
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

func (c *CloudphotoClient) SetMe(req *SetMeArgs) (resp *SetMeResponse, err error) {
	resp = &SetMeResponse{}
	err = c.Request("SetMe", req, resp)
	return
}

type SetMeArgs struct {
	LibraryId string
	StoreName string
	FaceId    int64
}
type SetMeResponse struct {
	Code      string
	Message   string
	RequestId string
	Action    string
}

func (c *CloudphotoClient) GetDownloadUrls(req *GetDownloadUrlsArgs) (resp *GetDownloadUrlsResponse, err error) {
	resp = &GetDownloadUrlsResponse{}
	err = c.Request("GetDownloadUrls", req, resp)
	return
}

type GetDownloadUrlsResult struct {
	Code        string
	Message     string
	PhotoId     int64
	DownloadUrl string
}
type GetDownloadUrlsArgs struct {
	LibraryId string
	PhotoIds  GetDownloadUrlsPhotoIdList
	StoreName string
}
type GetDownloadUrlsResponse struct {
	Code      string
	Message   string
	RequestId string
	Action    string
	Results   GetDownloadUrlsResultList
}

type GetDownloadUrlsPhotoIdList []int64

func (list *GetDownloadUrlsPhotoIdList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]int64)
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

type GetDownloadUrlsResultList []GetDownloadUrlsResult

func (list *GetDownloadUrlsResultList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]GetDownloadUrlsResult)
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

func (c *CloudphotoClient) GetPhotosByMd5s(req *GetPhotosByMd5sArgs) (resp *GetPhotosByMd5sResponse, err error) {
	resp = &GetPhotosByMd5sResponse{}
	err = c.Request("GetPhotosByMd5s", req, resp)
	return
}

type GetPhotosByMd5sPhoto struct {
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
}
type GetPhotosByMd5sArgs struct {
	LibraryId string
	StoreName string
	State     string
	Md5s      GetPhotosByMd5sMd5List
}
type GetPhotosByMd5sResponse struct {
	Code      string
	Message   string
	RequestId string
	Action    string
	Photos    GetPhotosByMd5sPhotoList
}

type GetPhotosByMd5sMd5List []string

func (list *GetPhotosByMd5sMd5List) UnmarshalJSON(data []byte) error {
	m := make(map[string][]string)
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

type GetPhotosByMd5sPhotoList []GetPhotosByMd5sPhoto

func (list *GetPhotosByMd5sPhotoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]GetPhotosByMd5sPhoto)
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

func (c *CloudphotoClient) ListPhotoFaces(req *ListPhotoFacesArgs) (resp *ListPhotoFacesResponse, err error) {
	resp = &ListPhotoFacesResponse{}
	err = c.Request("ListPhotoFaces", req, resp)
	return
}

type ListPhotoFacesFace struct {
	FaceId   int64
	FaceName string
	Axis     ListPhotoFacesAxiList
}
type ListPhotoFacesArgs struct {
	LibraryId string
	PhotoId   int64
	StoreName string
}
type ListPhotoFacesResponse struct {
	Code      string
	Message   string
	RequestId string
	Action    string
	Faces     ListPhotoFacesFaceList
}

type ListPhotoFacesAxiList []string

func (list *ListPhotoFacesAxiList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]string)
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

type ListPhotoFacesFaceList []ListPhotoFacesFace

func (list *ListPhotoFacesFaceList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListPhotoFacesFace)
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

func (c *CloudphotoClient) ListMomentPhotos(req *ListMomentPhotosArgs) (resp *ListMomentPhotosResponse, err error) {
	resp = &ListMomentPhotosResponse{}
	err = c.Request("ListMomentPhotos", req, resp)
	return
}

type ListMomentPhotosResult struct {
	PhotoId int64
	State   string
}
type ListMomentPhotosArgs struct {
	Cursor    string
	Size      int
	LibraryId string
	StoreName string
	State     string
	MomentId  int64
	Direction string
}
type ListMomentPhotosResponse struct {
	Code       string
	Message    string
	NextCursor string
	TotalCount int
	RequestId  string
	Action     string
	Results    ListMomentPhotosResultList
}

type ListMomentPhotosResultList []ListMomentPhotosResult

func (list *ListMomentPhotosResultList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListMomentPhotosResult)
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

func (c *CloudphotoClient) SetFaceCover(req *SetFaceCoverArgs) (resp *SetFaceCoverResponse, err error) {
	resp = &SetFaceCoverResponse{}
	err = c.Request("SetFaceCover", req, resp)
	return
}

type SetFaceCoverArgs struct {
	LibraryId string
	PhotoId   int64
	StoreName string
	FaceId    int64
}
type SetFaceCoverResponse struct {
	Code      string
	Message   string
	RequestId string
	Action    string
}
