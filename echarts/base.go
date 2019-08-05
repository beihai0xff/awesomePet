package echarts

import (
	"github.com/chenjiandongx/go-echarts/charts"
	"math/rand"
	"time"
)

const (
	AssetsHost = "https://127.0.0.1:718/assets/"
	maxNum     = 50
)

type router struct {
	name string
	charts.RouterOpts
}

var (
	nameItems = []string{"衬衫", "牛仔裤", "运动裤", "袜子", "冲锋衣", "羊毛衫"}
	foodItems = []string{"面包", "牛奶", "奶茶", "棒棒糖", "加多宝", "可口可乐"}
	seed      = rand.NewSource(time.Now().UnixNano())

	routers = []router{
		{"bar", charts.RouterOpts{URL: AssetsHost + "/bar", Text: "Bar-(柱状图)"}},
		{"bar3D", charts.RouterOpts{URL: AssetsHost + "/bar3D", Text: "Bar3D-(3D 柱状图)"}},
		{"boxPlot", charts.RouterOpts{URL: AssetsHost + "/boxPlot", Text: "BoxPlot-(箱线图)"}},
		{"effectScatter", charts.RouterOpts{URL: AssetsHost + "/effectScatter", Text: "EffectScatter-(动态散点图)"}},
		{"funnel", charts.RouterOpts{URL: AssetsHost + "/funnel", Text: "Funnel-(漏斗图)"}},
		{"gauge", charts.RouterOpts{URL: AssetsHost + "/gauge", Text: "Gauge-仪表盘"}},
		{"geo", charts.RouterOpts{URL: AssetsHost + "/geo", Text: "Geo-地理坐标系"}},
		{"graph", charts.RouterOpts{URL: AssetsHost + "/graph", Text: "Graph-关系图"}},
		{"heatMap", charts.RouterOpts{URL: AssetsHost + "/heatMap", Text: "HeatMap-热力图"}},
		{"kline", charts.RouterOpts{URL: AssetsHost + "/kline", Text: "Kline-K 线图"}},
		{"line", charts.RouterOpts{URL: AssetsHost + "/line", Text: "Line-(折线图)"}},
		{"line3D", charts.RouterOpts{URL: AssetsHost + "/line3D", Text: "Line3D-(3D 折线图)"}},
		{"liquid", charts.RouterOpts{URL: AssetsHost + "/liquid", Text: "Liquid-(水球图)"}},
		{"map", charts.RouterOpts{URL: AssetsHost + "/map", Text: "Map-(地图)"}},
		{"overlap", charts.RouterOpts{URL: AssetsHost + "/overlap", Text: "Overlap-(重叠图)"}},
		{"parallel", charts.RouterOpts{URL: AssetsHost + "/parallel", Text: "Parallel-(平行坐标系)"}},
		{"pie", charts.RouterOpts{URL: AssetsHost + "/pie", Text: "Pie-(饼图)"}},
		{"radar", charts.RouterOpts{URL: AssetsHost + "/radar", Text: "Radar-(雷达图)"}},
		{"sankey", charts.RouterOpts{URL: AssetsHost + "/sankey", Text: "Sankey-(桑基图)"}},
		{"scatter", charts.RouterOpts{URL: AssetsHost + "/scatter", Text: "Scatter-(散点图)"}},
		{"scatter3D", charts.RouterOpts{URL: AssetsHost + "/scatter3D", Text: "Scatter-(3D 散点图)"}},
		{"surface3D", charts.RouterOpts{URL: AssetsHost + "/surface3D", Text: "Surface3D-(3D 曲面图)"}},
		{"themeRiver", charts.RouterOpts{URL: AssetsHost + "/themeRiver", Text: "ThemeRiver-(主题河流图)"}},
		{"wordCloud", charts.RouterOpts{URL: AssetsHost + "/wordCloud", Text: "WordCloud-(词云图)"}},
		{"page", charts.RouterOpts{URL: AssetsHost + "/page", Text: "Page-(顺序多图)"}},
	}
)
