import Navbar from './components/Navbar'
import Header from './components/Header'
import Content from './components/Content';
import {HashRouter as Router} from 'react-router-dom';

function App() {
  return (
    <Router>
      <div className='min-h-screen bg-black text-white font-Poppins'>
          <Navbar />
          <Header text='Go Ticket!'/>
          <Content />
      </div>
    </Router>
    
  );
}

export default App;
