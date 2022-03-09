import React from 'react'
import { Link } from 'react-router-dom'
import Header from '../components/core/Header'
import TextLink from '../components/core/TextLink'

function Categories() {
    return (
        <div>
            <Header text={"Categories"} />
            <ul className='list-group list-group-flush'>
                <li className='list-group-item bg-black'>
                    <TextLink to='comedy'> Comedy </TextLink>
                </li>
                <li className='list-group-item bg-black'>
                    <TextLink to='drama'> Drama </TextLink>
                </li>
            </ul>
        </div>

    )
}

export default Categories