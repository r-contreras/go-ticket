import React, { Component } from 'react'
import { Link } from 'react-router-dom'

export default class Navbar extends Component {

    render() {
        return (
            <nav className="navbar navbar-expand-sm navbar-dark bg-dark ">
                <div className="container-fluid d-flex justify-content-between">
                    <a className="navbar-brand display-7 fw-bolder " href="#/">
                        <img alt="Go Ticket! logo" src="logo.png" height="60" className="d-inline-block" />
                        Go Ticket!
                    </a>
                    <div className='navbar-nav'>
                        <Link className='nav-link' to="/"> Home </Link>
                        <Link className='nav-link' to="/categories"> Categories </Link>
                        <Link className='nav-link' to="/movies"> Movies </Link>
                        <Link className='nav-link disabled' to="/admin"> Manage Catalogue </Link>
                    </div>
                    <div id="searchbar">
                        <form className="d-flex input-group w-auto">
                            <input
                                type="search"
                                className="form-control rounded bg-black text-white border-0"
                                placeholder="Search"
                                aria-label="Search"
                                aria-describedby="search-addon"
                            />
                            <span className="input-group-text border-0 bg-dark text-white" id="search-addon">
                                <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" fill="currentColor" className="bi bi-search" viewBox="0 0 16 16">
                                    <path d="M11.742 10.344a6.5 6.5 0 1 0-1.397 1.398h-.001c.03.04.062.078.098.115l3.85 3.85a1 1 0 0 0 1.415-1.414l-3.85-3.85a1.007 1.007 0 0 0-.115-.1zM12 6.5a5.5 5.5 0 1 1-11 0 5.5 5.5 0 0 1 11 0z" />
                                </svg>
                            </span>
                        </form>
                    </div>
                </div>
            </nav>
        )
    }
}