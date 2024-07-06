package controllerx

import (
	"github.com/shanluzhineng/app"
	"github.com/shanluzhineng/entity"
	"github.com/shanluzhineng/mongodbr"
)

func GetEntityService[T mongodbr.IEntity]() entity.IEntityService[T] {
	return app.Context.GetInstance(new(entity.IEntityService[T])).(entity.IEntityService[T])
}
