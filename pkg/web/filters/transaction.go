package filters

import (
	"github.com/Edward-Jackie/gotool/pkg/database/orm"
	"github.com/Edward-Jackie/gotool/pkg/global"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// Transaction 事务中间件
func Transaction() gin.HandlerFunc {
	return func(context *gin.Context) {
		db := orm.Default().Context.Session(&gorm.Session{
			Context: context.Request.Context(),
		})
		transaction := db.Begin()
		defer func() {
			transaction.Rollback()
		}()
		context.Set(global.TransactionKey, &orm.Transaction{Context: transaction})
		context.Next()
		_, exists := context.Get("error")
		if exists {
			return
		}
		transaction.Commit()
	}
}
