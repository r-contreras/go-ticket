import axios from "axios";
import Header from '../components/core/Header.js';
import React, { useEffect, useState } from 'react'
import { useParams } from "react-router-dom"
import MovieList from "../components/movies/MovieList.js";

export default function Genre() {

    let { id } = useParams()
    let [genre, setGenre] = useState([])
    let [movies, setMovies] = useState([])
    useEffect(() => {
        axios.get(`http://localhost:4000/v1/genre/${id}`)
            .then((response) => {
                setGenre(response.data.genre)
                setMovies(response.data.genre.movies)
            })
    }, [])

    return (
        <div className='container'>
            <Header text={`${genre.name}`}></Header>
            <MovieList movies={movies} />
        </div>
    )
}