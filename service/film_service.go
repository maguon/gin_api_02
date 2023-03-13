package service

import (
	"context"
	"fmt"
	"gin_api_02/global"
	res "gin_api_02/model/response"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/zap"
)

type FilmService struct{}

func (filmService *FilmService) GetFilmInfo(queryModel res.FilmQuery) (list interface{}, total int64, err error) {
	limit := queryModel.PageSize
	offset := queryModel.PageSize * (queryModel.PageNumber - 1)
	db := global.SYS_DB.Table("film_info").Model(&res.FilmInfo{})
	var filmInfoList []res.FilmInfo
	fmt.Println(queryModel)
	if queryModel.ID != 0 {
		db = db.Where("id = ?", queryModel.ID)
	}
	if queryModel.UniqueId != "" {
		db = db.Where("unique_id = ?", queryModel.UniqueId)
	}
	if queryModel.Studio != "" {
		db = db.Where("studio = ?", queryModel.Studio)
	}
	if queryModel.Status != 0 {
		db = db.Where("status = ?", queryModel.Status)
	}
	if !queryModel.PublishStart.IsZero() && !queryModel.PublishEnd.IsZero() {
		db = db.Where("publish_date >=?", queryModel.PublishStart)
		db = db.Where("publish_date <=?", queryModel.PublishEnd)
	}
	if queryModel.ActorId != "" {
		db = db.Where(" ? =ANY(casts)", queryModel.ActorId)
	}
	if queryModel.CategoryId != "" {
		db = db.Where(" ? =ANY(category)", queryModel.CategoryId)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Debug().Limit(limit).Offset(offset).Order("id desc").Find(&filmInfoList).Error
	return filmInfoList, total, err
}

func (filmService *FilmService) GetMfilmInfo(queryModel res.FilmQuery) (list interface{}, total int64, err error) {
	limit := queryModel.PageSize
	skip := queryModel.PageSize * (queryModel.PageNumber - 1)
	collection := global.SYS_MONGO.Collection("mav_film")
	opts := options.Find()
	var filmList []res.Mfilm
	filter := bson.M{}
	fmt.Println(primitive.NewDateTimeFromTime(queryModel.PublishStart))
	fmt.Println(primitive.NewDateTimeFromTime(queryModel.PublishEnd))

	if queryModel.UniqueId != "" {
		filter["unique_id"] = queryModel.UniqueId
	}
	if queryModel.ActorId != "" {
		fmt.Println(queryModel.ActorId)
		actorId, _ := primitive.ObjectIDFromHex(queryModel.ActorId)
		filter["actor._id"] = actorId
	}
	if queryModel.ActressId != "" {
		actressId, _ := primitive.ObjectIDFromHex(queryModel.ActressId)
		filter["actress._id"] = actressId
	}
	if queryModel.DirectorId != "" {
		directorId, _ := primitive.ObjectIDFromHex(queryModel.DirectorId)
		filter["director._id"] = directorId
	}
	if queryModel.StudioId != "" {
		studioId, _ := primitive.ObjectIDFromHex(queryModel.StudioId)
		filter["studio._id"] = studioId
	}
	if queryModel.CategoryId != "" {
		fmt.Println(queryModel.CategoryId)
		categoryId, _ := primitive.ObjectIDFromHex(queryModel.CategoryId)
		filter["category._id"] = categoryId
	}
	if !queryModel.PublishStart.IsZero() && !queryModel.PublishEnd.IsZero() {
		filter["publish_date"] = bson.M{"$gte": primitive.NewDateTimeFromTime(queryModel.PublishStart), "$lte": primitive.NewDateTimeFromTime(queryModel.PublishEnd)}
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	cur, err := collection.Find(ctx, filter, opts.SetSort(bson.M{"publish_date": -1}).SetSkip(int64(skip)).SetLimit(int64(limit)))
	if err != nil {
		global.SYS_LOG.Error("Mongo cursor error : ", zap.Error(err))
	}
	if err = cur.All(ctx, &filmList); err != nil {
		global.SYS_LOG.Error("Cursor convert error : ", zap.Error(err))
	}
	count, err := collection.CountDocuments(ctx, filter)
	for cur.Next(ctx) {
		var filmItem res.Mfilm
		cur.Decode(&filmItem)
		filmList = append(filmList, filmItem)
	}
	return filmList, count, err
}
