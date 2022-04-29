package demo

import (
	"github.com/dtm-labs/dtm-cases/utils"
	"github.com/dtm-labs/dtmcli"
	"github.com/dtm-labs/dtmcli/logger"
	"github.com/gin-gonic/gin"
)

func addBaseRoute(app *gin.Engine) {

	app.POST(BusiAPI+"/deleteKey", utils.WrapHandler(func(c *gin.Context) interface{} {
		req := MustReqFrom(c)
		logger.Infof("deleting key: %s", req.Key)
		_, err := rdb.Del(rdb.Context(), req.Key).Result()
		return err
	}))
	app.GET(BusiAPI+"/queryPrepared", utils.WrapHandler(func(c *gin.Context) interface{} {
		bb, err := dtmcli.BarrierFromQuery(c.Request.URL.Query())
		if err == nil {
			err = bb.QueryPrepared(db)
		}
		return err
	}))
	app.POST(BusiAPI+"/delayDeleteKey", utils.WrapHandler(func(c *gin.Context) interface{} {
		req := MustReqFrom(c)
		return dc.DelayDelete(req.Key)
	}))

}
