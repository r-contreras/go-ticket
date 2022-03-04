import React, { useEffect, useState } from 'react'
import axios from "axios";
import Header from '../components/core/Header.js';
import { Link } from 'react-router-dom';

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
            <ul>
                <nav>
                    {movies.map((m) =>
                        <li key={m.id}>
                            <Link to={`/movies/${m.title}`}> {m.title} </Link>
                        </li>
                    )}
                </nav>
            </ul>
        </div>
    )
}

export default Movies