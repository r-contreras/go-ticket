import React from 'react'
import Navbar from './components/core/Navbar'
import Content from './components/core/Content';
import { HashRouter as Router } from 'react-router-dom';

function App() {
  return (
    <Router>
      <div>
        <Navbar />
        <Content />
      </div>
    </Router>

  );
}

export default App;
