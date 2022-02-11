import { useParams } from "react-router-dom"

export default function Movie(movie) {

    let { id } = useParams();

    return (
        <div>
            <h1> Movie {movie.title}</h1>
            <h2> Runtime: {movie.runtime}</h2>
        </div>
    )
}