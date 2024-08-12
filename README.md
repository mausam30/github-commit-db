# GitHub Commit Viewer

GitHub Commit Viewer is a web application that allows users to fetch and view commits from a specified GitHub repository. Users can enter their GitHub username and an optional repository name to retrieve the commits they've made.

## Table of Contents
- [Features](#features)
- [Technologies Used](#technologies-used)
- [Project Structure](#project-structure)
- [Backend Installation](#backend-installation)
- [Frontend Installation](#frontend-installation)
- [Usage](#usage)
- [Testing](#testing)
- [Future Improvements](#future-improvements)
- [Contributing](#contributing)
- [License](#license)

## Features
- **Fetch Commits**: Users can enter their GitHub username and repository name to fetch the list of commits.
- **Grid View**: Commits are displayed in a grid format, with three commits per row.
- **Responsive Design**: The application is responsive and works well on various screen sizes.
- **Hover Effects**: Commit cards have hover effects to enhance the user experience.
- **Styled Search Form**: Nicely styled and positioned search form for entering GitHub credentials.

## Technologies Used
- **Frontend**: React.js
  - React Hooks
  - Axios for API requests
  - CSS for styling
- **Backend**: Golang
  - RESTful API using `gin-gonic/gin`
  - Integration with GitHub API to fetch commits

## Project Structure
# Project Structure

- **project-root**
  - **frontend**
    - **public**
    - **src**
      - **components**
        - `CommitItem.jsx`
        - `CommitList.jsx`
        - `NavBar.jsx`
      - `App.js`
      - `index.js`
      - `app.css`
      - `index.css`
    - `package.json`
  - **backend**
    - **cmd**
      - `main.go`
    - **pkg**
      - **api**
        - `commits.go`
        - `router.go`
      - **github**
        - `client.go`
        - `commits.go`
        - `rate_limit.go`
      - **middleware**
        - `auth.go`
      - **utils**
        - `config.go`
        - `errors.go`
    - **test**
      - `api_test.go`
    - `go.mod`
    - `go.sum`
    - `.env`
  - `README.md`



## Backend Installation

1. **Clone the repository**:
    ```bash
    git clone https://github.com/mausam30/github-commit-db.git
    cd project-root/backend
    ```

2. **Install Go dependencies**:
    ```bash
    go mod tidy
    ```

3. **Set up environment variables**:
   Create a `.env` file in the `backend` directory with the following content:
    ```env
    GITHUB_TOKEN=your_github_token
    ```

4. **Run the backend server**:
    ```bash
    go run cmd/main.go
    ```

## Frontend Installation

1. **Navigate to the frontend directory**:
    ```bash
    cd project-root/frontend
    ```

2. **Install dependencies**:
    ```bash
    npm install
    ```

3. **Start the React development server**:
    ```bash
    npm start
    ```

## Usage

1. **Open the application** in your browser at `http://localhost:3000`.
2. **Enter your GitHub username** and the repository name (optional).
3. **Click "Fetch Commits"** to view the list of commits.
4. **Scroll through the grid** to see commits displayed in a grid format with hover effects.

## Testing

To run the tests for the backend:

1. **Navigate to the backend directory**:
    ```bash
    cd project-root/backend
    ```

2. **Run tests**:
    ```bash
    go test ./...
    ```

## Future Improvements
- Implement user authentication for private repositories.
- Add pagination or load more commits on scroll.
- Improve error handling and user notifications.
- Enhance UI/UX with additional styling and features.

## Contributing

If you would like to contribute to this project, please follow these steps:

1. **Fork the repository**.
2. **Create a new branch** for your changes:
    ```bash
    git checkout -b feature-branch
    ```
3. **Commit your changes**:
    ```bash
    git add .
    git commit -m "Describe your changes"
    ```
4. **Push your branch**:
    ```bash
    git push origin feature-branch
    ```
5. **Create a pull request** on GitHub.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

