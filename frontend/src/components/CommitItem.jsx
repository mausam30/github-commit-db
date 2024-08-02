import React from 'react';

const CommitItem = ({ commit }) => {
  return (
    <div className="commit-item">
      <p><strong>Message:</strong> {commit.commit.message}</p>
      <p><strong>Author:</strong> {commit.commit.author.name}</p>
      <p><strong>Date:</strong> {new Date(commit.commit.author.date).toLocaleString()}</p>
      <p><strong>SHA:</strong> {commit.sha.substring(0, 7)}</p>
      {commit.repository && <p><strong>Repository:</strong> {commit.repository.name}</p>}
    </div>
  );
};

export default CommitItem;
