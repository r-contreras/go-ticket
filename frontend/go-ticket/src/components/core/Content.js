import React, { Component } from 'react';
import { Route, Routes } from 'react-router-dom';
import Admin from '../../pages/Admin';
import Genres from '../../pages/Genres';
import Genre from '../../pages/Genre';
import Home from '../../pages/Home';
import Movies from '../../pages/Movies';
import Movie from '../../pages/Movie';
export default class Content extends Component {

    render() {
        return (
            <div className='container text-white'>
                <Routes>
                    <Route path='/' element={<Home />} />
                    <Route path='admin' element={<Admin />} />
                    <Route path='genres' element={<Genres />} />
                    <Route path='genres/:id' element={<Genre />} />
                    <Route path='movies' element={<Movies />} />
                    <Route path='movies/:id' element={<Movie />} />
                </Routes>
            </div>
        )
    }
}