import React from 'react'
import { useParams } from 'react-router-dom'
import Header from '../components/core/Header'

function Category() {
    let { category } = useParams()
    let title = capitalizeFirstLetter(category)
    return (
        <div>
            <Header text={title} />
        </div>
    )
}

function capitalizeFirstLetter(string) {
    return string.charAt(0).toUpperCase() + string.slice(1)
}

export default Category