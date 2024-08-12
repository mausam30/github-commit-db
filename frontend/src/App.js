
import React, { useRef } from 'react';
import NavBar from './components/Navbar';
import CommitList from './components/CommitList';
import './App.css';

const App = () => {
  const commitListRef = useRef();

  const handleHomeClick = () => {
    if (commitListRef.current) {
      commitListRef.current.resetCommits();
    }
  };

  return (
    <div className="App">
      <NavBar onHomeClick={handleHomeClick} />
      <h1>Get Your Commits</h1>
      <CommitList ref={commitListRef} />
    </div>
  );
};

export default App;
