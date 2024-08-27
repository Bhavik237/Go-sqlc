package main

import (
	"Cousre-Go/course"
	"context"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
	_ "github.com/lib/pq"
	// Add the import statement for the package that contains the definition of `db`
)

var interruptSignals = []os.Signal{
	os.Interrupt,
	syscall.SIGTERM,
	syscall.SIGINT,
}

type Occupation struct {
	OccupationId   int    `json:"occupation_id"`
	OccupationName string `json:"occupation"`
}

type CourseWithOccupations struct {
	CourseID    uuid.UUID    `json:"CourseID"`
	CourseTitle string       `json:"CourseTitle"`
	Occupations []Occupation `json:"occupations"`
}

func main() {

	ctx, stop := signal.NotifyContext(context.Background(), interruptSignals...)
	defer stop()
	//dbConn, err := sql.Open("postgres", "postgresql://postgres:pass@localhost:5430/course?sslmode=disable")
	connPool, err := pgxpool.New(ctx, "postgresql://postgres:pass@localhost:5430/course?sslmode=disable")
	if err != nil {
		log.Fatal("Cannot connect to database:", err)
	}

	queries := course.New(connPool)

	r := gin.Default()

	r.GET("/courses", func(c *gin.Context) {
		rows, err := queries.GetAllCourse(c)
		// log.Println("rows", rows)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		//coursesMap := make(map[uuid.UUID]CourseWithOccupations)
		var coursesList []CourseWithOccupations
		for _, row := range rows {
			// Define a slice to hold the list of Occupation structs
			var occupations []Occupation

			// Unmarshal the JSON byte slice into the slice of structs
			err = json.Unmarshal(row.Occupations, &occupations)
			if err != nil {
				log.Fatalf("Error unmarshalling JSON data: %v", err)
			}
			courseID := row.ID
			var uuidVal uuid.UUID
			copy(uuidVal[:], courseID.Bytes[:])
			course := CourseWithOccupations{
				CourseID:    uuidVal,
				CourseTitle: row.Title.String,
				Occupations: occupations,
			}
			coursesList = append(coursesList, course)

		}
		result := map[string]interface{}{
			"courses": coursesList,
		}

		// var occupation []Occupation
		// //var result map[string]interface{}
		// for _, row := range rows {
		// 	fmt.Printf("%+v\n", row)
		// 	fmt.Printf("%+v\n", row.Occupations)

		// 	err1 := json.Unmarshal(row.Occupations, &occupation)
		// 	if err1 != nil {
		// 		log.Fatalf("Error unmarshalling byte data: %v", err)
		// 	}
		// }
		// Define a map to hold the JSON data

		// Post-process the rows to group occupations by courses
		//courseMap := make(map[uuid.UUID]*CourseWithOccupations)

		// for _, row := range rows {
		// 	courseID := row.ID
		// 	var uuidVal uuid.UUID
		// 	copy(uuidVal[:], courseID.Bytes[:])

		// 	if _, exists := courseMap[uuidVal]; !exists {
		// 		courseMap[uuidVal] = &CourseWithOccupations{
		// 			CourseID:    uuid.UUID(courseID.Bytes),
		// 			CourseTitle: row.Title.String,
		// 			Occupations: []Occupation{},
		// 		}
		// 	}
		// 	if row.CourseID_3.Valid {
		// 		courseMap[uuidVal].Occupations = append(courseMap[uuidVal].Occupations, Occupation{
		// 			Name: row.Occupation.String,
		// 		})
		// 	}

		// }
		// var result []CourseWithOccupations
		// for _, course := range courseMap {
		// 	result = append(result, *course)
		// }

		c.JSON(http.StatusOK, result)
	})

	r.Run()
}
