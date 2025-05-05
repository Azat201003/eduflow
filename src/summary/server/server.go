package server

import (
	"context"
	"log"

	"summary-service/server/db"

	pb "github.com/Azat201003/eduflow_service_api/gen/go/summary"
	"gorm.io/gorm"
)

type summaryServiceServer struct {
	pb.UnimplementedSummaryServiceServer
	dbm *db.DBManger
}

func (s *summaryServiceServer) GetSummaryById(context context.Context, id *pb.Id) (*pb.Summary, error) {
	summary := db.Summary{ID: id.Id}
	err := s.dbm.FindSummary(&summary)
	return &pb.Summary{
		Id:          &pb.Id{Id: summary.ID},
		Title:       summary.Title,
		Description: summary.Description,
		FilePath:    summary.FilePath,
		AuthorId:    &pb.Id{Id: summary.AuthorId},
	}, err
}

func (s *summaryServiceServer) CreateSummary(ctx context.Context, summary *pb.Summary) (*pb.Id, error) {
	db_summary := db.Summary{
		Title:       summary.Title,
		Description: summary.Description,
		FilePath:    summary.FilePath,
		AuthorId:    summary.AuthorId.Id,
	}
	err := s.dbm.CreateSummary(&db_summary)
	return &pb.Id{Id: db_summary.ID}, err
}

func (s *summaryServiceServer) GetFilteredSummaries(request *pb.FilterRequest, stream pb.SummaryService_GetFilteredSummariesServer) error {
	summary := request.Filter
	conds := &db.Summary{
		Title:       summary.Title,
		Description: summary.Description,
		FilePath:    summary.FilePath,
	}
	if summary.Id != nil {
		conds.ID = summary.Id.Id
	}
	if summary.AuthorId != nil {
		conds.AuthorId = summary.AuthorId.Id
	}
	summaries, err := s.dbm.ListSummaries(conds, int((request.Page.Number-1)*request.Page.Size), int(request.Page.Size))
	log.Println(summary)
	log.Println(*summaries)
	if err != nil {
		return err
	}
	for _, summary := range *summaries {
		if err := stream.Send(&pb.Summary{
			Id:          &pb.Id{Id: summary.ID},
			Title:       summary.Title,
			Description: summary.Description,
			FilePath:    summary.FilePath,
			AuthorId:    &pb.Id{Id: summary.AuthorId},
		}); err != nil {
			return err
		}
	}

	return nil
}

func NewServer(db_gorm *gorm.DB) pb.SummaryServiceServer {
	server := new(summaryServiceServer)
	server.dbm = &db.DBManger{DB: db_gorm}
	return server
}
