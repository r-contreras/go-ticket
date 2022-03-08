import React, { useEffect, useState } from 'react'
import axios from "axios";
import Header from '../components/core/Header.js';
import TextLink from '../components/core/TextLink.js';

function Movies() {
    const [movies, setMovies] = useState([])

    useEffect(() => {
        axios.get("http://localhost:4000/v1/movies")
            .then((response) => {
                setMovies(response.data.movies)
            })
    }, [])

    return (
        <div>
            <Header text={"Movies"} />
            <ul className='list-group list-group-flush'>
                <nav>
                    {movies.map((m) =>
                        <li key={m.id} className='list-group-item bg-black'>
                            <TextLink to={`/movies/${m.id}`}>
                                {m.title} ({m.year})
                            </TextLink>
                        </li>
                    )}
                </nav>
            </ul>
        </div>
    )
}

export default Movies