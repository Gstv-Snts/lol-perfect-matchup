import axios from "axios";
import React, { useEffect } from "react";

function App() {
  useEffect(() => {
    axios.get("https://www.leagueofgraphs.com/champions/counters/aatrox");
  }, []);

  return (
    <div className="App">
      <header className="App-header">
        <h1>WHAT ROLE ARE YOU PLAYING</h1>

        <h1>WHAT CHAMPION ARE YOU AGAINST</h1>
        <h1>WHO ARE YOUR TEAM:</h1>
        <h1>TOP:</h1>
        <h1>MID</h1>
        <h1>JUNGLE</h1>
        <h1>AD:</h1>
        <h1>SUP:</h1>
      </header>
    </div>
  );
}

export default App;
