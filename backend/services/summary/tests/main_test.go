package tests

import (
	"context"
	"fmt"
	"io"
	"math/rand/v2"
	"strconv"
	"summary-service/server/db"
	"testing"

	"github.com/Azat201003/eduflow/backend/libs/config"
	"github.com/Azat201003/eduflow/backend/libs/gen/go/summary"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type ClientTestSuite struct {
	suite.Suite
	Client *summary.SummaryServiceClient
	dbm    *db.DBManger
}

func TestClientSuite(t *testing.T) {
	t.Helper()
	t.Parallel()

	conf, err := config.GetConfig("../../config.yaml")
	assert.NoError(t, err)
	summary_conf, err := conf.GetServiceById(1)
	assert.NoError(t, err)
	summary_conn, err := grpc.NewClient(fmt.Sprintf("%v:%v", summary_conf.Host, summary_conf.Port), grpc.WithTransportCredentials(insecure.NewCredentials()))
	assert.NoError(t, err)
	summary_client := summary.NewSummaryServiceClient(summary_conn)

	db_conf := conf.Database
	conn_conf := summary_conf.Connect
	dsn := fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=%v sslmode=disable TimeZone=Europe/Moscow search_path=%v", db_conf.Host, conn_conf.User, conn_conf.Password, conn_conf.DB, db_conf.Port, conn_conf.Schema)
	db_conn, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	assert.NoError(t, err)

	s := ClientTestSuite{
		Client: &summary_client,
		dbm:    &db.DBManger{DB: db_conn},
	}
	suite.Run(t, &s)
}

func (s *ClientTestSuite) compareEqualSummaries(s1, s2 *summary.Summary) {
	s.Equal(s1.Title, s2.Title, "titles aren't same")
	s.Equal(s1.Description, s2.Description, "descriptions aren't same")
	s.Equal(s1.FilePath, s2.FilePath, "file paths aren't same")
	s.Equal(s1.AuthorId.Id, s2.AuthorId.Id, "authors aren't same")
	s.Equal(s1.Id.Id, s2.Id.Id, "IDs aren't same")
}

func (s *ClientTestSuite) TestCreatingGetting() {
	i := strconv.Itoa(rand.Int())
	obj := &summary.Summary{Title: "Test", Description: "testing file", FilePath: "test-" + i, AuthorId: &summary.Id{Id: 2}}
	id, err := (*s.Client).CreateSummary(context.Background(), obj)
	obj.Id = id
	s.NoError(err)
	resp, err := (*s.Client).GetSummaryById(context.Background(), id)
	s.NoError(err)
	s.compareEqualSummaries(obj, resp)
}

func (s *ClientTestSuite) TestGetting() {
	resp, err := (*s.Client).GetSummaryById(context.Background(), &summary.Id{Id: 1})
	s.NoError(err)
	obj := &db.Summary{ID: 1}
	err = s.dbm.FindSummary(obj)
	s.NoError(err)
	s.compareEqualSummaries(&summary.Summary{Title: obj.Title, FilePath: obj.FilePath, Description: obj.Description, Id: &summary.Id{Id: obj.ID}, AuthorId: &summary.Id{Id: obj.AuthorId}}, resp)
}

func (s *ClientTestSuite) TestCreating() {
	obj := &summary.Summary{
		Title:       "Test",
		Description: "Summary service testing creating",
		FilePath:    "Test-" + strconv.Itoa(rand.Int()),
		AuthorId:    &summary.Id{Id: 2},
	}
	resp, err := (*s.Client).CreateSummary(context.Background(), obj)
	s.NoError(err)
	obj.Id = &summary.Id{Id: resp.Id}
	obj2 := &db.Summary{ID: resp.Id}
	err = s.dbm.FindSummary(obj2)
	s.NoError(err)
	s.compareEqualSummaries(obj, &summary.Summary{
		Title:       obj2.Title,
		Description: obj2.Description,
		AuthorId:    &summary.Id{Id: obj2.AuthorId},
		FilePath:    obj2.FilePath,
		Id:          &summary.Id{Id: resp.Id},
	})
}

func (s *ClientTestSuite) TestList() {
	stream, err := (*s.Client).GetFilteredSummaries(context.Background(), &summary.FilterRequest{
		Filter: &summary.Summary{
			AuthorId: &summary.Id{Id: 2},
			Tags: []*summary.Tag{
				{Id: &summary.Id{Id: 7}},
			},
		},
		Page: &summary.Page{
			Number: 1,
			Size:   5,
		},
	})
	s.NoError(err)

	obj, err := stream.Recv()
	s.NoError(err)

	s.compareEqualSummaries(obj, &summary.Summary{
		Id:          &summary.Id{Id: 4},
		Title:       "Test",
		Description: "testing file",
		FilePath:    "test",
		AuthorId:    &summary.Id{Id: 2},
	})

	for {
		obj, err = stream.Recv()
		if err == io.EOF {
			break
		}
		s.Equal(obj.AuthorId.Id, uint64(2))
		fmt.Println(obj.Title)
		for _, tag := range obj.Tags {
			fmt.Println("--> ", tag.Name, "\t\t", tag.Color)
		}
	}
	err = stream.CloseSend()
	s.NoError(err)
}

func (s *ClientTestSuite) TestFilteredGettingNullPageSize() {
	stream, err := (*s.Client).GetFilteredSummaries(context.Background(), &summary.FilterRequest{
		Filter: &summary.Summary{
			Title:       "New",
			Description: "New file",
		},
		Page: &summary.Page{},
	})

	_, err = stream.Recv()

	s.ErrorContains(err, "Page size is 0")
}
