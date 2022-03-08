import axios from "axios";
import Header from '../components/core/Header.js';
import React, { useEffect, useState } from 'react'
import { useParams } from "react-router-dom"

export default function Movie() {

    let { id } = useParams();

    let [movie, setMovie] = useState({})

    useEffect(() => {
        axios.get(`http://localhost:4000/v1/movie/${id}`)
            .then((response) => {
                setMovie(response.data.movie)
            })
    }, [])

    return (
        <div className='container'>
            <Header text={`${movie.title} (${movie.year})`}></Header>
            <p className="fs-3 fw-lighter text-break fw-italic">
                {movie.description}
            </p>
            <p className="text-muted">
                Runtime: {movie.runtime} min.
            </p>

        </div>
    )
}