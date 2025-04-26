package server

import (
	"context"

	"summary-service/server/db"

	pb "github.com/Azat201003/eduflow_service_api/gen/summary"
	"gorm.io/gorm"
)

type summaryServiceServer struct {
	pb.UnimplementedSummaryServiceServer
	db *gorm.DB
}

func (s *summaryServiceServer) GetSummaryById(context context.Context, id *pb.Id) (*pb.Summary, error) {
	summary := db.Summary{ID: id.Id}
	err := db.FindSummary(s.db, &summary)
	return &pb.Summary{
		Id:          &pb.Id{Id: summary.ID},
		Title:       summary.Title,
		Description: summary.Description,
		FilePath:    summary.FilePath,
	}, err
}
func (s *summaryServiceServer) CreateSummary(ctx context.Context, summary *pb.Summary) (*pb.Empty, error) {
	db_summary := db.Summary{
		Title:       summary.Title,
		Description: summary.Description,
		FilePath:    summary.FilePath,
		ID:          summary.Id.Id,
	}
	err := db.CreateSummary(s.db, &db_summary)
	return &pb.Empty{}, err
}

func NewServer(db *gorm.DB) pb.SummaryServiceServer {
	server := new(summaryServiceServer)
	server.db = db
	return server
}
