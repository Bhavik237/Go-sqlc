-- name: GetAllCourse :many
SELECT  c.id AS course_id,
    c.title AS course_title,
    c.language AS course_language,
    co.id AS occupation_id,
    co.occupation AS occupation_name
FROM course.course c
LEFT JOIN course.course_information ci ON c.course_information_id = ci.id
LEFT JOIN course.course_statistics cs ON c.course_statistics_id = cs.id
LEFT JOIN course.certificate_detail cd on c.certificate_detail_id =cd.id
LEFT JOIN course.course_occupation co on c.id = co.course_id  
Group by 
c.id,
co.id,
c.code,
co.course_id ,
c.title,
c.short_description,
c.long_description,
c.language,
c.learning_outcome,
c.start_date,
c.end_date,
c.price,
c.availability,
c.assessment,
c.url,
c.source_system_name,
c.source_system_id,
c.learning_partner_type,
c.course_mode,
c.course_status_id,
c.course_type,
ci.id ;
