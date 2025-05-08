package server

import (
	"context"
	"log"

	"summary-service/server/db"
	"summary-service/server/validators"

	pb "github.com/Azat201003/eduflow_service_api/gen/go/summary"
	"github.com/lib/pq"
	"gorm.io/gorm"
)

type summaryServiceServer struct {
	pb.UnimplementedSummaryServiceServer
	dbm       *db.DBManger
	validator *validators.Validator
}

func convertTagsFromIds(dbm *db.DBManger, ids pq.Int64Array) ([]*pb.Tag, error) {
	var tags []*pb.Tag
	for _, id := range ids {
		tag := &db.Tag{ID: uint64(id)}
		err := dbm.FindTag(tag)
		if err != nil {
			return nil, err
		}
		tags = append(tags, &pb.Tag{Id: &pb.Id{Id: uint64(id)}, Name: tag.Name, Color: tag.Color})
	}
	return tags, nil
}

func convertTagsToIds(tags []*pb.Tag) (pq.Int64Array, error) {
	var ids pq.Int64Array
	for _, tag := range tags {
		ids = append(ids, int64(tag.Id.Id))
	}
	return ids, nil
}

func (s *summaryServiceServer) GetSummaryById(context context.Context, id *pb.Id) (*pb.Summary, error) {
	summary := db.Summary{ID: id.Id}
	err := s.dbm.FindSummary(&summary)
	if err != nil {
		return nil, err
	}
	tags, err := convertTagsFromIds(s.dbm, summary.Tags)
	return &pb.Summary{
		Id:          &pb.Id{Id: summary.ID},
		Title:       summary.Title,
		Description: summary.Description,
		FilePath:    summary.FilePath,
		AuthorId:    &pb.Id{Id: summary.AuthorId},
		Tags:        tags,
	}, err
}

func (s *summaryServiceServer) CreateSummary(ctx context.Context, summary *pb.Summary) (*pb.Id, error) {
	var tag_ids pq.Int64Array
	for _, tag := range summary.Tags {
		tag_ids = append(tag_ids, int64(tag.Id.Id))
	}
	db_summary := db.Summary{
		Title:       summary.Title,
		Description: summary.Description,
		FilePath:    summary.FilePath,
		AuthorId:    summary.AuthorId.Id,
		Tags:        tag_ids,
	}
	err := s.dbm.CreateSummary(&db_summary)
	return &pb.Id{Id: db_summary.ID}, err
}

func (s *summaryServiceServer) GetFilteredSummaries(request *pb.FilterRequest, stream pb.SummaryService_GetFilteredSummariesServer) error {
	if err := s.validator.Validate(request); err != nil {
		return err
	}
	summary := request.Filter
	tag_ids, err := convertTagsToIds(request.Filter.Tags)
	conds := &db.Summary{
		Title:       summary.Title,
		Description: summary.Description,
		FilePath:    summary.FilePath,
		Tags:        tag_ids,
	}
	if summary.Id != nil {
		conds.ID = summary.Id.Id
	}
	if summary.AuthorId != nil {
		conds.AuthorId = summary.AuthorId.Id
	}
	summaries, err := s.dbm.FilteredSummaries(conds, int((request.Page.Number-1)*request.Page.Size), int(request.Page.Size))
	log.Println(summary)
	log.Println(summaries)
	if err != nil {
		return err
	}
	for _, summary := range *summaries {
		tags, err := convertTagsFromIds(s.dbm, summary.Tags)
		if err != nil {
			return err
		}
		if err := stream.Send(&pb.Summary{
			Id:          &pb.Id{Id: summary.ID},
			Title:       summary.Title,
			Description: summary.Description,
			FilePath:    summary.FilePath,
			AuthorId:    &pb.Id{Id: summary.AuthorId},
			Tags:        tags,
		}); err != nil {
			return err
		}
	}

	return nil
}

func NewServer(db_gorm *gorm.DB, validator *validators.Validator) pb.SummaryServiceServer {
	server := new(summaryServiceServer)
	server.dbm = &db.DBManger{DB: db_gorm}
	server.validator = validator
	return server
}
