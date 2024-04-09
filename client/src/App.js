import React from "react";
import "./App.css";
import {Container} from "semantic-ui-react";
import Tarefas from './Tarefas.js';

function App() {
  return (
    <div>
      <Container className="main">
        <Tarefas />
      </Container>
    </div>
  )
}

export default App;