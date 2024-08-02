import React, { useState } from 'react';
import axios from 'axios';
import CommitItem from './CommitItem';

const CommitList = () => {
  const [username, setUsername] = useState('');
  const [repo, setRepo] = useState('');
  const [commits, setCommits] = useState([]);
  const [loading, setLoading] = useState(false);
  const [page, setPage] = useState(1);
  const [hasMore, setHasMore] = useState(true);

  const fetchCommits = async (reset = false) => {
    if (reset) {
      setPage(1);
      setCommits([]);
    }

    setLoading(true);

    try {
      const response = await axios.get('/commits', {
        params: { username, repo, page },
        headers: { Authorization: `token ${process.env.REACT_APP_GITHUB_TOKEN}` },
      });

      const newCommits = response.data;
      setCommits(prevCommits => (reset ? newCommits : [...prevCommits, ...newCommits]));
      setHasMore(newCommits.length > 0);
      setPage(prevPage => prevPage + 1);
    } catch (error) {
      console.error(error);
      setHasMore(false);
    } finally {
      setLoading(false);
    }
  };

  const handleSubmit = (event) => {
    event.preventDefault();
    fetchCommits(true);
  };

  const handleScroll = (event) => {
    if (!loading && hasMore && event.target.scrollHeight - event.target.scrollTop === event.target.clientHeight) {
      fetchCommits();
    }
  };

  return (
    <div className="commit-list" onScroll={handleScroll} style={{ height: '80vh', overflowY: 'scroll' }}>
      <form onSubmit={handleSubmit}>
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
      {commits.map((commit) => (
        <CommitItem key={commit.sha} commit={commit} />
      ))}
      {loading && <div>Loading...</div>}
    </div>
  );
};

export default CommitList;
