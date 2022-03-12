import React, { useEffect, useState } from 'react'
import Header from '../components/core/Header'
import TextLink from '../components/core/TextLink'
import axios from 'axios'

function Genres() {
    const [genres, setGenres] = useState([])

    useEffect(() => {
        axios.get('http://localhost:4000/v1/genres')
            .then((response) => {
                setGenres(response.data.genres)
            })
    }, [])

    return (
        <div>
            <Header text={"Genres"} />
            <ul className='list-group list-group-flush'>
                {genres.map((g) =>
                    <li key={g.id} className='list-group-item bg-black'>
                        <TextLink to={`${g.id}`}> {g.name} </TextLink>
                    </li>
                )}
            </ul>
        </div>

    )
}

export default Genres