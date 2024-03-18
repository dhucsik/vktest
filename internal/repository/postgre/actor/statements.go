package actor

const (
	createActorStmt = `INSERT INTO actors (full_name, gender, date_of_birth)
						VALUES ($1, $2, $3) RETURNING id`

	updateActorStmt = `UPDATE actors 
						SET full_name = $2,
							gender = $3,
							date_of_birth = $4
						WHERE id = $1`

	getActorByIDStmt = `SELECT id, full_name, gender, date_of_birth FROM actors WHERE id = $1`

	deleteActorStmt = `DELETE FROM actors WHERE id = $1`

	getActorsByMovieStmt = `SELECT actors.id, actors.full_name, actors.gender, actors.date_of_birth FROM actors
						INNER JOIN movies_actors ON actors.id = movies_actors.actor_id
						WHERE movies_actors.movie_id = $1`

	getActorsByMoviesStmt = `SELECT actors.id, actors.full_name, actors.gender, actors.date_of_birth, movies_actors.movie_id 
						FROM actors
						INNER JOIN movies_actors ON actors.id = movies_actors.actor_id
						WHERE movies_actors.movie_id = ANY($1)`

	searchActorsStmt = `SELECT id, full_name, gender, date_of_birth FROM actors
						WHERE full_name ILIKE '%' || $1 || '%' LIMIT $2 OFFSET $3`

	orderActorsStmt = `SELECT id, full_name, gender, date_of_birth FROM actors
						ORDER BY full_name LIMIT $1 OFFSET $2`
)
