-- name: Test :one
select * from Demo
where txt ~~ '%' || @val || '%';

-- -- -- name: Test2 :one
-- select * from Demo
-- where txt like '%' || @val || '%';

-- name: Test3 :one
select * from Demo
where txt like concat('%', @val, '%');