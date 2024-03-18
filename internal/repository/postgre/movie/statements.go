package movie

const (
	createMovieStmt = `INSERT INTO movies (title, description, release_date, rating)
		VALUES ($1, $2, $3, $4) RETURNING id
	`

	createMovieActorStmt = `INSERT INTO movies_actors (movie_id, actor_id) 
		VALUES ($1, $2)`

	updateMovieStmt = `UPDATE movies 
						SET title = $2,
							description = $3,
							release_date = $4,
							rating = $5
							WHERE id = $1`

	deleteMovieActorsStmt = `DELETE FROM movies_actors WHERE movie_id = $1`

	getByIDStmt = `SELECT id, title, description, release_date, rating FROM movies WHERE id = $1`

	deleteMovieStmt = `DELETE FROM movies WHERE id = $1`

	searchMoviesStmt = `SELECT id, title, description, release_date, rating FROM movies
		WHERE title ILIKE '%' || $1 || '%' LIMIT $2 OFFSET $3`

	getMoviesByActorIDsStmt = `SELECT movies.id, movies.title, movies.description, movies.release_date, movies.rating
									FROM movies	
									INNER JOIN movies_actors ON movies.id = movies_actors.movie_id
									WHERE movies_actors.actor_id = ANY($1)`

	getMoviesByActorsStmt = `SELECT movies.id, movies.title, movies.description, movies.release_date, movies.rating, movies_actors.actor_id	
								FROM movies
								INNER JOIN movies_actors ON movies.id = movies_actors.movie_id
								WHERE movies_actors.actor_id = ANY($1)`

	orderMoviesByRatingStmt = `SELECT id, title, description, release_date, rating FROM movies
		ORDER BY rating DESC LIMIT $1 OFFSET $2`

	orderMoviesByReleaseDateStmt = `SELECT id, title, description, release_date, rating FROM movies	
		ORDER BY release_date DESC LIMIT $1 OFFSET $2`

	orderMoviesByTitleStmt = `SELECT id, title, description, release_date, rating FROM movies
		ORDER BY title LIMIT $1 OFFSET $2`

	getMoviesByActorID = `SELECT movies.id, movies.title, movies.description, movies.release_date, movies.rating
		FROM movies
		INNER JOIN movies_actors ON movies.id = movies_actors.movie_id
		WHERE movies_actors.actor_id = $1`
)
