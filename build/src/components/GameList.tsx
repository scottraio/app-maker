import React, { useState, useEffect } from 'react';
import { db } from '../firebase';

interface GameData {
  id: string,
  game: string,
  winner: string
}

const GameList = () => {
  const [games, setGames] = useState<GameData[]>([])

  useEffect(() => {
    const unsubscribe = db.collection('games').onSnapshot((snapshot) => {
      const gamesData: GameData[] = []
      snapshot.forEach((doc) => {
        gamesData.push({ id: doc.id, ...doc.data() } as GameData)
      })
      setGames(gamesData)
    })
    return unsubscribe
  }, [])

  return (
    <div>
      {games.map((game) => {
        return (
          <div key={game.id}>
            <h3>{game.game}</h3>
            <p>Winner: {game.winner}</p>
          </div>
        )
      })}
    </div>
  )
}

export default GameList;