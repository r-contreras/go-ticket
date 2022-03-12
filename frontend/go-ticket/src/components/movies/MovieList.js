import React, { Fragment } from 'react'
import TextLink from '../core/TextLink'

function MovieList({ movies }) {
    if (movies != undefined) {
        return (
            <Fragment>
                <p className="fs-3 fw-lighter text-break fw-italic">
                    List of movies for this genre
                </p>
                <ul className='list-group list-group-flush'>
                    {movies.map((m) =>
                        <li key={m.id} className='list-group-item bg-black'>
                            <TextLink to={`/movies/${m.id}`}> {m.title} </TextLink>
                        </li>
                    )}
                </ul>
            </Fragment>
        )
    }
    return (
        <p className="fs-3 fw-lighter text-break fw-italic">
            Sorry, no movies were found for this genre :(
        </p>
    )
}

export default MovieList