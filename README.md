# fast-track-quiz-backend

## Things that could've been implemented better

- Proper error handling - try/catch blocks
- Algorithm for determining how a user did compared to others - a proper leaderboard implementation would've been faster and would allow for more options
- Error handling when user tries to submit a session with an id that is already in "Database"
- Implement proper environment configuration - currently the backend is hardcoded to be on localhost:8081

## Dependencies
- go get -u github.com/gorilla/mux
- go get -u github.com/gorilla/handlers