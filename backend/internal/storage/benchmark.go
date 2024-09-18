package storage

import (
	"context"

	"github.com/Infinities-ICT-Solutions/project-dashboard/data"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type BenchmarkStorage interface {
	GetById(hex string) (*data.Benchmark, error)
	GetAll() ([]*data.Benchmark, error)
	Create(benchmark *data.Benchmark) (*data.Benchmark, error)
	Update(hex string, benchmark *data.Benchmark) (*data.Benchmark, error)
	Delete(hex string) error
}

type mongoBenchmarkStorage struct {
	db *mongo.Database
}

func NewBenchmarkStorage(db *mongo.Database) *mongoBenchmarkStorage {
	return &mongoBenchmarkStorage{db: db}
}

func (p *mongoBenchmarkStorage) GetById(hex string) (*data.Benchmark, error) {
	id, err := primitive.ObjectIDFromHex(hex)
	if err != nil {
		return nil, err
	}

	benchmark := &data.Benchmark{}
	err = p.db.Collection("benchmarks").FindOne(context.TODO(), bson.M{"_id": id}).Decode(benchmark)
	if err != nil {
		return nil, err
	}

	return benchmark, nil
}

func (p *mongoBenchmarkStorage) GetAll() ([]*data.Benchmark, error) {
	cursor, err := p.db.Collection("benchmarks").Find(context.TODO(), bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.TODO())

	benchmarks := []*data.Benchmark{}
	for cursor.Next(context.TODO()) {
		benchmark := &data.Benchmark{}
		err := cursor.Decode(benchmark)
		if err != nil {
			return nil, err
		}
		benchmarks = append(benchmarks, benchmark)
	}

	return benchmarks, nil
}

func (p *mongoBenchmarkStorage) Create(benchmark *data.Benchmark) (*data.Benchmark, error) {
	res, err := p.db.Collection("benchmarks").InsertOne(context.TODO(), benchmark)
	if err != nil {
		return nil, err
	}

	benchmark.ID = res.InsertedID.(primitive.ObjectID)
	return benchmark, nil
}

func (p *mongoBenchmarkStorage) Update(hex string, benchmark *data.Benchmark) (*data.Benchmark, error) {
	id, err := primitive.ObjectIDFromHex(hex)
	if err != nil {
		return nil, err
	}

	_, err = p.db.Collection("benchmarks").UpdateOne(context.TODO(), bson.M{"_id": id}, bson.M{"$set": benchmark})
	if err != nil {
		return nil, err
	}

	return benchmark, nil
}

func (p *mongoBenchmarkStorage) Delete(hex string) error {
	id, err := primitive.ObjectIDFromHex(hex)
	if err != nil {
		return err
	}

	_, err = p.db.Collection("benchmarks").DeleteOne(context.TODO(), bson.M{"_id": id})
	if err != nil {
		return err
	}

	return nil
}
