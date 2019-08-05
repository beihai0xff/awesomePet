package echarts

import (
	"Wade/gorm_mysql"
	"fmt"
	"github.com/chenjiandongx/go-echarts/charts"
	"github.com/labstack/echo/v4"
	"math/rand"
	"os"
)


func BarHandler(c echo.Context) error {
	page := charts.NewPage(orderRouters("bar")...)
	page.Add(
		/*		barBase(),
				barTitle(),
				barShowLabel(),
				barXYName(),
				barColor(),
				barSplitLine(),
				barGap(),
				barYAxis(),
				barMultiYAxis(),
				barMultiXAxis(),
				barDataZoom(),
				barReverse(),
				barStack(),
				barMark(),
				barMarkCustom(),*/
	)

	page.PageTitle = "统计数据"
	f, err := os.Create(getRenderPath("total.html"))
	if err != nil {
		return err
	}
	if err = page.Render(f); err != nil {
		return err
	}
	return c.Inline("./echarts/assets/html/total.html", "temp.html")
}
