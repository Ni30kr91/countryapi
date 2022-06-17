package main

import (
	"encoding/csv"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	setupRoutes(r)
	r.Run(":8082") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")

}

func setupRoutes(r *gin.Engine) {
	r.GET("/cases/new/country/:country", Dummynew)
	r.GET("/cases/total/country/:from_date", Dummytotal)

}

//Dummy new
func Dummynew(c *gin.Context) {
	date, ok := c.GetQuery("date")
	country, ok := c.Params.Get("country")
	records := readCsvFile("./full_data.csv")
	Name := getNew(records, date, country)
	if ok == false {
		res := gin.H{
			"error": "country name is missing",
		}
		c.JSON(http.StatusOK, res)
		return
	}

	res := gin.H{
		"date": date,
		"Case": Name,
	}
	c.JSON(http.StatusOK, res)
}

func getNew(records [][]string, input string, country string) []string {

	var con = []string{}
	for i := 0; i < len(records); i++ {

		if records[i][1] == country {
			if records[i][0] == input {

				con = append(con, records[i][2])
			}
		}

	}
	return con
}
func Dummytotal(c *gin.Context) {

	date, ok := c.Params.Get("from_date")
	records := readCsvFile("./full_data.csv")
	total_cases := getTotal(records, date)
	if ok == false {
		res := gin.H{
			"error": "date is missing",
			"date":  date,
		}
		c.JSON(http.StatusOK, res)
		return
	}
	/*
		city := ""
	*/
	res := gin.H{
		"total_cases": total_cases,
		"date":        date,
		"count":       len(total_cases),
	}
	c.JSON(http.StatusOK, res)
}

//
//Dummy total...
/*func Dummytotal(c *gin.Context) {
	date, ok := c.Params.Get("from_date")
	records := readCsvFile("./full_data.csv")
	//country, ok := c.Params.Get("country")
	total_cases := gettotal(records, date)
	if ok == false {
		res := gin.H{
			"error": "date  is missing",
			"date":  date,
			"test":  "test",
		}
		c.JSON(http.StatusOK, res)
		return
	}

	res := gin.H{
		"total_cases": total_cases,
		"date":        date,
		"country":     len(total_cases),
	}
	c.JSON(http.StatusOK, res)
}

func gettotal(records [][]string, date string) []string {

	var country = []string{}
	for i := 1; i < len(records); i++ {

		if records[i][0] == date {

			country = append(country, records[i][2])

		}

	}
	return country
}
*/

func getTotal(records [][]string, date string) []int64 {

	var total_cases = []int64{}

	var sum int64
	for i := 1; i < len(records); i++ {

		//fmt.Println(records[0][0], i)
		if records[i][0] >= date {

			//total_cases = append(total_cases, records[i][4])

			temp, _ := strconv.ParseInt(records[i][4], 0, 8)
			sum = temp + sum

		}

	}
	total_cases = append(total_cases, sum)

	return total_cases
}
func readCsvFile(filePath string) [][]string {
	f, err := os.Open(filePath)
	if err != nil {
		log.Fatal("Unable to read input file "+filePath, err)
	}
	defer f.Close()

	csvReader := csv.NewReader(f)
	records, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal("Unable to parse file as CSV for "+filePath, err)
	}

	return records
}
