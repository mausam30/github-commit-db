import React from 'react';

const NavBar = ({ onHomeClick }) => {
  return (
    <nav className="navbar">
      <div className="navbar-brand">GitHub Commit Viewer</div>
      <button className="navbar-home-button" onClick={onHomeClick}>
        Home
      </button>
    </nav>
  );
};

export default NavBar;
