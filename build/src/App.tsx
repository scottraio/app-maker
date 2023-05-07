import React from 'react';
import GameForm from './components/GameForm';
import GameList from './components/GameList';
import './App.css';

function App() {
  return (
    <div className="App">
      <h1>Greatest Video Game Tracker</h1>
      <GameForm />
      <GameList />
    </div>
  );
}

export default App;