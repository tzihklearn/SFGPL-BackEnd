package handler

import (
	"SFGPL-End/biz/dal"
	"SFGPL-End/biz/model"
	"SFGPL-End/biz/pojo/param"
	"SFGPL-End/biz/pojo/result"
	"errors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"strconv"
	"time"
)

// GetAllProgram 获取所有书籍，进行分页
func GetAllProgram(c *gin.Context) {
	var pageNum int
	if c.Request.URL.Path == "/program/search" {
		pageNum = 1
	} else {
		var err error
		//获取请求参数
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

	//查询书籍列表，限制返回的数据量limit `offset` to `limit`
	dal.DB.Offset((pageNum - 1) * pageSize).Limit(pageNum * pageSize).Find(&programList)
	programs := make([]*result.Program, 0)

	for i := range programList {
		categorie := model.Categorie{}

		//查询分类，限制where id = programList[i].CategorieID
		dal.DB.Where("id = ?", programList[i].CategorieID).First(&categorie)

		actor := model.Actor{}
		//查询作者
		dal.DB.Where("id = ?", programList[i].ActorID).First(&actor)

		programs = append(programs, &result.Program{
			Id:        programList[i].ID,
			TypeName:  categorie.Name,
			Name:      programList[i].Title,
			View:      programList[i].View,
			ActorList: actor.Name,
			//Num:       int64(programList[i].ActorNum),
		})
	}

	if len(programs) == 0 {
		programs = nil
	}

	//拼接返回对象
	programResults := &result.ProgramResults{ProgramResults: programs}

	//返回
	c.JSON(200, &result.Result{
		Code:     result.Code_Success,
		Msg:      "查询成功",
		Response: programResults,
	})

}

// Search 对分类或书籍名进行查询
func Search(c *gin.Context) {
	//获取请求参数
	typeStr := c.Query("type")
	name := c.Query("name")

	programList := make([]*model.Program, 0)
	programs := make([]*result.Program, 0)

	//如果请求参数为空，则去获取书籍列表，不作筛选
	if len(typeStr) == 0 && len(name) == 0 {
		GetAllProgram(c)
		return
	} else if len(typeStr) == 0 {
		//分类参数为空
		//获取书籍列表，对title进行模糊查询
		dal.DB.Where("title like ?", "%"+name+"%").Find(&programList)

	} else if len(typeStr) != 0 && len(name) == 0 {
		//书籍名参数为空

		typeId, err := strconv.Atoi(typeStr)
		if err != nil {
			c.JSON(200, &result.Result{
				Code:     result.Code_RTErr,
				Msg:      err.Error(),
				Response: nil,
			})
			return
		}

		//获取书籍列表，where categorie_id = typeId
		dal.DB.Where("categorie_id = ?", typeId).Find(&programList)

	} else {
		//两个参数均有
		typeId, err := strconv.Atoi(typeStr)
		if err != nil {
			c.JSON(200, &result.Result{
				Code:     result.Code_RTErr,
				Msg:      err.Error(),
				Response: nil,
			})
			return
		}

		//获取书籍列表，where categorie_id = typeId and title like %name%
		dal.DB.Where("categorie_id = ? AND title like ?", typeId, "%"+name+"%").Find(&programList)

	}

	for i := range programList {
		categorie := model.Categorie{}

		//查询分类，限制where id = programList[i].CategorieID
		dal.DB.Where("id = ?", programList[i].CategorieID).First(&categorie)

		actor := model.Actor{}
		//查询作者
		dal.DB.Where("id = ?", programList[i].ActorID).First(&actor)

		//拼装返回数据
		programs = append(programs, &result.Program{
			Id:        programList[i].ID,
			TypeName:  categorie.Name,
			Name:      programList[i].Title,
			View:      programList[i].View,
			ActorList: actor.Name,
		})
	}
	if len(programs) == 0 {
		programs = nil
	}

	programResults := &result.ProgramResults{ProgramResults: programs}
	//返回
	c.JSON(200, &result.Result{
		Code:     result.Code_Success,
		Msg:      "查询成功",
		Response: programResults,
	})

}

// Add 添加书籍
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

	//事务处理
	err = dal.DB.Transaction(func(tx *gorm.DB) error {

		actor := model.Actor{}

		//查询作者id
		tx.Where("name = ?", addProgram.Actors).Find(&actor)
		if len(actor.Name) == 0 {
			actor.Name = addProgram.Actors
			actor.UpdateTime = time.Now()
			actor.CreateTime = time.Now()
			//插入作者
			resultDB := tx.Create(&actor)
			//如果插入数异常，返回
			if resultDB.RowsAffected != 1 {
				c.JSON(200, &result.Result{
					Code:     result.Code_RTErr,
					Msg:      "新增错误",
					Response: nil,
				})
				return errors.New("插入数异常")
			}
		}

		//拼装数据库实体类program
		var program model.Program

		program.Title = addProgram.Name
		program.View = addProgram.Point
		program.ActorID = actor.ID
		program.CategorieID = int32(addProgram.TypeName)
		program.UpdeateTine = time.Now()
		program.CreateTime = time.Now()

		//创建书籍，插入数据
		resultDB := tx.Create(&program)
		//如果插入数异常，返回
		if resultDB.RowsAffected != 1 {
			c.JSON(200, &result.Result{
				Code:     result.Code_RTErr,
				Msg:      "新增错误",
				Response: nil,
			})
			return errors.New("插入数异常")
		}
		return nil
	})
	if err != nil {
		c.JSON(200, &result.Result{
			Code:     result.Code_RTErr,
			Msg:      err.Error(),
			Response: nil,
		})
		return
	}

	//返回
	c.JSON(200, &result.Result{
		Code:     result.Code_Success,
		Msg:      "新增成功",
		Response: nil,
	})

}

// Deleted 删除书籍
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

	//事务处理
	err = dal.DB.Transaction(func(tx *gorm.DB) error {
		//执行软删除，where id = deletedProgram.Id,将is_deleted字段设为当前时间
		resultDB := tx.Delete(&model.Program{}, deletedProgram.Id)
		if resultDB.RowsAffected != 1 {
			c.JSON(200, &result.Result{
				Code:     result.Code_DBErr,
				Msg:      "删除错误",
				Response: nil,
			})
			return errors.New("删除数异常")
		}
		return nil
	})
	if err != nil {
		c.JSON(200, &result.Result{
			Code:     result.Code_DBErr,
			Msg:      err.Error(),
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

// Update 更新书籍
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

	//事务处理
	err = dal.DB.Transaction(func(tx *gorm.DB) error {

		actor := model.Actor{}

		//查询作者id
		tx.Where("name = ?", addProgram.Actors).Find(&actor)
		if len(actor.Name) == 0 {
			actor.Name = addProgram.Actors
			actor.UpdateTime = time.Now()
			actor.CreateTime = time.Now()
			//插入作者
			resultDB := tx.Create(&actor)
			//如果插入数异常，返回
			if resultDB.RowsAffected != 1 {
				c.JSON(200, &result.Result{
					Code:     result.Code_RTErr,
					Msg:      "新增错误",
					Response: nil,
				})
				return errors.New("插入数异常")
			}
		}

		var program model.Program

		program.ID = int32(addProgram.Id)
		program.Title = addProgram.Name
		program.View = addProgram.Point
		program.ActorID = actor.ID
		program.CategorieID = categorie.ID
		program.UpdeateTine = time.Now()
		program.CreateTime = time.Now()

		//更新书籍，以主键id作匹配
		resultDB := tx.Updates(program)

		if resultDB.RowsAffected != 1 {
			c.JSON(200, &result.Result{
				Code:     result.Code_RTErr,
				Msg:      "更新错误",
				Response: nil,
			})
			return errors.New("更新数异常")
		}
		return nil
	})
	if err != nil {
		c.JSON(200, &result.Result{
			Code:     result.Code_DBErr,
			Msg:      err.Error(),
			Response: nil,
		})
		return
	}

	//返回
	c.JSON(200, &result.Result{
		Code:     result.Code_Success,
		Msg:      "更新成功",
		Response: nil,
	})
}
