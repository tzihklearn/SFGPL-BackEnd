package handler

import (
	"SFGPL-End/biz/dal"
	"SFGPL-End/biz/model"
	"SFGPL-End/biz/pojo/param"
	"SFGPL-End/biz/pojo/result"
	"github.com/gin-gonic/gin"
	"strconv"
	"strings"
	"time"
)

func GetAllProgram(c *gin.Context) {
	var pageNum int
	if c.Request.URL.Path == "/program/search" {
		pageNum = 1
	} else {
		var err error
		pageNum, err = strconv.Atoi(c.Query("pageNum"))
		if err != nil {
			c.JSON(200, &result.Result{
				Code:     result.Code_ParamInvalid,
				Msg:      "参数不正确",
				Response: nil,
			})
			return
		}
	}

	pageSize := 10

	programList := make([]*model.Program, 0)

	dal.DB.Offset((pageNum - 1) * pageSize).Limit(pageNum * pageSize).Find(&programList)
	programs := make([]*result.Program, 0)

	for i := range programList {
		categorie := model.Categorie{}

		dal.DB.Where("id = ?", programList[i].CategorieID).First(&categorie)

		programs = append(programs, &result.Program{
			Id:        programList[i].ID,
			TypeName:  categorie.Name,
			Name:      programList[i].Title,
			View:      programList[i].View,
			ActorList: programList[i].Actors,
			//Num:       int64(programList[i].ActorNum),
		})
	}

	if len(programs) == 0 {
		programs = nil
	}
	programResults := &result.ProgramResults{ProgramResults: programs}

	c.JSON(200, &result.Result{
		Code:     result.Code_Success,
		Msg:      "查询成功",
		Response: programResults,
	})

}

func Search(c *gin.Context) {
	typeStr := c.Query("type")

	name := c.Query("name")

	programList := make([]*model.Program, 0)
	programs := make([]*result.Program, 0)

	if len(typeStr) == 0 && len(name) == 0 {
		GetAllProgram(c)
		return
	} else if len(typeStr) == 0 {

		dal.DB.Where("title like ?", "%"+name+"%").Find(&programList)

	} else if len(typeStr) != 0 && len(name) == 0 {
		typeId, err := strconv.Atoi(typeStr)
		if err != nil {
			c.JSON(200, &result.Result{
				Code:     result.Code_RTErr,
				Msg:      err.Error(),
				Response: nil,
			})
			return
		}

		dal.DB.Where("categorie_id = ?", typeId).Find(&programList)

	} else {
		typeId, err := strconv.Atoi(typeStr)
		if err != nil {
			c.JSON(200, &result.Result{
				Code:     result.Code_RTErr,
				Msg:      err.Error(),
				Response: nil,
			})
			return
		}
		dal.DB.Where("categorie_id = ? AND title like ?", typeId, "%"+name+"%").Find(&programList)

	}

	for i := range programList {
		categorie := model.Categorie{}
		dal.DB.Where("id = ?", programList[i].CategorieID).First(&categorie)

		programs = append(programs, &result.Program{
			Id:        programList[i].ID,
			TypeName:  categorie.Name,
			Name:      programList[i].Title,
			View:      programList[i].View,
			ActorList: programList[i].Actors,
		})
	}
	if len(programs) == 0 {
		programs = nil
	}
	programResults := &result.ProgramResults{ProgramResults: programs}
	c.JSON(200, &result.Result{
		Code:     result.Code_Success,
		Msg:      "查询成功",
		Response: programResults,
	})

}

func Add(c *gin.Context) {
	var addProgram param.AddProgram
	err := c.ShouldBind(&addProgram)
	if err != nil {
		c.JSON(200, &result.Result{
			Code:     result.Code_ParamInvalid,
			Msg:      "参数不正确",
			Response: nil,
		})
		return
	}

	var program model.Program

	program.Title = addProgram.Name
	program.View = addProgram.Point
	program.Actors = addProgram.Actors
	program.ActorNum = int32(len(strings.Split(addProgram.Actors, "，")))
	program.CategorieID = int32(addProgram.TypeName)
	program.UpdeateTine = time.Now()
	program.CreateTime = time.Now()

	resultDB := dal.DB.Create(&program)

	if resultDB.RowsAffected != 1 {
		c.JSON(200, &result.Result{
			Code:     result.Code_RTErr,
			Msg:      "新增错误",
			Response: nil,
		})
		return
	}

	c.JSON(200, &result.Result{
		Code:     result.Code_Success,
		Msg:      "新增成功",
		Response: nil,
	})

}

func Deleted(c *gin.Context) {
	var deletedProgram param.DeletedProgram
	err := c.ShouldBind(&deletedProgram)
	if err != nil {
		c.JSON(200, &result.Result{
			Code:     result.Code_ParamInvalid,
			Msg:      "参数不正确",
			Response: nil,
		})
		return
	}

	resultDB := dal.DB.Delete(&model.Program{}, deletedProgram.Id)
	if resultDB.RowsAffected != 1 {
		c.JSON(200, &result.Result{
			Code:     result.Code_DBErr,
			Msg:      "删除错误",
			Response: nil,
		})
		return
	}

	c.JSON(200, &result.Result{
		Code:     result.Code_Success,
		Msg:      "删除成功",
		Response: nil,
	})
}

func Update(c *gin.Context) {
	var addProgram param.UpdateProgram
	err := c.ShouldBind(&addProgram)
	if err != nil {
		c.JSON(200, &result.Result{
			Code:     result.Code_ParamInvalid,
			Msg:      "参数不正确",
			Response: nil,
		})
		return
	}

	categorie := model.Categorie{}
	dal.DB.Where("name = ?", addProgram.TypeName).Find(&categorie)

	if len(categorie.Name) == 0 {
		c.JSON(200, &result.Result{
			Code:     result.Code_ParamInvalid,
			Msg:      "参数不正确",
			Response: nil,
		})
		return
	}

	var program model.Program

	program.ID = int32(addProgram.Id)
	program.Title = addProgram.Name
	program.View = addProgram.Point
	program.Actors = addProgram.Actors
	program.ActorNum = int32(len(strings.Split(addProgram.Actors, "，")))
	program.CategorieID = categorie.ID
	program.UpdeateTine = time.Now()
	program.CreateTime = time.Now()

	resultDB := dal.DB.Updates(program)

	if resultDB.RowsAffected != 1 {
		c.JSON(200, &result.Result{
			Code:     result.Code_RTErr,
			Msg:      "更新错误",
			Response: nil,
		})
		return
	}

	c.JSON(200, &result.Result{
		Code:     result.Code_Success,
		Msg:      "更新成功",
		Response: nil,
	})
}
