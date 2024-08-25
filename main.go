package main

import (
	"Cousre-Go/course"
	"context"
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
	Name string `json:"name"`
}

type CourseWithOccupations struct {
	CourseID       uuid.UUID    `json:"CourseID"`
	CourseTitle    string       `json:"CourseTitle"`
	CourseLanguage *string      `json:"CourseLanguage"`
	Occupations    []Occupation `json:"OccupationName"`
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

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		// Post-process the rows to group occupations by courses
		courseMap := make(map[uuid.UUID]*CourseWithOccupations)

		for _, row := range rows {
			courseID := row.CourseID
			var uuidVal uuid.UUID
			copy(uuidVal[:], courseID.Bytes[:])

			// Initialize course if not already in the map
			if _, exists := courseMap[uuidVal]; !exists {

				courseMap[uuidVal] = &CourseWithOccupations{
					CourseID:    uuid.UUID(courseID.Bytes),
					CourseTitle: row.CourseTitle.String,
					Occupations: []Occupation{},
				}
			}

			// Add occupation to the course
			if row.OccupationID.Valid {
				courseMap[uuidVal].Occupations = append(courseMap[uuidVal].Occupations, Occupation{
					Name: row.OccupationName.String,
				})
			}
		}

		// Convert the map to a slice of CourseWithOccupations
		var result []CourseWithOccupations
		for _, course := range courseMap {
			result = append(result, *course)
		}
		c.JSON(http.StatusOK, result)
	})

	r.Run()
}
