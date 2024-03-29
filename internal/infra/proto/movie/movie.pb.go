// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.23.1
// source: proto/movie.proto

package movie

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	actor "api-gateway/internal/infra/proto/actor"
	director "api-gateway/internal/infra/proto/director"
	genre "api-gateway/internal/infra/proto/genre"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Movie struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         uint64             `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Title      string             `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty" validate:"required"`
	Synopsis   string             `protobuf:"bytes,3,opt,name=synopsis,proto3" json:"synopsis,omitempty" validate:"required"`
	Year       uint32             `protobuf:"varint,4,opt,name=year,proto3" json:"year,omitempty" validate:"number"`
	Rating     float32            `protobuf:"fixed32,5,opt,name=rating,proto3" json:"rating,omitempty"`
	Duration   uint32             `protobuf:"varint,6,opt,name=duration,proto3" json:"duration,omitempty" validate:"number"`
	DirectorId uint64             `protobuf:"varint,7,opt,name=director_id,json=directorId,proto3" json:"director_id,omitempty" validate:"number"`
	State      bool               `protobuf:"varint,8,opt,name=state,proto3" json:"state,omitempty"`
	Director   *director.Director `protobuf:"bytes,9,opt,name=director,proto3" json:"director,omitempty"`
	Actors     []*actor.Actor     `protobuf:"bytes,10,rep,name=actors,proto3" json:"actors,omitempty"`
	Genres     []*genre.Genre     `protobuf:"bytes,11,rep,name=genres,proto3" json:"genres,omitempty"`
	CreatedAt  string             `protobuf:"bytes,12,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt  string             `protobuf:"bytes,13,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
}

func (x *Movie) Reset() {
	*x = Movie{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_movie_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Movie) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Movie) ProtoMessage() {}

func (x *Movie) ProtoReflect() protoreflect.Message {
	mi := &file_proto_movie_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Movie.ProtoReflect.Descriptor instead.
func (*Movie) Descriptor() ([]byte, []int) {
	return file_proto_movie_proto_rawDescGZIP(), []int{0}
}

func (x *Movie) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Movie) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Movie) GetSynopsis() string {
	if x != nil {
		return x.Synopsis
	}
	return ""
}

func (x *Movie) GetYear() uint32 {
	if x != nil {
		return x.Year
	}
	return 0
}

func (x *Movie) GetRating() float32 {
	if x != nil {
		return x.Rating
	}
	return 0
}

func (x *Movie) GetDuration() uint32 {
	if x != nil {
		return x.Duration
	}
	return 0
}

func (x *Movie) GetDirectorId() uint64 {
	if x != nil {
		return x.DirectorId
	}
	return 0
}

func (x *Movie) GetState() bool {
	if x != nil {
		return x.State
	}
	return false
}

func (x *Movie) GetDirector() *director.Director {
	if x != nil {
		return x.Director
	}
	return nil
}

func (x *Movie) GetActors() []*actor.Actor {
	if x != nil {
		return x.Actors
	}
	return nil
}

func (x *Movie) GetGenres() []*genre.Genre {
	if x != nil {
		return x.Genres
	}
	return nil
}

func (x *Movie) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *Movie) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

type FilterCriteriaMovie struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GenreName  string `protobuf:"bytes,1,opt,name=genre_name,json=genreName,proto3" json:"genre_name,omitempty"`
	MovieTitle string `protobuf:"bytes,2,opt,name=movie_title,json=movieTitle,proto3" json:"movie_title,omitempty"`
}

func (x *FilterCriteriaMovie) Reset() {
	*x = FilterCriteriaMovie{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_movie_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FilterCriteriaMovie) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FilterCriteriaMovie) ProtoMessage() {}

func (x *FilterCriteriaMovie) ProtoReflect() protoreflect.Message {
	mi := &file_proto_movie_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FilterCriteriaMovie.ProtoReflect.Descriptor instead.
func (*FilterCriteriaMovie) Descriptor() ([]byte, []int) {
	return file_proto_movie_proto_rawDescGZIP(), []int{1}
}

func (x *FilterCriteriaMovie) GetGenreName() string {
	if x != nil {
		return x.GenreName
	}
	return ""
}

func (x *FilterCriteriaMovie) GetMovieTitle() string {
	if x != nil {
		return x.MovieTitle
	}
	return ""
}

type ListRequestMovie struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Page           uint64               `protobuf:"varint,1,opt,name=page,proto3" json:"page,omitempty"`
	PageSize       uint64               `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	FilterCriteria *FilterCriteriaMovie `protobuf:"bytes,3,opt,name=filter_criteria,json=filterCriteria,proto3" json:"filter_criteria,omitempty"`
}

func (x *ListRequestMovie) Reset() {
	*x = ListRequestMovie{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_movie_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListRequestMovie) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListRequestMovie) ProtoMessage() {}

func (x *ListRequestMovie) ProtoReflect() protoreflect.Message {
	mi := &file_proto_movie_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListRequestMovie.ProtoReflect.Descriptor instead.
func (*ListRequestMovie) Descriptor() ([]byte, []int) {
	return file_proto_movie_proto_rawDescGZIP(), []int{2}
}

func (x *ListRequestMovie) GetPage() uint64 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *ListRequestMovie) GetPageSize() uint64 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *ListRequestMovie) GetFilterCriteria() *FilterCriteriaMovie {
	if x != nil {
		return x.FilterCriteria
	}
	return nil
}

type RequestByIdMovie struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *RequestByIdMovie) Reset() {
	*x = RequestByIdMovie{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_movie_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RequestByIdMovie) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequestByIdMovie) ProtoMessage() {}

func (x *RequestByIdMovie) ProtoReflect() protoreflect.Message {
	mi := &file_proto_movie_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RequestByIdMovie.ProtoReflect.Descriptor instead.
func (*RequestByIdMovie) Descriptor() ([]byte, []int) {
	return file_proto_movie_proto_rawDescGZIP(), []int{3}
}

func (x *RequestByIdMovie) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type ResponseMovie struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Title      string   `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Message    string   `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	IsOk       bool     `protobuf:"varint,3,opt,name=is_ok,json=isOk,proto3" json:"is_ok,omitempty"`
	Status     int32    `protobuf:"varint,4,opt,name=status,proto3" json:"status,omitempty"`
	Movie      *Movie   `protobuf:"bytes,5,opt,name=movie,proto3" json:"movie,omitempty"`
	Movies     []*Movie `protobuf:"bytes,6,rep,name=movies,proto3" json:"movies,omitempty"`
	TotalPages uint64   `protobuf:"varint,7,opt,name=total_pages,json=totalPages,proto3" json:"total_pages,omitempty"`
}

func (x *ResponseMovie) Reset() {
	*x = ResponseMovie{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_movie_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResponseMovie) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResponseMovie) ProtoMessage() {}

func (x *ResponseMovie) ProtoReflect() protoreflect.Message {
	mi := &file_proto_movie_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResponseMovie.ProtoReflect.Descriptor instead.
func (*ResponseMovie) Descriptor() ([]byte, []int) {
	return file_proto_movie_proto_rawDescGZIP(), []int{4}
}

func (x *ResponseMovie) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *ResponseMovie) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *ResponseMovie) GetIsOk() bool {
	if x != nil {
		return x.IsOk
	}
	return false
}

func (x *ResponseMovie) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *ResponseMovie) GetMovie() *Movie {
	if x != nil {
		return x.Movie
	}
	return nil
}

func (x *ResponseMovie) GetMovies() []*Movie {
	if x != nil {
		return x.Movies
	}
	return nil
}

func (x *ResponseMovie) GetTotalPages() uint64 {
	if x != nil {
		return x.TotalPages
	}
	return 0
}

var File_proto_movie_proto protoreflect.FileDescriptor

var file_proto_movie_proto_rawDesc = []byte{
	0x0a, 0x11, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6d, 0x6f, 0x76, 0x69, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x06, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x1a, 0x14, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2f, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x11, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x65, 0x6e, 0x72,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x82, 0x03, 0x0a, 0x05, 0x4d, 0x6f, 0x76, 0x69,
	0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x73, 0x79, 0x6e, 0x6f, 0x70,
	0x73, 0x69, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73, 0x79, 0x6e, 0x6f, 0x70,
	0x73, 0x69, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x79, 0x65, 0x61, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x04, 0x79, 0x65, 0x61, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x61, 0x74, 0x69, 0x6e,
	0x67, 0x18, 0x05, 0x20, 0x01, 0x28, 0x02, 0x52, 0x06, 0x72, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x12,
	0x1a, 0x0a, 0x08, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x08, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1f, 0x0a, 0x0b, 0x64,
	0x69, 0x72, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x04,
	0x52, 0x0a, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05,
	0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x73, 0x74, 0x61,
	0x74, 0x65, 0x12, 0x2c, 0x0a, 0x08, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x18, 0x09,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x69,
	0x72, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x52, 0x08, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x6f, 0x72,
	0x12, 0x25, 0x0a, 0x06, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x73, 0x18, 0x0a, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x0d, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x63, 0x74, 0x6f, 0x72, 0x52,
	0x06, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x73, 0x12, 0x25, 0x0a, 0x06, 0x67, 0x65, 0x6e, 0x72, 0x65,
	0x73, 0x18, 0x0b, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31,
	0x2e, 0x47, 0x65, 0x6e, 0x72, 0x65, 0x52, 0x06, 0x67, 0x65, 0x6e, 0x72, 0x65, 0x73, 0x12, 0x1d,
	0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x0c, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a,
	0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x0d, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0x55, 0x0a, 0x13,
	0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x43, 0x72, 0x69, 0x74, 0x65, 0x72, 0x69, 0x61, 0x4d, 0x6f,
	0x76, 0x69, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x67, 0x65, 0x6e, 0x72, 0x65, 0x5f, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x67, 0x65, 0x6e, 0x72, 0x65, 0x4e, 0x61,
	0x6d, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x6d, 0x6f, 0x76, 0x69, 0x65, 0x5f, 0x74, 0x69, 0x74, 0x6c,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x6d, 0x6f, 0x76, 0x69, 0x65, 0x54, 0x69,
	0x74, 0x6c, 0x65, 0x22, 0x89, 0x01, 0x0a, 0x10, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x12, 0x1b, 0x0a, 0x09,
	0x70, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x44, 0x0a, 0x0f, 0x66, 0x69, 0x6c,
	0x74, 0x65, 0x72, 0x5f, 0x63, 0x72, 0x69, 0x74, 0x65, 0x72, 0x69, 0x61, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x46, 0x69, 0x6c, 0x74,
	0x65, 0x72, 0x43, 0x72, 0x69, 0x74, 0x65, 0x72, 0x69, 0x61, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x52,
	0x0e, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x43, 0x72, 0x69, 0x74, 0x65, 0x72, 0x69, 0x61, 0x22,
	0x22, 0x0a, 0x10, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x42, 0x79, 0x49, 0x64, 0x4d, 0x6f,
	0x76, 0x69, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x02, 0x69, 0x64, 0x22, 0xd9, 0x01, 0x0a, 0x0d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x4d, 0x6f, 0x76, 0x69, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x13, 0x0a, 0x05, 0x69, 0x73, 0x5f, 0x6f, 0x6b, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x04, 0x69, 0x73, 0x4f, 0x6b, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x12, 0x23, 0x0a, 0x05, 0x6d, 0x6f, 0x76, 0x69, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x0d, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x6f, 0x76, 0x69, 0x65,
	0x52, 0x05, 0x6d, 0x6f, 0x76, 0x69, 0x65, 0x12, 0x25, 0x0a, 0x06, 0x6d, 0x6f, 0x76, 0x69, 0x65,
	0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31,
	0x2e, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x52, 0x06, 0x6d, 0x6f, 0x76, 0x69, 0x65, 0x73, 0x12, 0x1f,
	0x0a, 0x0b, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x70, 0x61, 0x67, 0x65, 0x73, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x0a, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x50, 0x61, 0x67, 0x65, 0x73, 0x32,
	0xa5, 0x02, 0x0a, 0x09, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x43, 0x72, 0x75, 0x64, 0x12, 0x30, 0x0a,
	0x06, 0x49, 0x6e, 0x73, 0x65, 0x72, 0x74, 0x12, 0x0d, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31,
	0x2e, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x1a, 0x15, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x22, 0x00, 0x12,
	0x30, 0x0a, 0x06, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x0d, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x76, 0x31, 0x2e, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x1a, 0x15, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76,
	0x31, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x22,
	0x00, 0x12, 0x39, 0x0a, 0x04, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x18, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x4d, 0x6f,
	0x76, 0x69, 0x65, 0x1a, 0x15, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x22, 0x00, 0x12, 0x3b, 0x0a, 0x06,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x18, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x42, 0x79, 0x49, 0x64, 0x4d, 0x6f, 0x76, 0x69, 0x65,
	0x1a, 0x15, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x22, 0x00, 0x12, 0x3c, 0x0a, 0x07, 0x47, 0x65, 0x74,
	0x42, 0x79, 0x49, 0x64, 0x12, 0x18, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x42, 0x79, 0x49, 0x64, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x1a, 0x15,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x4d, 0x6f, 0x76, 0x69, 0x65, 0x22, 0x00, 0x42, 0x27, 0x5a, 0x25, 0x67, 0x72, 0x70, 0x63, 0x2d,
	0x6d, 0x6f, 0x76, 0x69, 0x65, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x69,
	0x6e, 0x66, 0x72, 0x61, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6d, 0x6f, 0x76, 0x69, 0x65,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_movie_proto_rawDescOnce sync.Once
	file_proto_movie_proto_rawDescData = file_proto_movie_proto_rawDesc
)

func file_proto_movie_proto_rawDescGZIP() []byte {
	file_proto_movie_proto_rawDescOnce.Do(func() {
		file_proto_movie_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_movie_proto_rawDescData)
	})
	return file_proto_movie_proto_rawDescData
}

var file_proto_movie_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_proto_movie_proto_goTypes = []interface{}{
	(*Movie)(nil),               // 0: api.v1.Movie
	(*FilterCriteriaMovie)(nil), // 1: api.v1.FilterCriteriaMovie
	(*ListRequestMovie)(nil),    // 2: api.v1.ListRequestMovie
	(*RequestByIdMovie)(nil),    // 3: api.v1.RequestByIdMovie
	(*ResponseMovie)(nil),       // 4: api.v1.ResponseMovie
	(*director.Director)(nil),   // 5: api.v1.Director
	(*actor.Actor)(nil),         // 6: api.v1.Actor
	(*genre.Genre)(nil),         // 7: api.v1.Genre
}
var file_proto_movie_proto_depIdxs = []int32{
	5,  // 0: api.v1.Movie.director:type_name -> api.v1.Director
	6,  // 1: api.v1.Movie.actors:type_name -> api.v1.Actor
	7,  // 2: api.v1.Movie.genres:type_name -> api.v1.Genre
	1,  // 3: api.v1.ListRequestMovie.filter_criteria:type_name -> api.v1.FilterCriteriaMovie
	0,  // 4: api.v1.ResponseMovie.movie:type_name -> api.v1.Movie
	0,  // 5: api.v1.ResponseMovie.movies:type_name -> api.v1.Movie
	0,  // 6: api.v1.MovieCrud.Insert:input_type -> api.v1.Movie
	0,  // 7: api.v1.MovieCrud.Update:input_type -> api.v1.Movie
	2,  // 8: api.v1.MovieCrud.List:input_type -> api.v1.ListRequestMovie
	3,  // 9: api.v1.MovieCrud.Delete:input_type -> api.v1.RequestByIdMovie
	3,  // 10: api.v1.MovieCrud.GetById:input_type -> api.v1.RequestByIdMovie
	4,  // 11: api.v1.MovieCrud.Insert:output_type -> api.v1.ResponseMovie
	4,  // 12: api.v1.MovieCrud.Update:output_type -> api.v1.ResponseMovie
	4,  // 13: api.v1.MovieCrud.List:output_type -> api.v1.ResponseMovie
	4,  // 14: api.v1.MovieCrud.Delete:output_type -> api.v1.ResponseMovie
	4,  // 15: api.v1.MovieCrud.GetById:output_type -> api.v1.ResponseMovie
	11, // [11:16] is the sub-list for method output_type
	6,  // [6:11] is the sub-list for method input_type
	6,  // [6:6] is the sub-list for extension type_name
	6,  // [6:6] is the sub-list for extension extendee
	0,  // [0:6] is the sub-list for field type_name
}

func init() { file_proto_movie_proto_init() }
func file_proto_movie_proto_init() {
	if File_proto_movie_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_movie_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Movie); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_movie_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FilterCriteriaMovie); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_movie_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListRequestMovie); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_movie_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RequestByIdMovie); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_movie_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResponseMovie); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_movie_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_movie_proto_goTypes,
		DependencyIndexes: file_proto_movie_proto_depIdxs,
		MessageInfos:      file_proto_movie_proto_msgTypes,
	}.Build()
	File_proto_movie_proto = out.File
	file_proto_movie_proto_rawDesc = nil
	file_proto_movie_proto_goTypes = nil
	file_proto_movie_proto_depIdxs = nil
}
