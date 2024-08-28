-- name: GetAllCourse :many
SELECT  
    c.id,c.title, json_agg(json_build_object('occupation_id', co.id, 'occupation', co.occupation)) AS occupations
FROM course.course c
LEFT JOIN course.course_information ci ON c.course_information_id = ci.id
LEFT JOIN course.course_statistics cs ON c.course_statistics_id = cs.id
LEFT JOIN course.certificate_detail cd on c.certificate_detail_id =cd.id
LEFT JOIN course.course_occupation co on c.id = co.course_id  
group by  c.title,c.id;

