package redis_manager

import (
	"context"
	"encoding/json"
	"fmt"
	"math/rand/v2"
	"time"

	pb "github.com/Azat201003/eduflow_service_api/gen/go/filager"
	"github.com/redis/go-redis/v9"
)

type RedisManager struct {
	Client *redis.Client
}

func NewRedisManager(redisClient *redis.Client) *RedisManager {
	return &RedisManager{Client: redisClient}
}

func (rm *RedisManager) CreateSendingSession(session *pb.StartWriteRequest) (uuid uint64, err error) {
	uuid = rand.NewPCG(uint64(time.Now().Nanosecond()), uint64(time.Now().Second())).Uint64()
	val, err := json.Marshal(session)
	if err != nil {
		return 0, err
	}
	return uuid, rm.Client.Set(context.Background(), fmt.Sprintf("sends[%v]", uuid), val, time.Second*time.Duration(session.FileSize)).Err()
}

func (rm *RedisManager) GetSendingSession(uuid uint64) (session *pb.StartWriteRequest, err error) {
	r := rm.Client.Get(context.Background(), fmt.Sprintf("sends[%v]", uuid))
	if r.Err() != nil {
		return nil, r.Err()
	}

	session = new(pb.StartWriteRequest)
	bytes, err := r.Bytes()
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(bytes, session)
	return session, err
}

func (rm *RedisManager) CloseSendingSession(uuid uint64) (err error) {
	return rm.Client.Del(context.Background(), fmt.Sprintf("sends[%v]", uuid)).Err()
}

type ReadingSession struct {
	FilePath  string      `json:"file_path"`
	ChunkSize uint16      `json:"chunk_size"`
	FileSize  uint64      `json:"file_size"`
	FileType  pb.FileType `json:"file_type"`
}

func (rm *RedisManager) CreateReadingSession(session *ReadingSession) (uuid uint64, err error) {
	uuid = rand.NewPCG(uint64(time.Now().Nanosecond()), uint64(time.Now().Second())).Uint64()
	val, err := json.Marshal(session)
	if err != nil {
		return 0, err
	}
	err = rm.Client.Set(context.Background(), fmt.Sprintf("reads[%v]", uuid), val, time.Second*time.Duration(session.FileSize)).Err()
	return uuid, err
}

func (rm *RedisManager) GetReadingSession(uuid uint64) (session *ReadingSession, err error) {
	r := rm.Client.Get(context.Background(), fmt.Sprintf("reads[%v]", uuid))
	if r.Err() != nil {
		return nil, r.Err()
	}

	session = new(ReadingSession)
	bytes, err := r.Bytes()
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(bytes, session)
	return session, err
}
