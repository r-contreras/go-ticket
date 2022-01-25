import React, { Component } from 'react'
import { Link } from 'react-router-dom'

export default class Navbar extends Component {

    render() {
        return (
            <nav>
                <div className="h-10 max-w-5xl rounded-full bg-gradient-to-r from-violet-500 to-fuchsia-500 mx-auto font-semibold">
                    <div class="flex justify-around">
                        <Link to="/"> Home </Link>
                        <Link to="/categories"> Categories </Link>
                        <Link to="/movies"> Movies </Link>
                        <Link to="/admin"> Manage Catalogue </Link>
                    </div>
                </div>
            </nav>
        )
    }
}