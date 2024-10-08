// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: courseSeach.sql

package course

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const getAllCourse = `-- name: GetAllCourse :many
SELECT  
    c.id,c.title, json_agg(json_build_object('occupation_id', co.id, 'occupation', co.occupation)) AS occupations
FROM course.course c
LEFT JOIN course.course_information ci ON c.course_information_id = ci.id
LEFT JOIN course.course_statistics cs ON c.course_statistics_id = cs.id
LEFT JOIN course.certificate_detail cd on c.certificate_detail_id =cd.id
LEFT JOIN course.course_occupation co on c.id = co.course_id  
group by  c.title,c.id
`

type GetAllCourseRow struct {
	ID          pgtype.UUID
	Title       pgtype.Text
	Occupations []byte
}

func (q *Queries) GetAllCourse(ctx context.Context) ([]GetAllCourseRow, error) {
	rows, err := q.db.Query(ctx, getAllCourse)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetAllCourseRow
	for rows.Next() {
		var i GetAllCourseRow
		if err := rows.Scan(&i.ID, &i.Title, &i.Occupations); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
