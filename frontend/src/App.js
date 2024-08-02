import React from 'react';
import CommitList from './components/CommitList';
import './App.css';

const App = () => {
  return (
    <div className="App">
      <h1>GitHub Commit Viewer</h1>
      <CommitList />
    </div>
  );
};

export default App;
