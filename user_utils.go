package controllerx

import (
	"fmt"

	"github.com/shanluzhineng/abmp/pkg/log"
	"github.com/shanluzhineng/abmp/pkg/utils/reflector"
	"github.com/shanluzhineng/entity"
	"github.com/kataras/iris/v12"
)

func GetUserId(ctx iris.Context) string {
	claims := GetCasdoorMiddleware().GetUserClaims(ctx)
	if claims != nil {
		return claims.Id
	}
	return ""
}

func checkEntityIsIEntityWithUser(entityValue interface{}) entity.IEntityWithUser {
	v, ok := entityValue.(entity.IEntityWithUser)
	if !ok {
		return nil
	}
	return v
}

func AddUserIdFilterIfNeed(filter map[string]interface{}, entity interface{}, ctx iris.Context) {
	if filter == nil {
		return
	}
	if checkEntityIsIEntityWithUser(entity) == nil {
		return
	}
	currentUserId := GetUserId(ctx)
	if currentUserId == "" {
		return
	}
	filter["creatorId"] = currentUserId
}

func FilterMustIsCurrentUserId(entity interface{}, ctx iris.Context) bool {
	entityWithUser := checkEntityIsIEntityWithUser(entity)
	if entityWithUser == nil {
		return true
	}
	userId := GetUserId(ctx)
	if userId == "" {
		return false
	}
	ok := (userId == entityWithUser.GetCreatorId())
	if !ok {
		log.Warn(fmt.Sprintf("用户修改的对象不属于当前登录的用户,对象类型:%s,userId:%s",
			reflector.GetFullName(entity),
			userId))
	}
	return ok
}
