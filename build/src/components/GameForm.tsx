import React, { useState } from 'react';
import { db } from '../firebase';

const GameForm = () => {
  const [game, setGame] = useState('')
  const [winner, setWinner] = useState('')

  const handleSubmit = (e: React.FormEvent) => {
    e.preventDefault()
    if (game && winner) {
      db.collection('games').add({
        game,
        winner
      })
      setGame('')
      setWinner('')
    }
  }

  return (
    <form onSubmit={handleSubmit}>
      <input type="text" placeholder="Enter game" value={game} onChange={e => setGame(e.target.value)} />
      <input type="text" placeholder="Enter winner" value={winner} onChange={e => setWinner(e.target.value)} />
      <button type="submit">Add Game</button>
    </form>
  )
}

export default GameForm;