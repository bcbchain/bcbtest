package webapi

import (
	"bcbtest/parse"
	"bcbtest/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

//Get get Preview
// @Tags Preview(统计预览)
// @Summary 查询总区块，总交易，总账户数
// @Description get preview
// @Accept  json
// @Produce json
// @Success 200
// @Router /api/v1/start [get]
func startHandler(c *gin.Context) {
	service.ExecBCBCase()
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": "success.",
	})
}

func viewCasesHandler(c *gin.Context) {
	//testKinds := service.GetTestKinds()
	testKinds := service.GetCases()
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": testKinds,
	})
}

func viewReportsHandler(c *gin.Context) {
	mode := c.Query("mode")
	if mode == "1" {
		detailReportTestKind := service.GetReportsDetail()
		c.JSON(http.StatusOK, gin.H{
			"code": 200,
			"data": detailReportTestKind,
		})
		return
	}
	reportTestKind := service.GetReports()
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": reportTestKind,
	})
}

func viewExcelHandler(c *gin.Context) {
	detailReportTestKind := parse.GetTestKinds()
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": detailReportTestKind,
	})
}

func viewReportHandler(c *gin.Context) {
	caseIDParam := c.Param("case_id")
	caseID, err := strconv.ParseInt(caseIDParam, 10, 64)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 53001,
			"data": "param error",
		})
	}

	reportTestKind := service.GetReport(caseID)
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": reportTestKind,
	})
}
