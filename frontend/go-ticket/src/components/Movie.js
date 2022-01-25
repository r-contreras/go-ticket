import { useParams } from "react-router-dom"

export default function Movie(movie) {

    let { id } = useParams();

    return (
        <div>
            <h1> Movie {id}</h1>
            <h2> {movie.runtime}</h2>
        </div>
    )
}