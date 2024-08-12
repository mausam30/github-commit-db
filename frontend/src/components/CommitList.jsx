import React, { useState } from 'react';
import axios from 'axios';
import CommitItem from './CommitItem';

const CommitList = () => {
  const [username, setUsername] = useState('');
  const [repo, setRepo] = useState('');
  const [commits, setCommits] = useState([]);
  const [loading, setLoading] = useState(false);

  const fetchCommits = async () => {
    setLoading(true);

    try {
      const response = await axios.get('/commits', {
        params: { username, repo },
        headers: { Authorization: `token ${process.env.REACT_APP_GITHUB_TOKEN}` },
      });

      setCommits(response.data);
    } catch (error) {
      console.error(error);
    } finally {
      setLoading(false);
    }
  };

  const handleSubmit = (event) => {
    event.preventDefault();
    fetchCommits();
  };

  return (
    <div className="commit-list-container">
      <form onSubmit={handleSubmit} className="commit-list-form">
        <input
          type="text"
          placeholder="GitHub Username"
          value={username}
          onChange={(e) => setUsername(e.target.value)}
          required
        />
        <input
          type="text"
          placeholder="Repository Name (Optional)"
          value={repo}
          onChange={(e) => setRepo(e.target.value)}
        />
        <button type="submit" disabled={loading}>
          {loading ? 'Loading...' : 'Fetch Commits'}
        </button>
      </form>
      <div className="commit-list">
        {commits.map((commit) => (
          <CommitItem key={commit.sha} commit={commit} />
        ))}
      </div>
    </div>
  );
};

export default CommitList;
