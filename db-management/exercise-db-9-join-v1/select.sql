-- TODO: answer here
SELECT 
r.id,
s.fullname,
s.class,
s.status,
r.study,
r.score
FROM reports as r
INNER JOIN students as s
ON r.student_id = s.id
WHERE r.score < 70
ORDER BY r.score ASC
