import React, { Component } from 'react';
import { Route, Routes } from 'react-router-dom';
import Admin from './Admin';
import Categories from './Categories';
import Category from './Category';
import Home from './Home';
import Movies from './Movies';
import Movie from './Movie';
export default class Content extends Component {

    render() {
        return (
            <div className='container'>
                <Routes>
                    <Route path='/' element={<Home />} />
                    <Route path='admin' element={<Admin />} />
                    <Route path='categories' element={<Categories />} />
                    <Route path='categories/:category' element={<Category />} />
                    <Route path='movies' element={<Movies />} />
                    <Route path='movies/:id' element={<Movie />} />
                </Routes>
            </div>
        )
    }
}