import React from 'react'
import { useParams } from 'react-router-dom'

function Category() {
    let { category } = useParams()
    let title = capitalizeFirstLetter(category)
    return (
        <div>
            <h2> {title} </h2>
        </div>
    )
}

function capitalizeFirstLetter(string) {
    return string.charAt(0).toUpperCase() + string.slice(1)
}

export default Category