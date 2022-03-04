import React from 'react'
import { Link } from 'react-router-dom'
import Header from '../components/core/Header'

function Categories() {
    return (
        <div>
            <Header text={"Categories"} />
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