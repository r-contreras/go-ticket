import React from 'react'
import { Link } from 'react-router-dom'

function Categories() {
    return (
        <div>
            <h2>Categories</h2>
            <ul>
                <li>
                    <Link to='comedy'> Comedy </Link>
                </li>
                <li>
                    <Link to='drama'> Drama </Link>
                </li>
            </ul>
        </div>

    )
}

export default Categories